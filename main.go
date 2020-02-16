package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	list := []WeightedItem{
		{"test_1", 10},
		{"test_2", 20},
		{"test_3", 0},
		{"test_4", 70},
	}

	item := GetItem(list)

	fmt.Println(item)
}

// GetItem takes in a weighted items list and returns a single item from the list
func GetItem(items []WeightedItem) interface{} {

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)

	total := sumWeights(items)
	weightRanges := generateWeightRanges(items, total)

	var selected interface{}

	r := random.Float64()
	for i := 0; i < len(items); i++ {
		if r <= weightRanges[i] {
			selected = items[i]
			break
		}
	}

	return selected
}

func sumWeights(items []WeightedItem) int {
	var total int = 0
	for i := 0; i < len(items); i++ {
		total += items[i].weight
	}

	return total
}

func generateWeightRanges(items []WeightedItem, total int) []float64 {
	var weightRanges []float64
	var count int = 0
	for _, item := range items {
		count += item.weight
		var weightRange float64 = -1.0
		if item.weight > 0 {
			weightRange = (100.0 / float64(total) * float64(count)) / 100
		}
		weightRanges = append(weightRanges, weightRange)
	}

	return weightRanges
}

// WeightedItem consists of an object and int representing the weight of chance to be returned
type WeightedItem struct {
	item   interface{}
	weight int
}
