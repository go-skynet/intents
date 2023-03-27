package intents_test

import (
	. "github.com/go-skynet/intents/core/chain"
	intents "github.com/go-skynet/intents/intents"

	"github.com/go-skynet/llama-cli/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("String intents", func() {
	Context("Question and summarization", func() {
		It("should be able to ask a question on a web page", func() {
			c := client.NewClient(apiAddress)
			chain := &Chain{}
			chain.Add(intents.NewWebScraper("https://kairos.io/community/"))
			chain.Add(intents.Summarize())
			chain.Add(intents.Question("Does kairos have office hours?"))
			a, err := chain.Execute(c, client.WithTokens(99999), client.WithTopK(10000))
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(ContainSubstring("Yes"))
		})
	})
})
