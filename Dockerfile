FROM golang:1.8

WORKDIR /go/src/examplegoapi

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["cmd"]

HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD [ "curl -f http://127.0.0.1:8000/Status || exit 1" ]

EXPOSE 8000
