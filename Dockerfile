FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN make build

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /app/build/docker-registry-proxy /main
ENTRYPOINT ["/main"]
CMD ["--address", "localhost:8989"]