package main

import (
	"github.com/OpenDataTelemetry/ingestion-api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Create a new gin router instance

	r.Use(cors.Default())

	api := r.Group("/api/ingestion/v0.1/IMT/")
	{
		api.POST("/NSPI/GenericJson/all", controller.HandleAllNspiGenericJsonUpIngestion)

		api.POST("/LNS/LnsDownlink/all", controller.HandleAllLnsDownlinkIngestion)

		// api.POST("/LNS/Alert/all", controller.HandleAllLnsAlertIngestion)
	}

	r.Run(":8888")
}
