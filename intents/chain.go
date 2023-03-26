// TODO: a chain should Add(Intents)
// TODO: When it adds, it sets as input the previous one (Interface must implement SetInput)
// When executes pass the context (client, and options)
// Executes the intent

package intents

import client "github.com/go-skynet/llama-cli/client"

type Chain struct {
	i IntentInput
}

func (c *Chain) Add(i IntentInput) *Chain {
	if c.i != nil {
		i.SetInput(c.i)
	}
	c.i = i
	return c
}

func (cc *Chain) Execute(c *client.Client, opts ...client.InputOption) (string, error) {
	return cc.i.Execute(c, opts...)
}
