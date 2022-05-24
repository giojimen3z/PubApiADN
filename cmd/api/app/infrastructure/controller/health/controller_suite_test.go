package health

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestControllerHealth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller")
}
