package loadeddice

import (
	"math/rand"
	"time"
)

// WeightedFace consists of an object and int representing the weight of chance to be returned
type WeightedFace struct {
	Item   interface{}
	Weight int
}

// Roll takes in a weighted items list and returns a single item from the list
func Roll(items []WeightedFace) interface{} {

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	total := sumWeights(items)
	weightRanges := generateWeightRanges(items, total)

	var selected interface{}

	r := random.Float64()
	for i := 0; i < len(items); i++ {
		if r <= weightRanges[i] {
			selected = items[i].Item
			break
		}
	}

	return selected
}

func sumWeights(items []WeightedFace) int {
	var total int = 0
	for i := 0; i < len(items); i++ {
		total += items[i].Weight
	}

	return total
}

func generateWeightRanges(items []WeightedFace, total int) []float64 {
	var weightRanges []float64
	var count int = 0
	for _, item := range items {
		count += item.Weight
		var weightRange float64 = -1.0
		if item.Weight > 0 {
			weightRange = (100.0 / float64(total) * float64(count)) / 100
		}
		weightRanges = append(weightRanges, weightRange)
	}

	return weightRanges
}
