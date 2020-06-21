package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/eduardocfalcao/money-tracker/http/auth/middleware"
	"github.com/eduardocfalcao/money-tracker/http/handlers"
	"github.com/eduardocfalcao/money-tracker/http/handlers/info"
	"github.com/gorilla/mux"
)

func get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

var logger *log.Logger = log.New(os.Stdout, "[money-tracker] ", 0)

func logRequest(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger.Print("Request received: Haders:", r.Header)
	next.ServeHTTP(w, r)
}

// InitHTTPServer  initializates a new http server listening the port provided
func InitHTTPServer(port int) {
	mux := mux.NewRouter()
	mux.Handle("/healthcheck", negroni.New(
		negroni.HandlerFunc(logRequest),
		negroni.Wrap(get(http.HandlerFunc(handlers.Healthcheck))),
	))

	authenticationMiddleware := middleware.CreateMiddleware()
	mux.Handle("/info/me", negroni.New(
		negroni.HandlerFunc(logRequest),
		negroni.HandlerFunc(authenticationMiddleware.HandlerWithNext),
		negroni.Wrap(get(http.HandlerFunc(info.Me))),
	))
	http.Handle("/", mux)
	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
