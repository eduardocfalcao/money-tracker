package info

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/eduardocfalcao/money-tracker/http/auth/middleware/jwtParser"
)

type MeData struct {
	Name  string `json:name`
	Email string `json:email`
}

func Me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
	token := jwtParser.Parse(authHeaderParts[1])

	claims, _ := token.Claims.(*jwtParser.CustomClaims)

	_ = json.NewEncoder(w).Encode(claims)
}
