package models

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Response struct {
	Product    interface{} `json:"product"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
}
