package auth

import (
	"net/http"

	"github.com/eduardocfalcao/money-tracker/internal/api"
)

const AuthHeader = "Authorization"

func (j *service) VerifyTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(AuthHeader)
		if header == "" {
			api.UnauthorizedResponse(w)
			return
		}
		tokenString := header[len("Bearer "):]

		err := j.VerifyToken(tokenString)
		if err != nil {
			api.UnauthorizedResponse(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}
