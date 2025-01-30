package main

import (
	"github.com/felipedsf/fullcycle-trannings/desafios/rate-limit/internal/config"
	"github.com/felipedsf/fullcycle-trannings/desafios/rate-limit/internal/database"
	"github.com/felipedsf/fullcycle-trannings/desafios/rate-limit/internal/middleware"
	"github.com/felipedsf/fullcycle-trannings/desafios/rate-limit/internal/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	//db := database.GetMemoryDb()
	db := database.GetRedisDb()
	svc := service.NewService(db)

	r := gin.Default()
	r.Use(middleware.RateLimit(svc))
	r.GET("/test", func(c *gin.Context) {
		client := c.ClientIP()
		c.JSON(200, gin.H{
			"message": "Hello World",
			"client":  client,
		})
	})
	if err := r.Run(config.AppConfig.Port); err != nil {
		log.Fatal(err)
	}
}
