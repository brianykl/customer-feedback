package controllers

import (
	"encoding/json"
	"log"
	"sync"

	beego "github.com/beego/beego/v2/server/web"
	grpcclient "github.com/brianykl/customer-feedback/api/grpc-client"
	"github.com/brianykl/customer-feedback/api/models"
	pb "github.com/brianykl/customer-feedback/grpc/proto"
	"github.com/brianykl/customer-feedback/services"
)

type FeedbackController struct {
	beego.Controller
	SentimentAnalysisClient pb.SentimentAnalysisClient
	TopicModellingClient    pb.TopicModellingClient
}

type FeedbackPayload struct {
	FeedbackID   string `json:"FeedbackID"`
	Email        string `json:"Email"`
	Category     string `json:"Category"`
	FeedbackText string `json:"FeedbackText"`
}

func (c *FeedbackController) CreateFeedback() {
	var payload FeedbackPayload
	payload.FeedbackID = services.GenerateID()
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "bad request"}
		c.ServeJSON()
		return
	}
	log.Printf("Received payload: %+v\n", payload)
	feedback := models.NewFeedback(payload.FeedbackID, payload.Email, payload.Category, payload.FeedbackText)
	if err := models.Insert(feedback); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "failed to insert feedback"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.SetStatus(201) // HTTP 201 Created
	c.Data["json"] = map[string]string{"message": "feedback created successfully"}
	c.ServeJSON()

	var wg sync.WaitGroup
	wg.Add(2)

	var sentiment string
	var topic string
	var mu sync.Mutex

	go func() {
		defer wg.Done()
		result := processSentimentAnalysis(feedback.FeedbackText)
		mu.Lock()
		sentiment = result
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		result := processTopicModelling(feedback.FeedbackText)
		mu.Lock()
		topic = result
		mu.Unlock()
	}()

	wg.Wait()
	analysis := models.NewAnalysis("", sentiment, topic)
	if err := models.Insert(analysis); err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "failed to insert feedback"}
		c.ServeJSON()
		return
	}
}

func processSentimentAnalysis(feedbackText string) string {
	saClient := grpcclient.NewSentimentAnalysisClient()
	saResponseChan := saClient.AnalyzeText(feedbackText)

	// Process sentiment analysis response
	saResponse := <-saResponseChan
	if saResponse != nil {
		log.Printf("Sentiment Analysis Response: %v", saResponse)
		return saResponse.Sentiment
		// Update feedback with sentiment analysis result here
	}
	return ""
}

func processTopicModelling(feedbackText string) string {
	tmClient := grpcclient.NewTopicModellingClient()
	tmResponseChan := tmClient.AnalyzeText(feedbackText)

	// Process topic modeling response
	tmResponse := <-tmResponseChan
	if tmResponse != nil {
		log.Printf("Topic Modelling Response: %v", tmResponse)
		return tmResponse.Topic
	}
	return ""
}
