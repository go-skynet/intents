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
			ws := func() *intents.ScrapeWeb { return intents.NewWebScraper("https://kairos.io/community/") }
			chain.Add(ws())
			chain.Add(intents.Summarize())
			a, err := chain.Execute(c, client.WithTokens(99999))
			Expect(err).ToNot(HaveOccurred())
			// We can't really check the summary content as can be variadic, we just check that it's not the same page that we pass by
			str, err := ws().Execute(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(str).To(ContainSubstring("Kairos"))
			Expect(a).ToNot(Equal(str))
		})
	})
})
