package main

import (
	"fmt"
	"math"
)

func main() {
	var monthlySalary float64 = 0
	var totalDeductions float64 = 0
	var sssContribution float64 = 0
	var pagIbigContribution float64 = 0
	var philHealthContribution float64 = 0
	var monthlyContributions float64 = 0
	var netPayAfterTax float64 = 0
	var incomeTax float64 = 0
	var netPayAfterDeductions float64 = 0

	fmt.Print("Enter your monthly salary: ")
	fmt.Scanf("%f", &monthlySalary)

	sssContribution = computeSss(monthlySalary)
	pagIbigContribution = computePagIbig(monthlySalary)
	philHealthContribution = computePhilHealth(monthlySalary)
	monthlyContributions = computeMonthlyContributions(monthlySalary)
	incomeTax = computeIncomeTax(monthlySalary)
	netPayAfterTax = monthlySalary - incomeTax
	totalDeductions = monthlyContributions + incomeTax
	netPayAfterDeductions = netPayAfterTax - totalDeductions

	// Display Computations
	fmt.Println("")

	fmt.Println("Tax Computation:")
	fmt.Printf("Income Tax: %.2f\n", incomeTax)
	fmt.Printf("Net Pay After Tax: %.2f\n", netPayAfterTax)

	fmt.Println("")

	fmt.Println("Monthly Contributions:")
	fmt.Printf("SSS: %.2f\n", sssContribution)
	fmt.Printf("PhilHealth: %.2f\n", philHealthContribution)
	fmt.Printf("Pag-ibig: %.2f\n", pagIbigContribution)
	fmt.Printf("Total Contributions: %.2f\n", monthlyContributions)

	fmt.Println("")

	fmt.Printf("Total Deductions: %.2f\n", totalDeductions)
	fmt.Printf("Net Pay after Deductions: %.2f\n", netPayAfterDeductions)
}

// computeMonthlyContributions() computes the monthly contributions for a given salary
func computeMonthlyContributions(monthlySalary float64) float64 {
	return computePagIbig(monthlySalary) + computePhilHealth(monthlySalary) + computeSss(monthlySalary)
}


// computeSss() computes the SSS contribution for a given salary
func computeSss(monthlySalary float64) float64 {
	var multiplier float64 = 0 // determines the multiplier for the SSS contribution
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
		
		return 180 + multiplier * 22.5
	} else {
		return 180 
	}
}

// computePhilHealth() computes the PhilHealth contribution for a given salary
func computePhilHealth(monthlySalary float64) float64 {
	if monthlySalary < 10000 {
		return 225.0
	} else if monthlySalary < 90000 {
		return monthlySalary * 0.045 / 2
	} else {
		return 4050.0
	}
}

// computePagIbig() computes the Pag-Ibig contribution for a given salary
func computePagIbig(monthlySalary float64) float64 {
	if monthlySalary < 1500 {
		return monthlySalary * .01
	} else if monthlySalary < 5000 {
		return monthlySalary * .02
	} else {
		return 100.0
	}
}

// computeIncomeTax() computes the taxable income for a given salary
func computeIncomeTax(monthlySalary float64) float64 {
	// TODO implement this function
	if monthlySalary < 20833 {
		return 0
	} else if monthlySalary < 33333 {
		return 0 + (monthlySalary - 20833) * 0.2
	} else if monthlySalary < 66667 {
		return 0 + (monthlySalary - 33333) * 0.25
	} else if monthlySalary < 166667 {
		return 0 + (monthlySalary - 66667) * 0.3
	} else if monthlySalary < 666667 {
		return 0 + (monthlySalary - 166667) * 0.32
	} else {
		return 0 + (monthlySalary - 666667) * 0.35
	}
}

