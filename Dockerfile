FROM golang:1.18-bullseye
FROM golang:1
WORKDIR /var/app
COPY --from=0 /var/app .
ENTRYPOINT ["/var/app/main"]
