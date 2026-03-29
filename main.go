package main

// Prints tonality formulas, as in "An Improviser's OS" by Wayne Krantz, in Markdown format

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println("# Tonality Formulas")
	fmt.Println()

	for i := 1; i <= 12; i++ {
		printFormulasOfLength(i)
		fmt.Println()
	}
}

var intervalNames = []string{
	"♭2", "2", "♭3", "3", "4", "♭5", "5", "♭6", "6", "♭7", "7",
}

func printFormulasOfLength(length int) {
	fmt.Printf("## %d-Note Formulas\n", length)

	if length == 1 {
		fmt.Println("- 1")
		return
	}

	numFormulas := uint(1) << len(intervalNames)
	for bitmask := uint(1); bitmask < numFormulas; bitmask++ {
		if bits.OnesCount(bitmask) == (length - 1) {
			printFormula(bitmask)
		}
	}
}

func printFormula(bitmask uint) {
	fmt.Print("- 1")
	for i := 0; i < len(intervalNames); i++ {
		if bitmask&(1<<i) != 0 {
			fmt.Printf(" %s", intervalNames[i])
		}
	}
	fmt.Println()
}
