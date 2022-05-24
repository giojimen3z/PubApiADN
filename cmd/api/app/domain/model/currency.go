package model

type Currency struct {
	Updated  string  `json:"updated"`
	Source   string  `json:"source"`
	Target   string  `json:"target"`
	Value    float64 `json:"value"`
	Quantity float64 `json:"quantity"`
	Amount   float64 `json:"amount"`
}
