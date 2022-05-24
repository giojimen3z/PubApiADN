package model

type CurrencyConversion struct {
	Currency Currency `json:"result"`
	Status   string   `json:"status"`
}
