# ingestion-api

## LNS - LnsDownlink
Post a Downlink Request to Specific LNS device:

Using Curl:
```bash
curl -X POST https://smartcampus-k8s.maua.br/api/ingestion/v0.1/IMT/LNS/LnsDownlink/all -d '{"application": "DET", "etc": "imt", "reference": "test-node-red", "deviceId": "0004a30b00286d19", "confirmed": false, "fPort": 100, "data": "AAE=", "timestamp": 1736459402000000000}' -H "Content-Type: application/json"
```

Http Method: Post
Host: https://smartcampus-k8s.maua.br/api/ingestion/v0.1/IMT/LNS/LnsDownlink/all
Data:
```json
{
  "application": "DET", // Application Name registered in the corresponding NetworkServer
  "etc": "imt", // NetworkServer to be queued
  "reference": "test-node-red", 
  "deviceId": "0004a30b00286d19", 
  "confirmed": false,
  "fPort": 100, // lora downlink fPort
  "data": "AAE=", // downlink data
  "timestamp": 1736459402000000000 // nanoseconds
  }
```

## NSPI - GenericJson
Post a message to OpenDataTelemetry from Specific NSPI device:

Using Curl:
```bash
curl -X POST https://smartcampus-k8s.maua.br/api/ingestion/v0.1/IMT/LNS/LnsDownlink/all -d '{"measurement": "Masak", "deviceId": "0004a30b00286d19", "etc": "imt", "data": "AAE=", "timestamp": 1736459402000000000}' -H "Content-Type: application/json"
```

Http Method: Post
Host: https://smartcampus-k8s.maua.br/api/ingestion/v0.1/IMT/NSPI/GenericJson/all
Data:
```json
{
  "measurement": "Masak", // mandatory
  "deviceId": "0004a30b00286d19", // mandatory
  "etc": "imt", // mandatory
  "data": "AAE=", 
  "timestamp": 1736459402000000000 // mandatory
  }
```