package intents_test

import (
	. "github.com/go-skynet/intents/intents"
	"github.com/go-skynet/llama-cli/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Intents", func() {
	Context("NewBaseIntent", func() {
		It("should return a new Intent", func() {
			Expect(NewBaseIntent("")).To(BeAssignableToTypeOf(&Intent{}))
		})
	})
	Context("NewIntent", func() {
		It("should return a new Intent", func() {
			Expect(NewIntent("")).To(BeAssignableToTypeOf(&Intent{}))
		})
	})
	Context("Run", func() {
		It("should return a string and an error", func() {
			c := client.NewClient(apiAddress)
			i := FindSubject().SetInput(NewBaseIntent("{{.Input}}").SetInput(StringIntent("What's an alpaca?")))
			a, err := i.Execute(c)
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(Equal("Alpaca"))
		})
	})
})
