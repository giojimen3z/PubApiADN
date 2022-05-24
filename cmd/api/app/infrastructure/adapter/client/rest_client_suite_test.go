package client

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRestClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rest client")
}
