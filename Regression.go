package main

import (
	"fmt"
	"gonum.org/v1/gonum/stat"
	"math/rand"
)

//Example using in built random number library
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

// Function to Calc the T-Stat of the intercept
func TValue(beta float64, x []float64) float64 {
	tval := beta / stat.StdErr(stat.StdDev(x, nil), float64(len(x)))
	return tval
}

func LinearRegression(iter int, coef chan float64, intercept chan float64, rsq chan float64, tval chan float64) {

	// Generate Random Numbers
	x, y := NumberGenerator(iter)

	//Regress
	alpha, beta := stat.LinearRegression(x, y, nil, false)

	//RSquared
	rsquared := stat.RSquared(x, y, nil, alpha, beta)

	//Find T Value
	tvalue := TValue(beta, x)

	// Assign to Channels
	coef <- alpha
	intercept <- beta
	rsq <- rsquared
	tval <- tvalue

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

	// Run Func
	go LinearRegression(iter, coef, intercept, rsq, tval)

	// Fill values
	alpha := <-coef
	beta := <-intercept
	rsquared := <-rsq
	tvalue := <-tval

	fmt.Printf("The Intercept is :%.4f\n", alpha)
	fmt.Printf("The Regression Coefficient is :%.4f\n", beta)
	fmt.Printf("The T-Value is :%.4f\n", tvalue)
	fmt.Printf("The R-Square is :%.4f\n", rsquared)
}
