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
	coef := make(chan float64)
	intercept := make(chan float64)
	rsq := make(chan float64)
	tval := make(chan float64)

	// Close Channels after finish
	defer close(coef)
	defer close(intercept)
	defer close(rsq)
	defer close(tval)

	// Number of Iterations
	const iter = 10000

	// Generate Random Numbers
	x, y := NumberGenerator(iter)

	// Run Regression with a Go Routine
	go Regression.LinearRegression(x, y, coef, intercept, rsq, tval)

	// Fill the channels with Result
	alpha := <-coef
	beta := <-intercept
	rsquared := <-rsq
	tvalue := <-tval

	fmt.Printf("The Intercept is :%.4f\n", alpha)
	fmt.Printf("The Regression Coefficient is :%.4f\n", beta)
	fmt.Printf("The T-Value is :%.4f\n", tvalue)
	fmt.Printf("The R-Square is :%.4f\n", rsquared)
}
