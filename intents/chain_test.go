package intents_test

import (
	. "github.com/go-skynet/intents/intents"
	"github.com/go-skynet/llama-cli/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chains", func() {

	Context("Run", func() {
		It("should return a string and an error", func() {
			c := client.NewClient(apiAddress)
			chain := &Chain{}
			chain.Add(StringIntent("What's an alpaca?"))
			chain.Add(NewBaseIntent("{{.Input}}"))
			chain.Add(FindSubject())
			a, err := chain.Execute(c)
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(Equal("Alpaca"))
		})
		It("should return a string and an error", func() {
			c := client.NewClient(apiAddress)
			chain := &Chain{}
			chain.Add(NewBaseIntent("What's an alpaca?"))
			chain.Add(FindSubject())
			a, err := chain.Execute(c)
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(Equal("Alpaca"))
		})
	})
})
