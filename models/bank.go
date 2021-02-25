package models

type Bank struct {
	Base
	Name string `json:"name"`
	Code int    `json:"code"`
}
