package main

import (
	"fmt"
	"math/rand"

	"github.com/knlmathur/GoRegress"
)

//NumberGenerator creates an example using the in built random number library
func NumberGenerator(iter int) (x []float64, y []float64) {

	// Generate a 6×6 matrix of random values
	x = make([]float64, iter)
	y = make([]float64, iter)

	for i := range x {
		x[i] = rand.NormFloat64()
		y[i] = rand.NormFloat64()
	}
	return x, y
}

func main() {

	//Create Channels
	result := make(chan []float64)

	// Close Channels after finish
	defer close(result)

	// Number of Iterations
	const iter = 10000

	// Generate Random Numbers
	x, y := NumberGenerator(iter)

	// Run Regression with a Go Routine
	go Regression.LinearRegression(x, y, result)

	// Fill the channels with Result
	res := <-result

	//Print Results
	fmt.Printf("The Intercept is :%.4f\n", res[0])
	fmt.Printf("The Regression Coefficient is :%.4f\n", res[1])
	fmt.Printf("The T-Value is :%.4f\n", res[2])
	fmt.Printf("The R-Square is :%.4f\n", res[3])
}
