package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

func main() {

	var fWeight, mWeight, nWeight, errorMargin int

	flag.IntVar(&fWeight, "f", 146, "F weight")
	flag.IntVar(&mWeight, "m", 162, "M weight")
	flag.IntVar(&nWeight, "n", 203, "N weight")
	flag.IntVar(&errorMargin, "err", 1, "Error margin for the total weight")
	flag.Parse()

	totalWeight, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println("Usage: glycan [-f number] [-m number] [-n number] [-err number] weight")
		return
	}

	for adjustment := errorMargin * -1; adjustment <= errorMargin; adjustment++ {
		testWeight := totalWeight + adjustment
		fmt.Printf("Testing value %d (Error of %+d):\n", testWeight, adjustment)
		fmt.Printf("F\tM\tN\n")
		fmt.Printf("---\t---\t---\n")

		maxIterations := int(math.Ceil(float64(testWeight) / float64(min(fWeight, mWeight, nWeight))))

		for f := 0; f <= maxIterations; f++ {
			fComboWeight := f * fWeight

			for m := 0; m <= maxIterations; m++ {
				mComboWeight := m * mWeight

				for n := 0; n <= maxIterations; n++ {
					nComboWeight := n * nWeight

					if fComboWeight+mComboWeight+nComboWeight == testWeight {
						fmt.Printf("%d\t%d\t%d\n", f, m, n)
					}
				}
			}
		}

		fmt.Printf("\n")
	}

}

func min(nums ...int) int {
	smallest := nums[0]

	for _, num := range nums {
		if num < smallest {
			smallest = num
		}
	}

	return smallest
}
