module github.com/eduardocfalcao/money-tracker

replace github.com/nats-io/go-nats => github.com/nats-io/nats.go v1.10.0

go 1.20

require (
	github.com/auth0/go-jwt-middleware v0.0.0-20200507191422-d30d7b9ece63
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-chi/chi/v5 v5.0.10
	github.com/sirupsen/logrus v1.9.3
)

require golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
