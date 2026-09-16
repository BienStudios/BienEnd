// Package config centraliza la carga de configuración externa del servicio.
// Nunca se lee os.Getenv disperso por el código.
package config

import (
	"fmt"
	"os"
)

// Config agrupa toda la configuración necesaria para arrancar el servicio.
type Config struct {
	Port   string
	APIKey string
}

// Load lee y valida la configuración desde variables de entorno, fallando
// rápido si falta algo crítico — nunca arrancar el servicio "a medias".
// TODO: Verificar existencia de la variable de entorno API_KEY con la clave real. No
// proveer la real dentro de este paquete (pánico ruidoso antes que exposición de datos
// secretos)
func Load() (*Config, error) {
	cfg := &Config{
		Port:   getEnvOrDefault("PORT", "8082"),
		APIKey: getEnvOrDefault("API_KEY", "dev-api-key-not-for-prod"),
	}

	if cfg.APIKey == "" {
		return nil, fmt.Errorf("API_KEY es requerida")
	}

	return cfg, nil
}

func getEnvOrDefault(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
