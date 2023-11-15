package utils

import (
	"context"

	"os"

	gl "cloud.google.com/go/ai/generativelanguage/apiv1beta2"
	pb "cloud.google.com/go/ai/generativelanguage/apiv1beta2/generativelanguagepb"

	"google.golang.org/api/option"
)

func AIGenearateText(message string) (string, error) {

	ctx := context.Background()
	client, err := gl.NewTextRESTClient(ctx, option.WithAPIKey(os.Getenv("PALM_KEY")))
	if err != nil {
		return "", err
	}

	defer client.Close()
	req := &pb.GenerateTextRequest{
		Model: "models/text-bison-001",
		Prompt: &pb.TextPrompt{
			Text: message,
		},
	}

	resp, err := client.GenerateText(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.Candidates[0].Output, nil

}
