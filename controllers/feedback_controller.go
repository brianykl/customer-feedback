package controllers

import (
	"encoding/json"
	"log"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/brianykl/customer-feedback/models"
)

type FeedbackController struct {
	beego.Controller
}

type FeedbackPayload struct {
	FeedbackID   string `json:"FeedbackID"`
	Email        string `json:"Email"`
	Category     string `json:"Category"`
	FeedbackText string `json:"FeedbackText"`
}

func (c *FeedbackController) CreateFeedback() {
	var payload FeedbackPayload
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

}
