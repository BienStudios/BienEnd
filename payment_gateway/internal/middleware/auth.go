package middleware

import (
	"crypto/subtle"

	"github.com/gin-gonic/gin"
)

// RequireAPIKey exige que el request incluya un header X-API-Key igual al
// configurado para el servicio. Es una alternativa más simple al JWT
// normalmente utilizado, suficiente para proteger este servicio de
// referencia; en producción se reemplaza por validación de JWT firmado.
//
// La comparación usa subtle.ConstantTimeCompare para no filtrar, por el
// tiempo de respuesta, cuántos caracteres iniciales del header coinciden
// con la clave real.
func RequireAPIKey(expectedKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		provided := c.GetHeader("X-API-Key")

		if provided == "" || subtle.ConstantTimeCompare([]byte(provided), []byte(expectedKey)) != 1 {
			c.AbortWithStatusJSON(401, gin.H{
				"code":    "unauthorized",
				"message": "credenciales inválidas o ausentes",
			})
			return
		}
		c.Next()
	}
}
