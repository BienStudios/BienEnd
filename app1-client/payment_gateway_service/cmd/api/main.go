// Command api es el punto de entrada del microservicio de pasarela de pagos.
// Se encarga únicamente de ensamblar las piezas concretas de cada puerto y
// arranca el servidor HTTP con apagado ordenado.
//
// Por defecto arranca con adaptadores en memoria/simulados, para poder
// correr el servicio completo sin infraestructura externa:
//
//	go mod tidy
//	go run ./cmd/api
//
// Para producción, reemplazar en este archivo:
//   - repository.NewMemoryPaymentRepository() -> repository.NewPostgresPaymentRepository(pool) (build tag "postgres")
//   - gateway.NewMockProcessor() -> gateway.NewHTTPProcessor(baseURL, apiKey)
//   - eventing.NewLogPublisher(logger) -> un publisher real sobre RabbitMQ
package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/BienStudios/BienEnd/payment_gateway/internal/config"
	"github.com/BienStudios/BienEnd/payment_gateway/internal/eventing"
	"github.com/BienStudios/BienEnd/payment_gateway/internal/gateway"
	"github.com/BienStudios/BienEnd/payment_gateway/internal/handler"
	"github.com/BienStudios/BienEnd/payment_gateway/internal/middleware"
	"github.com/BienStudios/BienEnd/payment_gateway/internal/repository"
	"github.com/BienStudios/BienEnd/payment_gateway/internal/service"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	cfg, err := config.Load()
	if err != nil {
		logger.Error("configuración inválida", slog.String("error", err.Error()))
		os.Exit(1)
	}

	// --- 1. Construir los adaptadores concretos de cada puerto ---
	paymentRepo := repository.NewMemoryPaymentRepository()
	idempotencyStore := repository.NewMemoryIdempotencyStore(24 * time.Hour)
	processor := gateway.NewMockProcessor()
	publisher := eventing.NewLogPublisher(logger)

	// --- 2. Inyectar los adaptadores en la capa de aplicación ---
	paymentService := service.NewPaymentService(paymentRepo, processor, publisher, idempotencyStore, logger)

	// --- 3. Inyectar el servicio de aplicación en el adaptador de entrada HTTP ---
	paymentHandler := handler.NewPaymentHandler(paymentService)

	// --- 4. Configurar el router ---
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/health/live", func(c *gin.Context) {
		// Si responde, ya está "vivo"
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	router.GET("/health/ready", func(c *gin.Context) {
		// TODO: Modificar cuando se necesiten las dependencias de Postgres y RabbitMQ reales
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	v1 := router.Group("/api/v26.00.0000_ALPHA")
	v1.Use(middleware.RequireAPIKey(cfg.APIKey))
	paymentHandler.RegisterRoutes(v1, middleware.RequireIdempotencyKey())

	// --- 5. Arrancar el servidor con apagado ordenado ---
	srv := &http.Server{Addr: ":" + cfg.Port, Handler: router}

	go func() {
		logger.Info("servidor escuchando", slog.String("port", cfg.Port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("error en el servidor", slog.String("error", err.Error()))
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("señal de apagado recibida, drenando conexiones...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("apagado forzado tras timeout", slog.String("error", err.Error()))
	}
	logger.Info("servidor apagado correctamente")
}
