FROM node:20-alpine AS base
COPY --from=golang:alpine /usr/local/go /usr/local/go
ENV GOPATH /go
ENV PATH "$GOPATH/bin:/usr/local/go/bin:$PATH"
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
RUN apk add --no-cache gcc musl-dev

FROM base AS builder
ARG GOOGLE_MAPS_API_KEY
ENV GOOGLE_MAPS_API_KEY=${GOOGLE_MAPS_API_KEY}
WORKDIR /app
COPY . .
RUN go install github.com/go-task/task/v3/cmd/task@latest
RUN task tools
RUN task assets
RUN task build

FROM alpine:latest
COPY --from=builder /app/eventmapd .
ENTRYPOINT ["/eventmapd"]
