package models

type Response struct {
	Data       interface{} `json:"product"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
}
