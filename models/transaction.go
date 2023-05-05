package models

type Transaction struct {
	Id          string `json:"id"`
	Customer_id string `json:"customer_id"`
	Menu        string `json:"menu"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Payment     string `json:"payment"`
	Total       int    `json:"total"`
	Created_at  string `json:"created_at"`
}
