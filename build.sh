# Build App Jenkins Script

export GOPATH="$WORKSPACE"

cd src/examplegoapi

go get -t ./...

figlet "Unit Tests"
go test -cover ./...

figlet "Build"
go build -o examplegoapi.exe ./cmd/
