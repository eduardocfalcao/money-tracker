package models

type (
	Me struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
