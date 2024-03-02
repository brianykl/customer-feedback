package services

import (
	"log"

	cohere "github.com/cohere-ai/cohere-go"
)

var examples = []cohere.Example{
	{Text: "The order came 5 days early", Label: "positive"},
	{Text: "The item exceeded my expectations", Label: "positive"},
	{Text: "I ordered more for my friends", Label: "positive"},
	{Text: "I would buy this again", Label: "positive"},
	{Text: "I would recommend this to others", Label: "positive"},
	{Text: "The package was damaged", Label: "negative"},
	{Text: "The order is 5 days late", Label: "negative"},
	{Text: "The order was incorrect", Label: "negative"},
	{Text: "I want to return my item", Label: "negative"},
	{Text: "The item's material feels low quality", Label: "negative"},
	{Text: "The product was okay", Label: "neutral"},
	{Text: "I received five items in total", Label: "neutral"},
	{Text: "I bought it from the website", Label: "neutral"},
	{Text: "I used the product this morning", Label: "neutral"},
	{Text: "The product arrived yesterday", Label: "neutral"},
}

func Client(apiKey string) *cohere.Client {
	co, err := cohere.CreateClient(apiKey)
	if err != nil {
		log.Print(err)
		return nil
	}
	return co
}

type CohereClient interface {
	Classify(options cohere.ClassifyOptions) (*cohere.ClassifyResponse, error)
	Generate(options cohere.GenerateOptions) (*cohere.GenerateResponse, error)
}

type MockCohereClient struct {
	ClassifyResponse *cohere.ClassifyResponse
	ClassifyErr      error
	GenerateResponse *cohere.GenerateResponse
	GenerateErr      error
}

func (m *MockCohereClient) Classify(options cohere.ClassifyOptions) (*cohere.ClassifyResponse, error) {
	return m.ClassifyResponse, m.ClassifyErr
}

func (m *MockCohereClient) Generate(options cohere.GenerateOptions) (*cohere.GenerateResponse, error) {
	return m.GenerateResponse, m.GenerateErr
}

func SentimentAnalysis(client CohereClient, feedback_text []string) (string, error) {
	response, err := client.Classify(cohere.ClassifyOptions{
		Inputs:   feedback_text,
		Examples: examples,
	})

	if err != nil {
		log.Print(err)
		return "", err
	}

	sentiment := response.Classifications[0].Prediction
	return sentiment, nil
}

func TopicModelling(client CohereClient, feedback_text string) ([]cohere.Generation, error) {
	var max_tokens = uint(3)
	var prompt = "Please summarize this feedback in one-word: " + feedback_text
	response, err := client.Generate(cohere.GenerateOptions{
		Prompt:    prompt,
		MaxTokens: &max_tokens,
	})

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return response.Generations, nil
}
