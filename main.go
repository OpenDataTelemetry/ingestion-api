package main

import (
	"github.com/OpenDataTelemetry/ingestion-api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Create a new gin router instance

	r.Use(cors.Default())

	api := r.Group("/api/ingestion/v0.1/")
	{

		api.POST("IMT/LNS/Command/all", controller.HandleAllLnsCommandIngestion)
		api.POST("IMT/LNS/Alert/all", controller.HandleAllLnsAlertIngestion)

		api.POST("IMT/NSPI/GenericJson/all", controller.HandleAllNspiGenericJsonIngestion)
		api.POST("IMT/NSPI/Alert/all", controller.HandleAllNspiAlertIngestion)

	}

	r.Run(":8888")
}
