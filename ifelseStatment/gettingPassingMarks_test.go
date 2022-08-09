package ifelseStatment

import (
	"testing"
)

// test function
func TestPassingMarks(t *testing.T) {
	marks := 20
	want := "pass"

	result := gettingPassingMarks(marks)
	if result != want {
		t.Errorf("passing marks should be more than 33")
	}
}
