FROM golang:1.25-trixie AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM debian:trixie-slim AS runtime

RUN apt-get update && \
    apt-get install -y --no-install-recommends chromium && \
    ln -sf /usr/bin/chromium /usr/bin/chromium-browser && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server

COPY --from=builder /app/public /app/public

WORKDIR /app

ENV CHROME_PATH=/usr/bin/chromium

EXPOSE 8080

CMD ["/app/server"]
