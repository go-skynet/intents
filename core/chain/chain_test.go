package chain_test

import (
	. "github.com/go-skynet/intents/core/chain"
	"github.com/go-skynet/intents/core/intent"
	intents "github.com/go-skynet/intents/intents"

	"github.com/go-skynet/llama-cli/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Chains", func() {

	Context("Run", func() {
		It("should return a string and an error", func() {
			c := client.NewClient(apiAddress)
			chain := &Chain{}
			chain.Add(intents.StringIntent("What's an alpaca?"))
			chain.Add(intent.NewBaseIntent("{{.Input}}"))
			chain.Add(intents.FindSubject())
			a, err := chain.Execute(c)
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(Equal("Alpaca"))
		})
		It("should return a string and an error", func() {
			c := client.NewClient(apiAddress)
			chain := &Chain{}
			chain.Add(intent.NewBaseIntent("What's an alpaca?"))
			chain.Add(intents.FindSubject())
			a, err := chain.Execute(c)
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(Equal("Alpaca"))
		})

		It("should able to mix chains", func() {
			c := client.NewClient(apiAddress)
			chain := &Chain{}
			chain.Add(intent.NewBaseIntent("What's an alpaca?"))
			chain.Add(intents.Summarize())

			chain2 := &Chain{}
			chain2.Add(chain)
			chain2.Add(intents.FindSubject())

			a, err := chain2.Execute(c)
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(Equal("Alpaca"))
		})

		It("should able to mix chains", func() {
			c := client.NewClient(apiAddress)
			chain := &Chain{}
			chain.Add(intent.NewBaseIntent("What's a robot?"))
			chain.Add(intents.Summarize())

			chain2 := &Chain{}
			chain2.Add(intent.NewBaseIntent("What's an alpaca?"))
			chain2.Add(chain)
			chain2.Add(intents.FindSubject())

			a, err := chain2.Execute(c)
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(Equal("Alpaca"))
		})
	})
})
