# STEP 1 build executable binary

FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /build
COPY . /build

# Fetch dependencies.
# Using go get.
RUN go get -v ...
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/money-tracker ./cmd/server

# STEP 2 build a small image

FROM scratch

COPY --from=builder /go/bin/money-tracker /go/bin/money-tracker
EXPOSE 8080

ENTRYPOINT  ["/go/bin/money-tracker"]