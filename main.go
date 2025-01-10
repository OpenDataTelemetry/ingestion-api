package main

import (
	"github.com/OpenDataTelemetry/ingestion-api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Create a new gin router instance

	r.Use(cors.Default())

	api := r.Group("/api/ingestion/v0.3/IMT/LNS/")
	{
		api.GET("SmartLight/all", controller.GetAllSmartLight)
	}

	r.Run(":8888")
}
