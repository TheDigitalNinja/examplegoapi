# Example Go API

I needed some services for testing deployments so I thought I would learn GoLang and just roll my own. 

## Run

```
cd cmd
go run main.go
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

## Tests

Run tests with

```
go test -cover ./...
```