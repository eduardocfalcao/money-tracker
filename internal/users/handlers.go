package users

import (
	"encoding/json"
	"errors"
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
	service service
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
		logrus.Warnf("[user handler] Login endpoint received a malformed json: %s", err)
		api.MalformedJsonResponse(w)
		return
	}

	user, err := h.service.GetUserByEmailAndPassword(r.Context(), loginRequest)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			api.WriteApiError(w, http.StatusUnauthorized, api.APIError{
				Message: "There is no user with the given email and password.",
			})
			return
		}
		logrus.Errorf("[user handler] Error retrieving user to login: %s", err)
		api.InternalErrorResponse(w)
		return
	}

	token, err := h.JWTService.CreateToken(auth.CreateTokenArgs{
		Username: user.Email,
		Email:    user.Email,
	})
	if err != nil {
		logrus.Errorf("[user handler] Error generating the user token: %s", err)
		api.InternalErrorResponse(w)
		return
	}

	api.JsonResponse(w, token)
}

func (h *Handlers) CreateUser(w http.ResponseWriter, r *http.Request) {

	var request models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		logrus.Warnf("[user handler] Create user endpoint received a malformed json: %s", err)
		api.MalformedJsonResponse(w)
		return
	}

	err := h.service.CreateUser(r.Context(), request)
	if err != nil {
		logrus.Errorf("[user handler] Error creating the user: %s", err)
		api.InternalErrorResponse(w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
