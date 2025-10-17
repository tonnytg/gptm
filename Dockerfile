FROM golang:1.25.3-alpine3.21

WORKDIR /var/app

ENTRYPOINT ["/var/app/main"]
