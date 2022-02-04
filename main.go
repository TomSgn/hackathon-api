package main

import (
	"hackathon-api/configs"
	"hackathon-api/metrics"
	"hackathon-api/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	metrics.RegisterMetrics()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"OPTIONS", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(metrics.PrometheusMiddleware())

	//run database
	configs.ConnectDB()

	//routes
	routes.DonationRoute(router)
	routes.StatisticsRoute(router)
	routes.CacheRoute(router)
	routes.Metrics(router)

	router.Run("0.0.0.0:8080")
}
