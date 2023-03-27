package intents

import (
	"fmt"
	"strings"

	intent "github.com/go-skynet/intents/core/intent"
	client "github.com/go-skynet/llama-cli/client"
	"github.com/gocolly/colly/v2"
)

type DOMBrowser struct {
	input   intent.IntentInput
	segment string
}

func NewDOMBrowser(url, segment string) *DOMBrowser {
	return &DOMBrowser{input: StringIntent(url), segment: segment}
}

func fetchSegment(url, query string) (string, error) {
	// Instantiate default collector
	c := colly.NewCollector()
	html := ""
	// Extract comment
	c.OnHTML(query, func(e *colly.HTMLElement) {
		str, _ := e.DOM.Html()
		html = str
	})

	if err := c.Visit(url); err != nil {
		return "", err
	}
	fmt.Println(html)

	if html == "" {
		return "", fmt.Errorf("no matches")
	}
	f := strings.Fields(html)
	return fmt.Sprintf("This is the content of %s: %s", url, strings.Join(f, " ")), nil
}

func (s *DOMBrowser) SetInput(ii intent.IntentInput) intent.IntentInput {
	s.input = ii
	return s
}

func (i *DOMBrowser) Execute(c *client.Client, opts ...client.InputOption) (string, error) {
	var inputresult string
	var err error
	if i.input != nil {
		inputresult, err = i.input.Execute(c, opts...)
		if err != nil {
			return "", err
		}
	}

	return fetchSegment(inputresult, i.segment)
}
