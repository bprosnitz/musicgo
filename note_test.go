package musicgo

import (
	"testing"
)

func TestNormalize(t *testing.T) {
	type expectedResult struct {
		input  Note
		output Note
	}
	tests := []expectedResult{
		expectedResult{0, 0},
		expectedResult{12, 0},
		expectedResult{1, 1},
		expectedResult{-1, 11},
		expectedResult{-12, 0},
		expectedResult{-13, 11},
		expectedResult{13, 1},
		expectedResult{0.5, 0.5},
		expectedResult{12.5, 0.5},
		expectedResult{-0.5, 11.5},
	}
	for _, test := range tests {
		if normalize(test.input) != test.output {
			t.Errorf("Unexpected normalized value: %f. Expected: %f.", test.input, test.output)
		}
	}
}
