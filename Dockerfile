FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add bash

WORKDIR /build
COPY . /build

RUN go get -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/money-tracker ./cmd/server

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/money-tracker /go/bin/money-tracker
EXPOSE 8080

ENTRYPOINT  ["/go/bin/money-tracker"]