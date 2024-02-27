package controllers

import (
	"encoding/json"
	"log"

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

	go func() {
		saClient := grpcclient.NewSentimentAnalysisClient()
		saResponseChan := saClient.AnalyzeText(feedback.FeedbackText)
		tmClient := grpcclient.NewTopicModellingClient()
		tmResponseChan := tmClient.AnalyzeText(feedback.FeedbackText)

		// Process sentiment analysis and topic modeling responses as they arrive
		// Note: This is for demonstration; adapt logging, error handling, and response processing as needed
		for i := 0; i < 2; i++ { // Expecting two responses
			select {
			case saResponse := <-saResponseChan:
				if saResponse != nil {
					log.Printf("Sentiment Analysis Response: %v", saResponse)
					// Update feedback with sentiment analysis result here
				}
			case tmResponse := <-tmResponseChan:
				if tmResponse != nil {
					log.Printf("Topic Modelling Response: %v", tmResponse)
					// Update feedback with topic modeling result here
				} else {
					log.Printf("BOOBA")
				}
			}
			log.Printf("ITERATED")
		}
	}()
}
