package models_test

import (
	"fmt"
	"testing"

	"github.com/brianykl/customer-feedback/models"
)

func TestConnect(*testing.T) {
	fmt.Println(models.Connect())
}
