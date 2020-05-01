FROM golang:alpine as builder
ARG APP_NAME=hellohttp
WORKDIR /go/src/$APP_NAME
ENV GO111MODULE=on
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags '-s -w -extldflags "-static"' -o /go/bin/app main.go conf.go

FROM alpine:latest as alpine
RUN apk --no-cache add tzdata zip ca-certificates
WORKDIR /usr/share/zoneinfo
# -0=no compression. Go's tz loader doesn't handle compressed data.
RUN zip -r -0 /zoneinfo.zip .

FROM scratch
WORKDIR /app/
USER 1000
COPY --from=builder /go/bin/app .
ENV ZONEINFO /zoneinfo.zip
COPY --from=alpine /zoneinfo.zip /
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
CMD ["./app"]
