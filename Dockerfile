FROM golang:1.24.1-alpine3.21

WORKDIR /var/app

ENTRYPOINT ["/var/app/main"]
