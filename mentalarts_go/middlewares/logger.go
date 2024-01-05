package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *MentalArtsMiddleware) Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		fmt.Printf("Request to %s took %s\n", c.Request.URL.Path, time.Since(start))
	}
}
