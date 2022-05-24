package config

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Environments", func() {

	Context("Scopes", func() {
		AfterEach(func() {
			os.Clearenv()

		})

		When("Is Productive Scope ", func() {
			It("should return  true", func() {

				_ = os.Setenv(scope, writeScope)
				defer os.Clearenv()
				got := isInProductiveScopes()

				Expect(true).Should(Equal(got))
			})
		})
		When("Isn´t Productive Scope ", func() {
			It("should return  false", func() {

				_ = os.Setenv(scope, "test")
				got := isInProductiveScopes()

				Expect(false).Should(Equal(got))
			})
		})

		When("Is Local Scope ", func() {
			It("should return  true", func() {

				_ = os.Setenv(scope, localScope)
				got := IsLocalScope()

				Expect(true).Should(Equal(got))
			})
		})
		When("Isn´t Local Scope ", func() {
			It("should return  false", func() {

				_ = os.Setenv(scope, "test")
				got := IsLocalScope()

				Expect(false).Should(Equal(got))
			})
		})
		When("try get currency apikey ", func() {
			It("should apikey", func() {

				apikeyExpected := "6392|h_2OeBxS2ibfZ^D1cA1o_3cYBQNUD*Pm"
				got := GetCurrencyApiKey()

				Expect(apikeyExpected).Should(Equal(got))
			})
		})
	})

})
