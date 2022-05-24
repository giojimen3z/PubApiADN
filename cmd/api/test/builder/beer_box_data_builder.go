package builder

import "github.com/PubApiADN/cmd/api/app/domain/model"

type BeerBoxDataBuilder struct {
	price float64
}

func NewBeerBoxDataBuilder() *BeerBoxDataBuilder {
	return &BeerBoxDataBuilder{
		price: 2.4966,
	}
}

func (builder *BeerBoxDataBuilder) WithPrice(price float64) *BeerBoxDataBuilder {
	builder.price = price
	return builder
}
func (builder *BeerBoxDataBuilder) Build() model.BeerBox {
	return model.BeerBox{
		Price: builder.price,
	}
}
