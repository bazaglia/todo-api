FROM golang:1.11-alpine as builder

WORKDIR /src/app
COPY . ./

RUN apk add --no-cache git openssh-client ca-certificates tzdata && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /src/app/main /app
CMD ["/app"]

EXPOSE 8000
