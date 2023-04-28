FROM golang:1.20-alpine AS base

# Install dependency based on package manager
FROM base AS deps
WORKDIR /app
COPY ../go.mod go.sum ./
RUN go mod download

# Rebuild the source code only when needed
FROM base AS go-builder
WORKDIR /app
COPY --from=deps /go/pkg/ /go/pkg/
COPY .. .

RUN go build -ldflags="-X github.com/zjc17/exif-web/internal/version.Version=1.0" -o ./exif-web

# Runner
FROM alpine AS runner
RUN rm -rf /var/cache/apk/*
WORKDIR /
COPY --from=go-builder /app/exif-web /bin/exif-web

EXPOSE 8080
ENTRYPOINT ["exif-web"]