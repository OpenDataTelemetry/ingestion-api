package mqtt

import (
	// "context"
	"fmt"
	"os"
	"strings"

	// "github.com/OpenDataTelemetry/ingestion-api/controller"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
)

func connLostHandler(c MQTT.Client, err error) {
	fmt.Printf("Connection lost, reason: %v\n", err)
	os.Exit(1)
}

func PubToBroker(t string, m string) {
	id := uuid.New().String()
	var sbMqttPubClientId strings.Builder
	// var sbPubTopic strings.Builder
	sbMqttPubClientId.WriteString("ingestion-api-")
	sbMqttPubClientId.WriteString(id)

	mqttPubBroker := "mqtt://mqtt.maua.br:1883"
	mqttPubClientId := sbMqttPubClientId.String()
	mqttPubUser := ""
	mqttPubPassword := ""
	mqttPubQos := 0

	mqttPubOpts := MQTT.NewClientOptions()
	mqttPubOpts.AddBroker(mqttPubBroker)
	mqttPubOpts.SetClientID(mqttPubClientId)
	mqttPubOpts.SetUsername(mqttPubUser)
	mqttPubOpts.SetPassword(mqttPubPassword)

	pClient := MQTT.NewClient(mqttPubOpts)
	if token := pClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", mqttPubBroker)
	}

	fmt.Printf("topic: %s, & message: %v\n", t, m)
	token := pClient.Publish(t, byte(mqttPubQos), false, m)
	token.Wait()
	pClient.Disconnect(250)
	fmt.Printf("Disconnected from %s\n", mqttPubBroker)
}
