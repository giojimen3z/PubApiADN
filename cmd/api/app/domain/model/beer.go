package model

type Beer struct {
	BeerId   int64   `json:"id" validate:"required"`
	Name     string  `json:"Name" validate:"required"`
	Brewery  string  `json:"Brewery" validate:"required"`
	Country  string  `json:"Country" validate:"required"`
	Price    float64 `json:"Price" validate:"required"`
	Currency string  `json:"Currency" validate:"required"`
}
