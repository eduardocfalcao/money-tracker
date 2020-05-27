package http

import (
	"fmt"
	"net/http"

	"github.com/eduardocfalcao/money-tracker/http/handlers"
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

// InitHTTPServer  initializates a new http server listening the port provided
func InitHTTPServer(port int) {
	mux := http.NewServeMux()

	mux.Handle("/healthcheck", get(http.HandlerFunc(handlers.Healthcheck)))

	http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
