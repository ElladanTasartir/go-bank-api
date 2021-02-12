package models

type Bank struct {
	ID   int
	Name string `json:"name"`
	Code int    `json:"code"`
}
