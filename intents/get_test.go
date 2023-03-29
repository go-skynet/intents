package intents_test

import (
	intents "github.com/go-skynet/intents/intents"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Get intent", func() {
	Context("retrieve pages", func() {
		It("should return a web page summary", func() {
			str, err := intents.NewWebGet("https://mysafeinfo.com/api/data?list=englishmonarchs&format=json").Execute(nil)
			Expect(err).ToNot(HaveOccurred())
			Expect(str).To(ContainSubstring("Edward the Elder"))
		})
	})
})
