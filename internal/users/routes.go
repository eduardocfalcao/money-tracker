package users

import "github.com/go-chi/chi/v5"

type userRouter struct {
	handlers *Handlers
}

func NewRoutes(h *Handlers) *userRouter {
	return &userRouter{h}
}

func (u *userRouter) RegisterRoutes(router chi.Router, privateRoutes chi.Router) {
	router.Post("/oauth/token", u.handlers.Login)

	//private routes
	privateRoutes.Get("/users/me", u.handlers.Me)

}
