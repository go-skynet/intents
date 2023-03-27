package chain_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var apiAddress string = os.Getenv("API_ADDRESS")

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "chains test Suite")
}
