package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strings"

	"github.com/OpenDataTelemetry/ingestion-api/mqtt"
	"github.com/gin-gonic/gin"
)

func HandleAllNspiGenericJsonUpIngestion(c *gin.Context) {
	var jsonMessageMap map[string]interface{}
	// var data string
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

	// jsonStr := string(jsonMessage)
	// fmt.Println(jsonStr)

	// OpenDataTelemetry/IMT/LNS/MEASUREMENT/DEVICE_ID/up/etc
	// VERIFY IF EACH KEY IS NOT NIL
	var organization = "IMT" // FORCE TO IMT DEVICE_TYPE
	// var organization = jsonMessageMap["organization"]
	var deviceType = "NSPI" // FORCE TO NSPI DEVICE_TYPE
	// var deviceType = jsonMessageMap["deviceType"]
	var measurement = jsonMessageMap["measurement"]
	if measurement == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'measurement' key in json.",
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
	// sb.WriteString(organization.(string))
	sb.WriteString(organization)
	sb.WriteString(`/`)
	sb.WriteString(deviceType)
	// sb.WriteString(deviceType.(string))
	sb.WriteString(`/`)
	sb.WriteString(measurement.(string))
	sb.WriteString(`/`)
	sb.WriteString(deviceId.(string))
	sb.WriteString(`/up/`)
	sb.WriteString(etc.(string))

	go mqtt.PubToBroker(sb.String(), string(jsonMessage))
	fmt.Println("topic:", sb.String())

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

func HandleAllLnsDownlinkIngestion(c *gin.Context) {
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
	// OpenDataTelemetry/IMT/LNS/LnsDownlink/0004a30b00286d19/down/imt
	// VERIFY IF EACH KEY IS NOT NIL
	var measurement = "LnsDownlink"
	var organization = "IMT" // FORCE TO IMT DEVICE_TYPE
	var deviceType = "LNS"   // FORCE TO NSPI DEVICE_TYPE
	var etc = jsonMessageMap["etc"]
	if etc == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'etc' key in json.",
		})
		return
	}
	var application = jsonMessageMap["application"] //"DET"
	if application == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'application' key in json.",
		})
		return
	}
	var reference = jsonMessageMap["reference"] //"test-node-red"
	if reference == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'reference' key in json.",
		})
		return
	}
	var deviceId = jsonMessageMap["deviceId"] //"0004a30b00286d19"
	if deviceId == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'deviceId' key in json.",
		})
		return
	}
	var confirmed = jsonMessageMap["confirmed"] //false
	if confirmed == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'confirmed' key in json.",
		})
		return
	}
	var fPort = jsonMessageMap["fPort"] //100
	if fPort == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"msg":    "Invalid input, please check your data. Missing 'fPort' key in json.",
		})
		return
	}
	var data = jsonMessageMap["data"] //"AAE="
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

// func HandleAllLnsAlertIngestion(c *gin.Context) {
// 	var jsonMessageMap map[string]interface{}

// 	if err := c.ShouldBindJSON(&jsonMessageMap); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	jsonMessage, err := json.Marshal(jsonMessageMap)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	// OpenDataTelemetry/IMT/LNS/LnsDownlink/0004a30b00286d19/down/imt
// 	// VERIFY IF EACH KEY IS NOT NIL
// 	var measurement = "LnsAlert"
// 	var organization = "IMT" // FORCE TO IMT DEVICE_TYPE
// 	var deviceType = "LNS"   // FORCE TO NSPI DEVICE_TYPE
// 	var etc = jsonMessageMap["etc"]
// 	if etc == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": "error",
// 			"msg":    "Invalid input, please check your data. Missing 'etc' key in json.",
// 		})
// 		return
// 	}
// 	var application = jsonMessageMap["application"] //"DET"
// 	if application == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": "error",
// 			"msg":    "Invalid input, please check your data. Missing 'application' key in json.",
// 		})
// 		return
// 	}

// 	var deviceId = jsonMessageMap["deviceId"] //"0004a30b00286d19"
// 	if deviceId == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": "error",
// 			"msg":    "Invalid input, please check your data. Missing 'deviceId' key in json.",
// 		})
// 		return
// 	}

// 	var data = jsonMessageMap["data"] //"AAE="
// 	if data == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": "error",
// 			"msg":    "Invalid input, please check your data. Missing 'data' key in json.",
// 		})
// 		return
// 	}
// 	var timestamp = jsonMessageMap["timestamp"]
// 	if timestamp == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status": "error",
// 			"msg":    "Invalid input, please check your data. Missing 'timestamp' key in json.",
// 		})
// 		return
// 	}

// 	var sb strings.Builder
// 	sb.WriteString(`OpenDataTelemetry/`)
// 	sb.WriteString(organization)
// 	sb.WriteString(`/`)
// 	sb.WriteString(deviceType)
// 	sb.WriteString(`/`)
// 	sb.WriteString(measurement)
// 	sb.WriteString(`/`)
// 	sb.WriteString(deviceId.(string))
// 	sb.WriteString(`/down/`)
// 	sb.WriteString(etc.(string))

// 	go mqtt.PubToBroker(sb.String(), string(jsonMessage))
// 	fmt.Println("topic:", sb.String())

// 	c.JSON(http.StatusOK, gin.H{"data": "ok"})
// }
