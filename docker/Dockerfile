FROM golang:alpine as builder

ENV GO111MODULE=on

WORKDIR /build

COPY . .

RUN apk add --no-cache git \
    && go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-multitenancy

# dist image
FROM alpine

COPY --from=builder /build/go-multitenancy /app/
COPY migrations/ /app/migrations
RUN addgroup -S arkanjoms && adduser -S arkanjoms -G arkanjoms \
    && chown -R arkanjoms:arkanjoms /app
USER arkanjoms
WORKDIR /app

EXPOSE 8080
ENTRYPOINT ["./go-multitenancy"]