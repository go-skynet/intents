package intents

import (
	"fmt"

	intent "github.com/go-skynet/intents/core/intent"

	client "github.com/go-skynet/llama-cli/client"
)

type StringIntent string

func (s StringIntent) Execute(c *client.Client, opts ...client.InputOption) (string, error) {
	return string(s), nil
}

func (s StringIntent) SetInput(intent.IntentInput) intent.IntentInput {
	return s
}

func FindSubject() *intent.Intent {
	return intent.NewBaseIntent("Return only the subject of the following sentence and nothing else in singular form: {{.Input}}")
}

// There should be a "recursive" summarize that tries to cut inputs that exceed a token size.
// We should then try to summarize each batch, join them and summarize it again (there will be loss of information).
func Summarize() *intent.Intent {
	return intent.NewBaseIntent("Given the following input text write a short summary: {{.Input}}. Now write a short summary.")
}

func Question(q string) *intent.Intent {
	return intent.NewBaseIntent(fmt.Sprintf("Given the following input text: {{.Input}}. %s", q))
}

func Proofread() *intent.Intent {
	return intent.NewBaseIntent("Proofread, improving clarity and flow: {{.Input}}")
}

func ProofreadWithTone(tone string) *intent.Intent {
	return intent.NewBaseIntent(fmt.Sprintf("Proofread, improving clarity and flow as a %s: {{.Input}}", tone))
}
