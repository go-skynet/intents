package intents

import client "github.com/go-skynet/llama-cli/client"

type StringIntent string

func (s StringIntent) Execute(c *client.Client, opts ...client.InputOption) (string, error) {
	return string(s), nil
}

func (s StringIntent) SetInput(IntentInput) IntentInput {
	return s
}

func FindSubject() *Intent {
	return &Intent{
		template: baseTemplate("Return only the subject of the following sentence and nothing else in singular form: {{.Input}}"),
	}
}
