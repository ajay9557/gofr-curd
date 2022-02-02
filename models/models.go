package models

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}
