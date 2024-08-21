FROM golang:1.22.3-alpine as builder
COPY . /app
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint

FROM alpine AS release-stage
WORKDIR /
COPY --from=builder /entrypoint /entrypoint
RUN apk add libcap && setcap 'cap_net_bind_service=+ep' /entrypoint
USER 1000
ENTRYPOINT ["/entrypoint"]
