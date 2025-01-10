package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OpenDataTelemetry/ingestion-api/database"
	"github.com/gin-gonic/gin"
)

func GetAllAllSmartLight(c *gin.Context) {
	intervalStr := c.Query("interval")
	interval, err := strconv.Atoi(intervalStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid interval value"})
		return
	}

	if interval > 57600 {
		c.JSON(400, gin.H{"error": "Interval must be less than 57600"})
		return
	}

	var objs = []gin.H{}
	influxDB, err := database.ConnectToDB()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer influxDB.Close()

	query := `
		SELECT *
		FROM "EnergyMeter"
		WHERE "time" >= now() - interval '` + intervalStr + ` minutes'
		ORDER BY time DESC;
	`

	iterator, err := influxDB.Query(context.Background(), query) // Create iterator from query response

	if err != nil {
		panic(err)
	}

	for iterator.Next() { // Iterate over query response
		value := iterator.Value() // Value of the current row
		obj := gin.H{
			"fields": gin.H{
				"forwardEnergy":  value["forwardEnergy"],
				"reverseEnergy":  value["reverseEnergy"],
				"boardVoltage":   value["boardVoltage"],
				"data":           value["data"],
				"fCnt":           value["fCnt"],
				"fPort":          value["fPort"],
				"rxAlt_0":        value["rxAlt_0"],
				"rxLat_0":        value["rxLat_0"],
				"rxLon_0":        value["rxLon_0"],
				"rxRssi_0":       value["rxRssi_0"],
				"rxSnr_0":        value["rxSnr_0"],
				"txBandWidth":    value["txBandWidth"],
				"txFrequency":    value["txFrequency"],
				"txSpreadFactor": value["txSpreadFactor"],
			},
			"name": "EnergyMeter",
			"tags": gin.H{
				"deviceId":   value["deviceId"],
				"deviceType": value["deviceType"],
				"direction":  value["direction"],
				"host":       value["host"],
				"origin":     value["origin"],
				"rxMac_0":    value["rxMac_0"],
				// "txCodeRate":   value["txCodeRate"],
				// "txModulation": value["txModulation"],
				"type": value["type"],
			},
			"timestamp": value["time"],
		}
		// Convert the row to a gin.H map (JSON)
		objs = append(objs, obj) // Append the row to the objs slice
	}
	fmt.Println(len(objs))
	c.IndentedJSON(http.StatusOK, objs)
}
