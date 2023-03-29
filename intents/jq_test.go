package intents_test

import (
	"os"

	. "github.com/go-skynet/intents/core/chain"

	intents "github.com/go-skynet/intents/intents"
	client "github.com/go-skynet/llama-cli/client"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var _ = Describe("JQ intent", func() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	Context("JQ", func() {
		It("Can filter json", func() {
			chain := &Chain{}
			chain.Add(intents.NewWebGet("https://mysafeinfo.com/api/data?list=englishmonarchs&format=json"))
			chain.Add(intents.NewJQ(".[0].Name"))
			a, err := chain.Execute(nil, client.WithTokens(99999))
			Expect(err).ToNot(HaveOccurred())
			Expect(a).To(Equal("Edward the Elder"))
		})
	})
})
