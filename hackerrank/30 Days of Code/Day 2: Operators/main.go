package main

import "fmt"

func solve(meal_cost float64, tip_percent int32, tax_percent int32) {
	tip := meal_cost * float64(tip_percent) / 100
	tax := meal_cost * float64(tax_percent) / 100
	total := meal_cost + tip + tax
	fmt.Printf("%.0f\n", total)
}

func main() {
	solve(100.80, 9, 12)
}
