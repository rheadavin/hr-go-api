package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rheadavin/hr-go-api/pkg/jwt"
	response "github.com/rheadavin/hr-go-api/pkg/reponse"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			response.ErrorResponse(c, http.StatusUnauthorized, "Token diperlukan")
			c.Abort()
			return
		}

		claims, err := jwt.ValidateToken(token)
		if err != nil {
			response.ErrorResponse(c, http.StatusUnauthorized, "Token tidak valid atau kadaluarsa")
			c.Abort()
			return
		}
		c.Set("user_id", claims.UserId)
		c.Set("user_email", claims.Email)
		c.Set("user_role", claims.Role)
		c.Next()
	}

}
