package config

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Data Base client", func() {

	Context("data base client", func() {
		BeforeEach(func() {
			_ = os.Setenv("SCOPE", "local")
		})
		AfterEach(func() {
			os.Clearenv()

		})

		When("Get Write Connection", func() {
			It("should return  OK", func() {

				con, _ := GetWriteConnection()

				Expect(con).Should(Not(Equal(nil)))

			})
		})
		When("Get Read Connection", func() {
			It("should return  OK", func() {
				con, _ := GetReadConnection()
				Expect(con).Should(Not(Equal(nil)))
			})
		})

	})

})
