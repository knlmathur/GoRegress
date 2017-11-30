package Regression

import "gonum.org/v1/gonum/stat"

//LinearRegression performs the actual Regression
//Outputs the results into the respective channels
func LinearRegression(x []float64, y []float64, res chan []float64) {

	// Make array Containing the result
	result := make([]float64, 4)

	//Regress
	alpha, beta := stat.LinearRegression(x, y, nil, false)

	//RSquared
	rsquared := stat.RSquared(x, y, nil, alpha, beta)

	//Find T Value
	tvalue := beta / stat.StdErr(stat.StdDev(x, nil), float64(len(x)))

	//Assign Alpha,Beta,T-Value & Rsquared to Array
	result[0] = alpha
	result[1] = beta
	result[2] = tvalue
	result[3] = rsquared

	// Assign to Channels
	res <- result
}
