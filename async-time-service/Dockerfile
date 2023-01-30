FROM golang:1.19 AS builder
WORKDIR /async-time-service
COPY . /async-time-service/
RUN go mod download
RUN go build -o ./async-time-service main.go
#deploy stage
FROM alpine:latest
WORKDIR /
COPY --from=builder /async-time-service /async-time-service
RUN chmod +x ./async-time-service
EXPOSE 8080
ENTRYPOINT [ "./async-time-service" ]