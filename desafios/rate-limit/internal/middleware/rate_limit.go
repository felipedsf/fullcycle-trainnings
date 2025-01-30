package middleware

import (
	"github.com/felipedsf/fullcycle-trannings/desafios/rate-limit/internal/service"
	"github.com/gin-gonic/gin"
)

func RateLimit(svc *service.RateLimitService) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("API_KEY")
		clientIp := c.ClientIP()
		if apiKey == "" {
			apiKey = clientIp
		}
		if err := svc.ProcessRequest(apiKey, clientIp); err != nil {
			c.AbortWithStatusJSON(
				429,
				gin.H{
					"message": "you have reached the maximum number of requests or actions allowed within a certain time frame",
				},
			)
		}
		c.Next()
	}
}
