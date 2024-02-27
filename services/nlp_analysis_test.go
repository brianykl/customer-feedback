package services

import (
	"testing"

	cohere "github.com/cohere-ai/cohere-go"
	"github.com/stretchr/testify/assert"
)

// TestSentimentAnalysis tests the SentimentAnalysis function with a mock Cohere client.
func TestSentimentAnalysis(t *testing.T) {
	// Mock response setup
	mockResponse := &cohere.ClassifyResponse{
		ID: "mock_response_id", // Simulate a response ID from Cohere
	}
	mockClient := &MockCohereClient{
		ClassifyResponse: mockResponse, // Use the mock response
		ClassifyErr:      nil,          // Simulate no error
	}

	// Test data
	feedbackTexts := []string{"This product is great!"}

	// Execute the function with the mock client and test data
	responseID, _ := SentimentAnalysis(mockClient, feedbackTexts)

	// Assertions
	assert.Equal(t, "mock_response_id", responseID, "The response ID should match the mock response ID.")
}

// TestTopicModelling tests the TopicModelling function with a mock Cohere client.
func TestTopicModelling(t *testing.T) {
	// Mock response setup
	mockGenerations := []cohere.Generation{{Text: "Sample topic modeling response"}}
	mockResponse := &cohere.GenerateResponse{
		Generations: mockGenerations, // Use the mock generation
	}
	mockClient := &MockCohereClient{
		GenerateResponse: mockResponse, // Use the mock response
		GenerateErr:      nil,          // Simulate no error
	}

	// Test data
	feedbackText := "What are the main themes discussed?"

	// Execute the function with the mock client and test data
	generations, _ := TopicModelling(mockClient, feedbackText)

	// Assertions
	assert.Equal(t, mockGenerations, generations, "The generations should match the mock generations.")
}
