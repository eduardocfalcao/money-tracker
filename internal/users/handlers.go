package users

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/eduardocfalcao/money-tracker/internal/api"
	"github.com/eduardocfalcao/money-tracker/internal/auth"
	"github.com/eduardocfalcao/money-tracker/internal/users/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type Handlers struct {
	auth.JWTService
}

func NewHandler(jwtService auth.JWTService) *Handlers {
	return &Handlers{JWTService: jwtService}
}

func (u *Handlers) Me(w http.ResponseWriter, r *http.Request) {
	authHeaderParts := strings.Split(r.Header.Get("Authorization"), " ")
	token, err := u.JWTService.ParseToken(authHeaderParts[1])
	if err != nil {
		logrus.Errorf("[user handler] Error parsing the user token: %s", err)
		api.InternalErrorResponse(w)
		return
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	api.JsonResponse(w, claims)
}

// temp soliution: The final solution must be an oauth kind authentication,
// with better security schema
func (h *Handlers) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var loginRequest models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		logrus.Warnf("[user handler] Login request endpoint received a malformed json: %s", err)
		api.MalformedJsonResponse(w)
	}

	if loginRequest.Email == "eduardo.cfalcao@gmail.com" && loginRequest.Password == "123" {
		token, err := h.JWTService.CreateToken(auth.CreateTokenArgs{
			Username: loginRequest.Email,
			Email:    loginRequest.Email,
		})
		if err != nil {
			logrus.Errorf("[user handler] Error generating the user token: %s", err)
			api.InternalErrorResponse(w)
			return
		}

		api.JsonResponse(w, token)
	} else {
		api.WriteApiError(w, http.StatusUnauthorized, api.APIError{
			Message: "There is no user with the given email and password.",
		})
	}
}
