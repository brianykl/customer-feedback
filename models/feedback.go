package models

type Feedback struct {
	FeedbackID   string `gorm:"column:feedback_id"`
	Email        string `gorm:"column:email"`
	Category     string `gorm:"column:category"`
	FeedbackText string `gorm:"column:feedback_text"`
}

func NewFeedback(feedback_id, email, category, feedback_text string) *Feedback {
	f := Feedback{
		FeedbackID:   feedback_id,
		Email:        email,
		Category:     category,
		FeedbackText: feedback_text,
	}

	return &f
}
