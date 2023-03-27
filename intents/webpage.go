package intents

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	intent "github.com/go-skynet/intents/core/intent"
	client "github.com/go-skynet/llama-cli/client"
)

func NewWebScraper(s string) *ScrapeWeb {
	return &ScrapeWeb{input: StringIntent(s)}
}

func Scrape() *ScrapeWeb {
	return &ScrapeWeb{}
}

type ScrapeWeb struct {
	input intent.IntentInput
}

func fetchURL(url string) (string, error) {
	// Make sure the URL starts with http:// or https://
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	// Make a GET request to the URL
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	p := strings.NewReader(string(body))
	doc, _ := goquery.NewDocumentFromReader(p)

	for _, t := range []string{"style", "script", "nav"} {
		doc.Find(t).Each(func(i int, el *goquery.Selection) {
			el.Remove()
		})
	}

	// Reduce multiple whitespaces
	f := strings.Fields(doc.Text())
	return strings.Join(f, " "), nil
}

func (s *ScrapeWeb) SetInput(ii intent.IntentInput) intent.IntentInput {
	s.input = ii
	return s
}

func (i *ScrapeWeb) Execute(c *client.Client, opts ...client.InputOption) (string, error) {
	var inputresult string
	var err error
	if i.input != nil {
		inputresult, err = i.input.Execute(c, opts...)
		if err != nil {
			return "", err
		}
	}

	return fetchURL(inputresult)
}
