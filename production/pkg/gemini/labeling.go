package gemini

import (
	"context"
	"fmt"
	"local-test/internal/model"
	"log"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func parseLabels(input string) []string {
	// Split the input by newline
	lines := strings.Split(input, "\n")

	var labels []string
	for _, line := range lines {
		// Remove the "- " prefix and trim whitespace
		line = strings.TrimSpace(strings.TrimPrefix(line, "- "))
		// Add non-empty labels to the list
		if line != "" {
			labels = append(labels, line)
		}
		// Stop if we already have 3 labels
		if len(labels) == 3 {
			break
		}
	}

	return labels
}

func generateCaptionPrompt(url string) string {
	return fmt.Sprintf("日本語で次の画像のキャプションを生成してください: %s", url)
}

func generateLabelingPrompt(content *string, code *model.Code, mediaCaption string) string {
	prompt := "次のツイート内容に以下のラベルから最大3つを選んでラベル付けしてください:\n"

	if content != nil {
		contentStr := *content
		prompt += fmt.Sprintf("内容: %s\n", contentStr)
	}
	if code != nil {
		codeStr := code.Content
		prompt += fmt.Sprintf("コード: %s\n", codeStr)
	}
	if mediaCaption != "" {
		prompt += fmt.Sprintf("メディアキャプション: %s\n", mediaCaption)
	}

	prompt += "ラベル:\n"
	labels := model.GetLabels()
	for _, label := range labels {
		prompt += fmt.Sprintf("- %s\n", label)
	}

	return prompt
}

func generateRelabelingPrompt(labels []string) string {
	prompt := fmt.Sprintf("The following %d label are invalid. Please select %d valid label that is close in meaning instead.:\n", len(labels), len(labels))

	for _, label := range labels {
		prompt += fmt.Sprintf("- %s\n", label)
	}

	prompt += "valid labels:\n"
	collectLabels := model.GetLabels()
	for _, label := range collectLabels {
		prompt += fmt.Sprintf("- %s\n", label)
	}

	return prompt
}

func LabelingTweet(ctx context.Context, content *string, code *model.Code, media *model.Media) ([]string, error) {
	// Media validation
	if media != nil {
		if err := media.Validate(); err != nil {
			return nil, err
		}
	}

	// Connect to generative AI client
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		return nil, err
	}
	defer client.Close()

	gptModel := client.GenerativeModel("gemini-pro")

	// Step 1: Generate caption for the media (image processing)
	var mediaCaption string
	if media != nil {
		if media.Type == "image" {
			captionPrompt := genai.Text(generateCaptionPrompt(media.URL))
			captionResp, err := gptModel.GenerateContent(ctx, captionPrompt)
			if err != nil {
				return nil, err
			}

			if len(captionResp.Candidates) > 0 && len(captionResp.Candidates[0].Content.Parts) > 0 {
				var parts []string
				for _, part := range captionResp.Candidates[0].Content.Parts {
					parts = append(parts, fmt.Sprintf("%s", part))
				}
				mediaCaption = strings.Join(parts, " ")
			} else {
				mediaCaption = "No caption generated"
			}
		} else {
			mediaCaption = "Media type not supported for caption generation"
		}
	}

	// Step 2: Combine content, code, and media caption for labeling
	labelingPrompt := genai.Text(generateLabelingPrompt(content, code, mediaCaption))
	labelingResp, err := gptModel.GenerateContent(ctx, labelingPrompt)
	if err != nil {
		return nil, err
	}

	// Parse the response into labels
	var labels string
	if len(labelingResp.Candidates) > 0 && len(labelingResp.Candidates[0].Content.Parts) > 0 {
		var parts []string
		for _, part := range labelingResp.Candidates[0].Content.Parts {
			parts = append(parts, fmt.Sprintf("%s", part))
		}
		labels = parts[0]
	}
	parsedLabels := parseLabels(labels)

	// Extract Invalid Labels
	var invalidLabels []string
	for _, label := range parsedLabels {
		temp := model.Label(label)
		if temp.Validate() != nil {
			invalidLabels = append(invalidLabels, label)
		}
	}

	// Relabeling
	if len(invalidLabels) > 0 {
		log.Println("Invalid labels detected, prompting for relabeling...")
		log.Printf("Invalid labels: %v", invalidLabels)
		// Generate relabeling prompt
		relabelingPrompt := genai.Text(generateRelabelingPrompt(invalidLabels))
		relabelingResp, err := gptModel.GenerateContent(ctx, relabelingPrompt)
		if err != nil {
			return nil, err
		}
		// Parse the response into labels
		var relabels string
		if len(relabelingResp.Candidates) > 0 && len(relabelingResp.Candidates[0].Content.Parts) > 0 {
			var parts []string
			for _, part := range relabelingResp.Candidates[0].Content.Parts {
				parts = append(parts, fmt.Sprintf("%s", part))
			}
			relabels = parts[0]
		}
		parsedRelabels := parseLabels(relabels)
		log.Printf("Relabeled: %v", parsedRelabels)

		// Replace invalid labels with valid labels
		for i, label := range parsedLabels {
			if model.Label(label).Validate() != nil {
				if model.Label(parsedRelabels[0]).Validate() == nil {
					parsedLabels[i] = parsedRelabels[0]
				}
				relabels = relabels[1:]
			}
		}
	}

	return parsedLabels, nil
}