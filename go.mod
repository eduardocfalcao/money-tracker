module github.com/eduardocfalcao/money-tracker

replace github.com/nats-io/go-nats => github.com/nats-io/nats.go v1.10.0

go 1.21

require (
	github.com/go-chi/chi/v5 v5.0.10
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/sirupsen/logrus v1.9.3
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
