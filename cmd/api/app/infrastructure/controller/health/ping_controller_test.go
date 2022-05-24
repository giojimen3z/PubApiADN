package health

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ping Controller", func() {

	Context("Ping", func() {

		When("ping controller is ok ", func() {
			It("should return  pong", func() {

				response := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(response)
				c.Request, _ = http.NewRequest(http.MethodGet, "/ping", nil)

				PingController.Ping(c)

				Expect(http.StatusOK).Should(Equal(response.Code))
				Expect("pong").Should(Equal(response.Body.String()))
			})
		})

	})

})
