package intents_test

import (
	. "github.com/go-skynet/intents/core/chain"
	intents "github.com/go-skynet/intents/intents"

	"github.com/go-skynet/llama-cli/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Webpage intent", func() {
	Context("summarize", func() {
		It("should return a web page summary", func() {
			c := client.NewClient(apiAddress)
			chain := &Chain{}
			chain.Add(intents.NewWebScraper("https://kairos.io/community/"))
			chain.Add(intents.Summarize())
			a, err := chain.Execute(c, client.WithTokens(99999))
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(ContainSubstring("Kairos is a community-driven project"))
		})
	})
})
