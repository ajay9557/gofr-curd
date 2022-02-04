package models

type Product struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Response struct {
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message,omitempty"`
	StatusCode int         `json:"statusCode,omitempty"`
}

type ErrorResponse struct {
	StatusCode   int    `json:"statusCode"`
	ErrorMessage string `json:"error"`
}
