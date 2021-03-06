# rest template

A template for writing rest services in Go. Based on a post by Mat Ryer: 

https://medium.com/statuscode/how-i-write-go-http-services-after-seven-years-37c208122831 How I write Go HTTP services after seven years

## Features

This sample rest client has the following features pre-added:

- Prometheus metrics export
- GOPS

## History

|Version|Description|
|---|---|
|1.0.1|update dependencies|
|1.0.0|switch to go modules|
|0.2.0|make listen and port configurable by default|
|0.1.1|add .go-template.yml|
|0.1.0|- sample static pages|
||- wildcard route|
||- fixed route|
||- gops https://github.com/google/gops|
||- prometheus metric at /metrics|
||- templates route|
