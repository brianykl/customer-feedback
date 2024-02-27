package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	_, file, _, _ := runtime.Caller(0)
	// Adjust to go one level up to the root from the current file's directory
	appPath, err := filepath.Abs(filepath.Dir(filepath.Join(file, "../")))
	if err != nil {
		panic("Failed to compute app path: " + err.Error())
	}
	fmt.Println("Computed appPath:", appPath)
	beego.TestBeegoInit(appPath)
}

func TestCreateFeedback(t *testing.T) {
	feedback := struct {
		FeedbackID   string `json:"feedback_id"`
		Email        string `json:"email"`
		Category     string `json:"category"`
		FeedbackText string `json:"feedback_text"`
	}{
		FeedbackID:   "uniqueID",
		Email:        "test@example.com",
		Category:     "Service",
		FeedbackText: "This is a test feedback.",
	}
	body, err := json.Marshal(feedback)
	if err != nil {
		t.Fatalf("Marshaling feedback failed: %v", err)
	}
	req, err := http.NewRequest("POST", "/feedback", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Creating request failed: %v", err)
	}
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusCreated { // Adjust according to your handler's expected response
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
	// Implement additional checks on the response body if necessary
}
