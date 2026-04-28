FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init -g cmd/api/main.go -o docs

RUN CGO_ENABLED=0 GOOS=linux go build \
-ldflags="-w -s" \
-o /app/bin/api \
./cmd/api

FROM alpine:3.23

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

RUN apk --no-cache add ca-certificates tzdata
ENV TZ=Asia/Jakarta

WORKDIR /app

COPY --from=builder /app/bin/api .
COPY .env .
COPY --from=builder /app/docs ./docs

USER appuser

EXPOSE 8080
CMD ["./api"]