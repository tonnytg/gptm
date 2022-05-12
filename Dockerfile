#Build Builder aplication
FROM golang:1.18-bullseye
WORKDIR /var/app
COPY . .
RUN go build -o main

#Build scratch less than 16kb + binary total 1.94Mb
FROM golang:1.18-bullseye
WORKDIR /var/app
COPY --from=0 /var/app .
ENTRYPOINT ["/var/app/main"]
