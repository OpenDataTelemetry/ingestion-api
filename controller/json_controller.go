package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strings"

	"github.com/OpenDataTelemetry/ingestion-api/mqtt"
	"github.com/gin-gonic/gin"
)

func HandleAllLnsCommandIngestion(c *gin.Context) {
	var jsonMessageMap map[string]interface{}
	if err := c.ShouldBindJSON(&jsonMessageMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonMessage, err := json.Marshal(jsonMessageMap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var measurement = "Downlink"
	var organization = "IMT"
	var deviceType = "LNS"
	var etc = jsonMessageMap["etc"]

	if etc == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'etc' key in json.",
		})
		return
	}
	var application = jsonMessageMap["application"]
	if application == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'application' key in json.",
		})
		return
	}
	var reference = jsonMessageMap["reference"]
	if reference == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'reference' key in json.",
		})
		return
	}
	var deviceId = jsonMessageMap["deviceId"]
	if deviceId == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'deviceId' key in json.",
		})
		return
	}
	var confirmed = jsonMessageMap["confirmed"]
	if confirmed == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'confirmed' key in json.",
		})
		return
	}
	var fPort = jsonMessageMap["fPort"]
	if fPort == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'fPort' key in json.",
		})
		return
	}
	var data = jsonMessageMap["data"]
	if data == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'data' key in json.",
		})
		return
	}
	var timestamp = jsonMessageMap["timestamp"]
	if timestamp == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'timestamp' key in json.",
		})
		return
	}

	var sb strings.Builder
	sb.WriteString(`OpenDataTelemetry/`)
	sb.WriteString(organization)
	sb.WriteString(`/`)
	sb.WriteString(deviceType)
	sb.WriteString(`/`)
	sb.WriteString(measurement)
	sb.WriteString(`/`)
	sb.WriteString(deviceId.(string))
	sb.WriteString(`/down/`)
	sb.WriteString(etc.(string))

	go mqtt.PubToBroker(sb.String(), string(jsonMessage))
	fmt.Println("topic:", sb.String())

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func HandleAllLnsAlertIngestion(c *gin.Context) {
	var jsonMessageMap map[string]interface{}
	if err := c.ShouldBindJSON(&jsonMessageMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonMessage, err := json.Marshal(jsonMessageMap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var organization = "IMT"
	var deviceType = "LNS"
	var measurement = "Alert"
	var deviceId = jsonMessageMap["deviceId"]

	if deviceId == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'deviceId' key in json.",
		})
		return
	}
	var etc = jsonMessageMap["etc"]
	if etc == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'etc' key in json.",
		})
		return
	}
	var timestamp = jsonMessageMap["timestamp"]
	if timestamp == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'timestamp' key in json.",
		})
		return
	}

	var sb strings.Builder
	sb.WriteString(`OpenDataTelemetry/`)
	sb.WriteString(organization)
	sb.WriteString(`/`)
	sb.WriteString(deviceType)
	sb.WriteString(`/`)
	sb.WriteString(measurement)
	sb.WriteString(`/`)
	sb.WriteString(deviceId.(string))
	sb.WriteString(`/alert/`)
	sb.WriteString(etc.(string))

	go mqtt.PubToBroker(sb.String(), string(jsonMessage))
	fmt.Println("topic:", sb.String())

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func HandleAllNspiGenericJsonIngestion(c *gin.Context) {
	var jsonMessageMap map[string]interface{}
	if err := c.ShouldBindJSON(&jsonMessageMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// fmt.Println("jsonMessage: %v\n", jsonMessage)

	jsonMessage, err := json.Marshal(jsonMessageMap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var organization = "IMT"
	var deviceType = "NSPI"
	var measurement = "GenericJson"
	var deviceId = jsonMessageMap["deviceId"]

	if deviceId == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'deviceId' key in json.",
		})
		return
	}
	var etc = jsonMessageMap["etc"]
	if etc == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'etc' key in json.",
		})
		return
	}
	var timestamp = jsonMessageMap["timestamp"]
	if timestamp == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'timestamp' key in json.",
		})
		return
	}

	var sb strings.Builder
	sb.WriteString(`OpenDataTelemetry/`)
	sb.WriteString(organization)
	sb.WriteString(`/`)
	sb.WriteString(deviceType)
	sb.WriteString(`/`)
	sb.WriteString(measurement)
	sb.WriteString(`/`)
	sb.WriteString(deviceId.(string))
	sb.WriteString(`/up/`)
	sb.WriteString(etc.(string))

	go mqtt.PubToBroker(sb.String(), string(jsonMessage))
	fmt.Println("topic:", sb.String())

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func HandleAllNspiAlertIngestion(c *gin.Context) {
	var jsonMessageMap map[string]interface{}
	if err := c.ShouldBindJSON(&jsonMessageMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jsonMessage, err := json.Marshal(jsonMessageMap)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var organization = "IMT"
	var deviceType = "NSPI"
	var measurement = "Alert"
	var deviceId = jsonMessageMap["deviceId"]

	if deviceId == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'deviceId' key in json.",
		})
		return
	}
	var etc = jsonMessageMap["etc"]
	if etc == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'etc' key in json.",
		})
		return
	}
	var timestamp = jsonMessageMap["timestamp"]
	if timestamp == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'timestamp' key in json.",
		})
		return
	}

	var sb strings.Builder
	sb.WriteString(`OpenDataTelemetry/`)
	sb.WriteString(organization)
	sb.WriteString(`/`)
	sb.WriteString(deviceType)
	sb.WriteString(`/`)
	sb.WriteString(measurement)
	sb.WriteString(`/`)
	sb.WriteString(deviceId.(string))
	sb.WriteString(`/up/`)
	sb.WriteString(etc.(string))

	go mqtt.PubToBroker(sb.String(), string(jsonMessage))
	fmt.Println("topic:", sb.String())

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}
