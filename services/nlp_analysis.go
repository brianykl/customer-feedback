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
	co, err := cohere.CreateClient("<<apiKey>>")
	if err != nil {
		log.Print(err)
		return nil
	}
	return co
}

func SentimentAnalysis(feedback_text []string) string {
	co := Client("APIkey")
	response, err := co.Classify(cohere.ClassifyOptions{
		Inputs:   feedback_text,
		Examples: examples,
	})

	if err != nil {
		log.Print(err)
		return "error"
	}

	return response.ID
}

func TopicModelling(feedback_text string) []cohere.Generation {
	co := Client("APIkey")
	var max_tokens uint
	max_tokens = 10
	response, err := co.Generate(cohere.GenerateOptions{
		Model:     "large", // You can choose from different model sizes
		Prompt:    feedback_text,
		MaxTokens: &max_tokens,
	})

	if err != nil {
		log.Print(err)
		return nil
	}

	return response.Generations
}
