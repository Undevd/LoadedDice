# Loaded Dice

A utility which provides the ability to select a random item from a list based on a weight.

Usage:

```bash
go get github.com/undevd/loadeddice
```

```go
package main

import (
	"fmt"

	"github.com/undevd/loadeddice"
)

func main() {

	faces := []loadeddice.WeightedFace{
		{"Epic Fail", 1},
		{"Miss", 12},
		{"Hit", 6},
		{"Crit", 1},
	}

	dice := loadeddice.NewDice(faces)
	fmt.Println(dice.Roll().(string))
}
```