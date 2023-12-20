package users

import "github.com/go-chi/chi/v5"

type userRouters struct {
	handlers *Handlers
}

func (u *userRouters) RegisterRoutes(r *chi.Mux) {
	r.Get("/users/me", u.handlers.Me)
}
