package models_test

import (
	"fmt"
	"testing"

	"github.com/brianykl/customer-feedback/models"
)

func TestConnect(*testing.T) {
	fmt.Println(models.Connect())
}

func TestInsert(t *testing.T) {
	f := models.NewFeedback("1", "brian.yikun.li@gmail.com", "Product", "testingtesting woohah")
	err := models.Insert(f)
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	t.Logf("Feedback text: %s", f.FeedbackText)
	fmt.Println(f.FeedbackText)
}
