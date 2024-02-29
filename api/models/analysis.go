package models

type Analysis struct {
	FeedbackID string `gorm:"column:feedback_id"`
	Sentiment  string `gorm:"column:sentiment"`
	Topic      string `gorm:"column:topic"`
}

func NewAnalysis(feedback_id, sentiment, topic string) *Analysis {
	a := Analysis{
		FeedbackID: feedback_id,
		Sentiment:  sentiment,
		Topic:      topic,
	}
	return &a
}
