package intents_test

import (
	"os"

	. "github.com/go-skynet/intents/core/chain"
	intents "github.com/go-skynet/intents/intents"
	"github.com/go-skynet/llama-cli/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var _ = Describe("DOM intent", func() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	Context("DOM", func() {
		It("Can ask question to the dom", func() {
			c := client.NewClient(apiAddress)
			chain := &Chain{}
			chain.Add(intents.NewDOMBrowser("https://kairos.io/docs/", ".td-search--offline"))
			chain.Add(intents.Question("What's the json search url?"))
			a, err := chain.Execute(c, client.WithTokens(99999), client.WithTopK(10000))
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(ContainSubstring(".json"))
		})
	})
})
