package api

type APIError struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type ValidationError struct {
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (v ValidationError) Error() string {
	return v.Message
}
