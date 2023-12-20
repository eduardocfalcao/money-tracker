package users

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/eduardocfalcao/money-tracker/internal/auth/middleware/jwtParser"
)

type Handlers struct{}

func NewHandler() *Handlers {
	return &Handlers{}
}

type MeData struct {
	Name  string `json:name`
	Email string `json:email`
}

func (u *Handlers) Me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
	token := jwtParser.Parse(authHeaderParts[1])

	claims, _ := token.Claims.(*jwtParser.CustomClaims)

	_ = json.NewEncoder(w).Encode(claims)
}
