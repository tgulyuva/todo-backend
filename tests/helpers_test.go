package tests

import (
	"testing"

	"github.com/tgulyuva/todo-backend/helpers"
)

func TestDoubleSpace(t *testing.T) {
	text := "test  value"
	result := helpers.CheckDoubleSpace(text)
	if result != true {
		t.Fail()
	}
}

func TestRemoveDoubleSpace(t *testing.T) {
	text := "my  test  data"
	expected := "my test data"
	result := helpers.RemoveDoubleSpace(text)

	if expected != result {
		t.Error("Expected value: " + text + " but return value" + result)
	}
}
