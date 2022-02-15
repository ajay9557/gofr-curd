package product


type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Response struct {
	Data       interface{} `json:"product"`
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
}