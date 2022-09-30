package Middlewares

import (
	"Seller/Seller.Application/Common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		Claims, _ := Common.GetClaimsFromToken(strings.ReplaceAll(strings.ReplaceAll(token, "bearer ", ""), "Bearer ", ""))
		if Claims.VerifyExpiresAt(time.Now().Unix(), false) == false {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		if token == "" {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
