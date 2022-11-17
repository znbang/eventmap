FROM node:18-alpine AS base
COPY --from=golang:alpine /usr/local/go /usr/local/go
ENV GOPATH /go
ENV PATH "$GOPATH/bin:/usr/local/go/bin:$PATH"
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
RUN apk add gcc musl-dev
RUN go install github.com/go-task/task/v3/cmd/task@latest
RUN go install github.com/bufbuild/buf/cmd/buf@latest
RUN go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN npm i -g @quasar/cli @bufbuild/protoc-gen-connect-web @bufbuild/protoc-gen-es

FROM base AS builder
WORKDIR /app
COPY . .
RUN task build

FROM alpine:latest
COPY --from=builder /app/eventmapd .
ENTRYPOINT ["/eventmapd"]
