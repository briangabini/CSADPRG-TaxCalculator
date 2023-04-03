package model

import (
	"math"
)

// computeMonthlyContributions() computes the monthly contributions for a given salary
func ComputeMonthlyContributions(monthlySalary float64) float64 {
	return computePagIbig(monthlySalary) + computePhilHealth(monthlySalary) + computeSss(monthlySalary)
}

// computeSss() computes the SSS contribution for a given salary
func ComputeSss(monthlySalary float64) float64 {
	var multiplier float64 = 0    // determines the multiplier for the SSS contribution
	var isWholeNumber bool = true // determines if the multiplier is a whole number

	if monthlySalary > 29750 {
		return 1350.0
	} else if monthlySalary >= 4250 {
		multiplier = monthlySalary - 4250
		multiplier /= 500

		isWholeNumber = math.Mod(multiplier, 1) == 0

		if isWholeNumber {
			multiplier++
		} else {
			multiplier = math.Ceil(multiplier)
		}

		return 180 + multiplier*22.5
	} else {
		return 180
	}
}

// computePhilHealth() computes the PhilHealth contribution for a given salary
func ComputePhilHealth(monthlySalary float64) float64 {
	if monthlySalary < 10000 {
		return 225.0
	} else if monthlySalary < 90000 {
		return monthlySalary * 0.045 / 2
	} else {
		return 4050.0
	}
}

// computePagIbig() computes the Pag-Ibig contribution for a given salary
func ComputePagIbig(monthlySalary float64) float64 {
	if monthlySalary < 1500 {
		return monthlySalary * .01
	} else if monthlySalary < 5000 {
		return monthlySalary * .02
	} else {
		return 100.0
	}
}

// computeIncomeTax() computes the taxable income for a given salary
func ComputeIncomeTax(taxableIncome float64) float64 {
	// TODO implement this function
	if taxableIncome < 20833 {
		return 0
	} else if taxableIncome < 33333 {
		return 0 + (taxableIncome-20833)*0.15
	} else if taxableIncome < 66667 {
		return 1875 + (taxableIncome-33333)*0.2
	} else if taxableIncome < 166667 {
		return 8541.80 + (taxableIncome-66667)*0.25
	} else if taxableIncome < 666667 {
		return 33541.80 + (taxableIncome-166667)*0.30
	} else {
		return 183541.80 + (taxableIncome-666667)*0.35
	}
}
