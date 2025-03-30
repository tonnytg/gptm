FROM golang:latest

WORKDIR /var/app

ENTRYPOINT ["/var/app/main"]
