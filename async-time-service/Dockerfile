#builder stage
FROM golang:1.19 AS builder
WORKDIR /async-time-service
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o bin/async-time-service main.go
#deploy stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /async-time-service/bin/async-time-service ./
RUN chmod -R 775 ./
EXPOSE 8080
ENTRYPOINT [ "./async-time-service" ]