FROM asia.gcr.io/om-hkt21-team23/golang:1.18-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./build/service ./service/main.go
RUN CGO_ENABLED=0 go build -o ./build/worker ./worker/*.go

FROM asia.gcr.io/om-hkt21-team23/golang:1.18-alpine
WORKDIR /app
COPY --from=builder /app/build/service .
RUN chmod +x /app/service
COPY --from=builder /app/build/worker .
RUN chmod +x /app/worker
