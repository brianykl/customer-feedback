package models

type Feedback struct {
	FeedbackID   string `gorm:"column:feedback_id"`
	Email        string `gorm:"column:email"`
	Category     string `gorm:"column:category"`
	FeedbackText string `gorm:"column:feedback_text"`
}
