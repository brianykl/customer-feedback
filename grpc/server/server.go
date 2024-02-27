package main

import (
	"context"
	"log"
	"net"

	pb "github.com/brianykl/customer-feedback/grpc/proto"
	services "github.com/brianykl/customer-feedback/services"
	cohere "github.com/cohere-ai/cohere-go"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cc := services.Client("MYAPIKEY")
	s := grpc.NewServer()

	saServer := &SentimentAnalysisServer{CohereClient: *cc}
	tmServer := &TopicModellingServer{CohereClient: *cc}

	pb.RegisterSentimentAnalysisServer(s, saServer)
	pb.RegisterTopicModellingServer(s, tmServer)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type SentimentAnalysisServer struct {
	pb.UnimplementedSentimentAnalysisServer
	CohereClient cohere.Client
}

type TopicModellingServer struct {
	pb.UnimplementedTopicModellingServer
	CohereClient cohere.Client
}

func (s *SentimentAnalysisServer) AnalyzeText(ctx context.Context, in *pb.TextRequest) (*pb.SentimentResponse, error) {

	sentiment, err := services.SentimentAnalysis(&s.CohereClient, []string{in.Text})
	if err != nil {
		log.Printf("Error processing Sentiment Analysis: %v", err)
		return nil, err // Ensure to return the error to the client.
	}
	log.Println(sentiment)
	return &pb.SentimentResponse{Sentiment: sentiment}, nil
}

func (s *TopicModellingServer) AnalyzeText(ctx context.Context, in *pb.TextRequest) (*pb.TopicResponse, error) {

	topic, err := services.TopicModelling(&s.CohereClient, in.Text)

	if err != nil {
		log.Printf("Error processing Topic Modelling: %v", err)
		return nil, err // Ensure to return the error to the client.
	}
	log.Printf(topic[0].Text)
	return &pb.TopicResponse{Topic: topic[0].Text}, nil
}
