package intents

import (
	"bytes"
	"text/template"

	client "github.com/go-skynet/llama-cli/client"
)

const baseIntent = `
Below is an instruction that describes a task. Write a response that appropriately completes the request. 

### Instruction:

{{.Input}}

### Response:`

func baseTemplate(s string) string {
	str, _ := templateString(baseIntent, struct {
		Input string
	}{Input: s})
	return str
}

func NewBaseIntent(s string) *Intent {
	return &Intent{template: baseTemplate(s)}
}

func NewIntent(template string) *Intent {
	return &Intent{
		template: template,
	}
}

type IntentInput interface {
	Execute(c *client.Client, opts ...client.InputOption) (string, error)
	SetInput(IntentInput) IntentInput
}

type Intent struct {
	input    IntentInput
	template string
}

func (i *Intent) SetInput(ii IntentInput) IntentInput {
	i.input = ii
	return i
}

func (i *Intent) Execute(c *client.Client, opts ...client.InputOption) (string, error) {
	var inputresult string
	var err error
	if i.input != nil {
		inputresult, err = i.input.Execute(c, opts...)
		if err != nil {
			return "", err
		}
	}
	str, err := templateString(i.template, struct {
		Input string
	}{Input: inputresult})
	if err != nil {
		return "", err
	}

	return c.Predict(str, opts...)
}

func templateString(t string, in interface{}) (string, error) {
	// Parse the template
	tmpl, err := template.New("prompt").Parse(t)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, in)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
