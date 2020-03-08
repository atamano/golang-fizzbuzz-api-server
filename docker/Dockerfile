# Step1: build
FROM golang:alpine AS build
RUN apk update && \
    apk add curl \
            git \
            bash \
            make \
            ca-certificates && \
    rm -rf /var/cache/apk/*

WORKDIR /app

# copy module files first so that they don't need to be downloaded again if no change
COPY go.* ./
RUN go mod download
RUN go mod verify

# copy source files and build the binary
COPY . .
ENV CGO_ENABLED=0
RUN go build -a -o server /app/cmd/server

# Step2: run
FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
RUN mkdir -p /var/log/app
WORKDIR /app/

COPY --from=build /app/server .
RUN ls -la

CMD ["./server"]