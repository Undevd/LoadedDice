package loadeddice

import (
	"fmt"
	"testing"
)

func TestSumWeights(t *testing.T) {
	list100 := []WeightedFace{
		{"item_1", 10},
		{"item_2", 20},
		{"item_3", 0},
		{"item_4", 70},
	}

	list20 := []WeightedFace{
		{"item_1", 5},
		{"item_2", 5},
		{"item_3", 5},
		{"item_4", 5},
	}

	var cases = []struct {
		items []WeightedFace
		want  int
	}{
		{list100, 100},
		{list20, 20},
	}

	for _, tt := range cases {
		testname := fmt.Sprintf("expect %d", tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := sumWeights(tt.items)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
