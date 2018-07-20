package main

import (
	"flag"
	"fmt"
	"math"
	"strconv"
)

// I'm told these are constant
const (
	FWeight = 146
	MWeight = 162
	NWeight = 203
)

const usage = `
Usage: glycan [-b number] weight

Options:
-b number 
	Optional defaults to 1. Specify how widely the script should bracket
	the input weights in case of measurement error. For example, given the
	default value and an input of 2000, glycan will produce reports
	for 1999 (-1), 2000 (+0), and 2001 (+1).
`

func init() {
	flag.Usage = func() {
		fmt.Println(usage)
	}
}

func main() {

	var bracketMargin int

	flag.IntVar(&bracketMargin, "b", 1, "Bracketing margin for the input weight")
	flag.Parse()

	totalWeight, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println("You must specify a numeric total weight")
		flag.Usage()
		return
	}

	for bracketOffset := bracketMargin * -1; bracketOffset <= bracketMargin; bracketOffset++ {
		testWeight := totalWeight + bracketOffset
		fmt.Printf("Input Value: %d (Adjusted by %+d):\n", testWeight, bracketOffset)
		fmt.Printf("F\tM\tN\n")
		fmt.Printf("---\t---\t---\n")

		// FWeight is the smallest of the three, so we can fit the most of it in one testWeight
		maxIterations := int(math.Ceil(float64(testWeight) / FWeight))

		// Putting the big-O in 'Ouch'
		for f := 0; f <= maxIterations; f++ {
			fComboWeight := f * FWeight

			for m := 0; m <= maxIterations; m++ {
				mComboWeight := m * MWeight

				for n := 0; n <= maxIterations; n++ {
					nComboWeight := n * NWeight

					if fComboWeight+mComboWeight+nComboWeight == testWeight {
						fmt.Printf("%d\t%d\t%d\n", f, m, n)
					}
				}
			}
		}

		fmt.Printf("\n")
	}

}
