package models

type Product struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Size        string  `json:"size"`
	Price       float64 `json:"price"`
}
