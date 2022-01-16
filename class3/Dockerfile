#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go mod init
RUN go mod tidy

RUN go build -o /go/bin/app -v ./src/

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app
LABEL Name=httpserver Version=1.0
EXPOSE 8000

