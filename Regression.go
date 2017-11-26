package Regression

import "gonum.org/v1/gonum/stat"

// TValue Calculates the T-Stat of the intercept
func TValue(beta float64, x []float64) float64 {
	tval := beta / stat.StdErr(stat.StdDev(x, nil), float64(len(x)))
	return tval
}

//LinearRegression performs the actual Regression
//Outputs the results into the respective channels
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
