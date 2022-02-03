package product


type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type HttpResponse struct {
	Message string `json:"message"`
	StatusCode int `json:"statusCode"`
}

type HttpErr struct {
	ErrorMsg string `json:"errMsg"`
	ErrorCode int `json:"errCode"`
}
