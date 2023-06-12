package main

import (
	"fmt"
)

func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed.\n", area/10)
}

func main() {
	paintNeeded(4.2, 3.4)
	paintNeeded(5.2, 3.6)
	paintNeeded(5.0, 3.3)
	/*
		var height, width, area float64
		width = 5.2
		height = 3
		area = width * height
		fmt.Printf("%.2f liters needed.\n", area)
		width = 5.3
		height = 4
		area = width * height
		fmt.Printf("%.2f liters needed.\n", area)
	*/
	/*
		fmt.Printf("the %s cost %d cents each.\n", "gum", 23)
		fmt.Printf("That will be $%f please.\n", 0.21*6)
	*/
}
