package intents

import (
	"io/ioutil"
	"net/http"
	"strings"

	intent "github.com/go-skynet/intents/core/intent"
	client "github.com/go-skynet/llama-cli/client"
)

func NewWebGet(s string) *ScrapeWeb {
	return &ScrapeWeb{input: StringIntent(s)}
}

func Get() *ScrapeWeb {
	return &ScrapeWeb{}
}

type WebGet struct {
	input intent.IntentInput
}

func getURL(url string) (string, error) {
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

	return string(body), nil
}

func (s *WebGet) SetInput(ii intent.IntentInput) intent.IntentInput {
	s.input = ii
	return s
}

func (i *WebGet) Execute(c *client.Client, opts ...client.InputOption) (string, error) {
	var inputresult string
	var err error
	if i.input != nil {
		inputresult, err = i.input.Execute(c, opts...)
		if err != nil {
			return "", err
		}
	}

	return getURL(inputresult)
}
