package models

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	fmt.Println(Connect())
}

func TestInsert(t *testing.T) {
	f := NewFeedback("1", "brian.yikun.li@gmail.com", "Product", "testingtesting woohah")
	err := Insert(f)
	if err != nil {
		t.Errorf("Insert failed: %v", err)
	}
	t.Logf("Feedback text: %s", f.FeedbackText)
	fmt.Println(f.FeedbackText)
}
