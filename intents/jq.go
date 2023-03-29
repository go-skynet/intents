package intents

import (
	"encoding/json"
	"fmt"
	"log"

	intent "github.com/go-skynet/intents/core/intent"
	client "github.com/go-skynet/llama-cli/client"
	"github.com/itchyny/gojq"
)

func jq(q string, a any) (string, error) {
	query, err := gojq.Parse(q)
	if err != nil {
		return "", err
	}
	iter := query.Run(a) // or query.RunWithContext
	res := ""
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return "", err
		}
		res += fmt.Sprint(v)
	}

	return res, nil
}

type JQ struct {
	input   intent.IntentInput
	segment string
}

func NewJQ(segment string) *JQ {
	return &JQ{segment: segment}
}

func (s *JQ) SetInput(ii intent.IntentInput) intent.IntentInput {
	s.input = ii
	return s
}

func (i *JQ) Execute(c *client.Client, opts ...client.InputOption) (string, error) {
	input := []any{}

	var inputresult string
	var err error
	if i.input != nil {
		inputresult, err = i.input.Execute(c, opts...)
		if err != nil {
			return "", err
		}
	}
	log.Print("Unmarshallling", inputresult)
	if err := json.Unmarshal([]byte(inputresult), &input); err != nil {
		input := map[string]any{}
		if err := json.Unmarshal([]byte(inputresult), &input); err != nil {
			return "", err
		}
	}

	return jq(i.segment, input)
}
