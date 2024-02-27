package grpcclient

import (
	"context"
	"log"
	"time"

	pb "github.com/brianykl/customer-feedback/grpc/proto"
	"google.golang.org/grpc"
)

var (
	serverAddr = "localhost:50051" // Adjust the address to your gRPC server's address
)

// SentimentAnalysisClient is a client for communicating with the Sentiment Analysis service
type SentimentAnalysisClient struct {
	client pb.SentimentAnalysisClient
}

type TopicModellingClient struct {
	client pb.TopicModellingClient
}

// NewSentimentAnalysisClient creates a new client for Sentiment Analysis service
func NewSentimentAnalysisClient() *SentimentAnalysisClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", serverAddr, err)
	}

	return &SentimentAnalysisClient{
		client: pb.NewSentimentAnalysisClient(conn),
	}
}

// NewTopicModellingClient creates a new client for Topic Modelling service
func NewTopicModellingClient() *TopicModellingClient {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", serverAddr, err)
	}

	return &TopicModellingClient{
		client: pb.NewTopicModellingClient(conn),
	}
}

// AnalyzeText calls the AnalyzeText method of Sentiment Analysis service
func (c *SentimentAnalysisClient) AnalyzeText(text string) <-chan *pb.SentimentResponse {
	resultChan := make(chan *pb.SentimentResponse, 1)

	go func() {
		defer close(resultChan)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		response, err := c.client.AnalyzeText(ctx, &pb.TextRequest{Text: text})
		if err != nil {
			log.Printf("Error analyzing text for sentiment analysis: %v", err)
			return
		}
		resultChan <- response
	}()
	return resultChan
}

// AnalyzeText calls the AnalyzeText method of Topic Modelling service
func (c *TopicModellingClient) AnalyzeText(text string) <-chan *pb.TopicResponse {
	resultChan := make(chan *pb.TopicResponse, 1)

	go func() {
		defer close(resultChan)
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		response, err := c.client.AnalyzeText(ctx, &pb.TextRequest{Text: text})
		if err != nil {
			log.Printf("Error analyzing text for topic modelling: %v", err)
			return
		}
		resultChan <- response
	}()
	return resultChan
}
