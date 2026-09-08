package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IdempotencyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-Idempotency-Key")
		if key == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": "El header X-Idempotency-Key es requerido",
			})
			return
		}

		// 1. Verificar en Redis/DB si la clave 'key' ya fue procesada.
		// 2. Si ya fue procesada, retornar la respuesta guardada previamente.
		// 3. Si no, continuar con c.Next() y guardar el resultado al finalizar.

		c.Next()
	}
}
