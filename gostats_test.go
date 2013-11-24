package numgo

import (
	"testing"
)

func TestMean(t *testing.T) {
    input := []int64{32, 46, 12, 43, 36}

	if mean := Mean(input); mean != 33.8 {
		t.Errorf("expected mean 33.8, got %v", mean)
	}
}

func TestAverageDeviation(t *testing.T) {
    input := []int64{32, 46, 12, 43, 36}

	if sd := standardDeviation(input); sd - 11.973303637 > 0.0001 {
		t.Errorf("expected standard deviation 11.973303637, got %v", sd)
	}
}
