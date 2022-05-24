package beer_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/PubApiADN/cmd/api/app/domain/service/beer"
	"github.com/PubApiADN/cmd/api/test/builder"
)

const (
	zeroQuantity = 0
)

var _ = Describe("Service", func() {
	Context("Get Beer Box Price", func() {
		var (
			getBeerBoxPriceService beer.GetBeerBoxPrice
		)
		BeforeEach(func() {
			getBeerBoxPriceService = beer.GetBeerBoxPrice{}

		})

		AfterEach(func() {
			os.Clearenv()
		})
		When("a new valid request is received", func() {
			It("should return beer box price default", func() {

				beerBoxExpected := builder.NewBeerBoxDataBuilder().WithPrice(2.4966).Build()
				currencyConversion := builder.NewCurrencyConversionDataBuilder().Build()

				beerBox := getBeerBoxPriceService.GetBeerBoxPrice(zeroQuantity, currencyConversion)

				Expect(beerBoxExpected).Should(Equal(beerBox))
			})
		})
		When("a new valid request is received", func() {
			It("should return beer box price", func() {

				quantity := 3
				beerBoxExpected := builder.NewBeerBoxDataBuilder().WithPrice(7.489800000000001).Build()
				currencyConversion := builder.NewCurrencyConversionDataBuilder().Build()

				beerBox := getBeerBoxPriceService.GetBeerBoxPrice(int64(quantity), currencyConversion)

				Expect(beerBoxExpected).Should(Equal(beerBox))
			})
		})
	})
})
