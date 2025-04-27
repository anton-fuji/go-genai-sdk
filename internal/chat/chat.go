package chat

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"github.com/google/generative-ai-go/genai"
)

type Chat struct {
	chat *genai.ChatSession
}

func New(ctx context.Context, client *genai.Client, model string, temp float32) (*Chat, error) {
	mdl := client.GenerativeModel(model)
	// cfg := &genai.GenerationConfig{Temperature: genai.Ptr(temp)}
	mdl.GenerationConfig = genai.GenerationConfig{
		Temperature: genai.Ptr(temp),
	}

	cht := mdl.StartChat()
	return &Chat{chat: cht}, nil
}

func (c *Chat) Run(ctx context.Context, in io.Reader, out io.Writer) error {
	sc := bufio.NewScanner(in)
	fmt.Fprint(out, ">>>")

	for sc.Scan() {
		user := sc.Text()
		resp, err := c.chat.SendMessage(ctx, genai.Text(user))
		if err != nil {
			fmt.Fprintln(out, "error:", err)
		} else if len(resp.Candidates) > 0 {
			for _, p := range resp.Candidates[0].Content.Parts {
				if txt, ok := p.(genai.Text); ok {
					fmt.Fprintln(out, string(txt))
				}
			}
		}
		fmt.Fprint(out, ">>>")
	}
	return sc.Err()
}
