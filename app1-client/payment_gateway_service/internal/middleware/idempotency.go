// Package middleware contiene los middlewares de Gin del servicio: exigir
// una clave de idempotencia, autenticación, y cualquier otro cross-cutting
// concern que no pertenece a un handler específico.
package middleware

import "github.com/gin-gonic/gin"

// RequireIdempotencyKey exige el header Idempotency-Key en el request y lo
// deja disponible en el contexto de Gin bajo la clave "idempotency_key" para
// que el handler lo pase al servicio de aplicación. La reserva atómica de la
// clave ocurre dentro del propio caso de uso (service.PaymentService),
// no aquí — este middleware solo valida su presencia.
func RequireIdempotencyKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("Idempotency-Key")
		if key == "" {
			c.AbortWithStatusJSON(400, gin.H{
				"code":    "missing_idempotency_key",
				"message": "el header Idempotency-Key es requerido",
			})
			return
		}
		c.Set("idempotency_key", key)
		c.Next()
	}
}
