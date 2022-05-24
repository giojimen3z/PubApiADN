package builder

import "github.com/PubApiADN/cmd/api/app/domain/model"

type BeerDataBuilder struct {
	beerId   int64
	name     string
	brewery  string
	country  string
	price    float64
	currency string
}

func NewBeerDataBuilder() *BeerDataBuilder {
	return &BeerDataBuilder{
		beerId:   1,
		name:     "Golden",
		brewery:  "Kross",
		country:  "Chile",
		price:    10.5,
		currency: "EUR",
	}
}

func NewBeerDataBuilderWithZeroID() *BeerDataBuilder {
	return &BeerDataBuilder{
		beerId:   0,
		name:     "",
		brewery:  "",
		country:  "",
		price:    0,
		currency: "",
	}
}
func (builder *BeerDataBuilder) Build() model.Beer {
	return model.Beer{
		BeerId:   builder.beerId,
		Name:     builder.name,
		Brewery:  builder.brewery,
		Country:  builder.country,
		Price:    builder.price,
		Currency: builder.currency,
	}
}
