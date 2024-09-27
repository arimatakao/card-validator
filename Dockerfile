# 1 stage. Builder
FROM golang:1.22.4 AS builder

ARG CGO_ENABLED=0
WORKDIR /app

COPY . .
RUN go mod tidy
RUN go build -o ./card-validator main.go

# 2 stage. Runner
FROM scratch
COPY --from=builder /app/card-validator /bin/card-validator

EXPOSE 8080

ENTRYPOINT ["/bin/card-validator"]