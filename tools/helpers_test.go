package tools

import (
	"testing"
)

func TestScaleValue(t *testing.T) {
	actual1 := ScaleValue(2500, 0, 5000, 500, 1000)
	expected1 := uint16(750)
	if expected1 != actual1 {
		t.Fatalf("Actual value was %v; expected %v", actual1, expected1)
	}

	actual2 := ScaleValue(0, 0, 5000, 1000, 2000)
	expected2 := uint16(1000)
	if expected2 != actual2 {
		t.Fatalf("ScaleValue returned %v; expected %v", actual2, expected2)
	}
}
