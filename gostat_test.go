package gostat

import (
	"testing"
)

func TestTakeExtremes(t *testing.T) {
    input := []int64{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
    correct := []int64{2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19}

    result := TakeExtremes(input)
    if len(result) != 18 {
        t.Errorf("expected len == 18, got %v", len(result))
    }

    for i, _ := range correct {
        if correct[i] != result[i] {
            t.Errorf("expected %v, got %v", correct, result)
        }
    }
}


func TestMean(t *testing.T) {
    input := []int64{32, 46, 12, 43, 36}

	if mean := Mean(input); mean != 33.8 {
		t.Errorf("expected mean 33.8, got %v", mean)
	}
}

func TestAverageDeviation(t *testing.T) {
    input := []int64{32, 46, 12, 43, 36}

	if sd := StandardDeviation(input); sd - 11.973303637 > 0.0001 {
		t.Errorf("expected standard deviation 11.973303637, got %v", sd)
	}
}
