package intent_test

import (
	. "github.com/go-skynet/intents/core/intent"
	intents "github.com/go-skynet/intents/intents"

	"github.com/go-skynet/llama-cli/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Intent", func() {
	Context("NewBaseIntent", func() {
		It("should return a new Intent", func() {
			Expect(NewBaseIntent("")).To(BeAssignableToTypeOf(&Intent{}))
		})
	})
	Context("NewIntent", func() {
		It("should return a new Intent", func() {
			Expect(New("")).To(BeAssignableToTypeOf(&Intent{}))
		})
	})
	Context("Run intent", func() {
		It("should be able to run a baseintent as an input for other", func() {
			c := client.NewClient(apiAddress)
			i := intents.FindSubject().SetInput(NewBaseIntent("{{.Input}}").SetInput(intents.StringIntent("What's an alpaca?")))
			a, err := i.Execute(c)
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(Equal("Alpaca"))
		})
	})
})
