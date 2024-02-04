module github.com/eduardocfalcao/money-tracker

replace github.com/nats-io/go-nats => github.com/nats-io/nats.go v1.10.0

go 1.21

require (
	github.com/go-chi/chi/v5 v5.0.10
	github.com/golang-jwt/jwt/v5 v5.2.0
	github.com/google/uuid v1.6.0
	github.com/jackc/pgx/v5 v5.5.3
	github.com/sirupsen/logrus v1.9.3
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
