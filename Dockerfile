FROM golang
WORKDIR /var/app
COPY . .
RUN go build -o main

FROM golang
WORKDIR /var/app
ENV GCP_API_KEY="dontForget"
COPY --from=0 /var/app .
CMD ["/var/app/main"]
EXPOSE 8080/tcp
