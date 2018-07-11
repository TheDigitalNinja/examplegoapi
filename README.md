# Example Go API

[![Go Report Card](https://goreportcard.com/badge/github.com/TheDigitalNinja/examplegoapi)](https://goreportcard.com/report/github.com/TheDigitalNinja/examplegoapi)
[![Build Status](https://travis-ci.org/TheDigitalNinja/examplegoapi.svg?branch=master)](https://travis-ci.org/TheDigitalNinja/examplegoapi)

I needed some services for testing deployments so I thought I would learn GoLang and just roll my own. 

## Run

```
cd cmd
go run main.go
```

## Tests

Run tests with

```
go test -cover ./...
```

## Directories

### `/cmd`

Main applications for this project.

### `/web`

Web application specific components: static web assets, server side templates and SPAs.

## Endpoints

### `/`

Example Web Page. Should return `Content-Type: text/html; charset=utf-8`

### `/Status`

Health Check. Should return json `{"STATUS": "UP"}`
