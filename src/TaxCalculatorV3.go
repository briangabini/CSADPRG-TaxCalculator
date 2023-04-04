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
	var philHealthContribution float64 = 03
	var monthlyContributions float64 = 0
	var netPayAfterTax float64 = 0
	var taxableIncome float64 = 0
	var incomeTax float64 = 0
	var netPayAfterDeductions float64 = 0

	for {
		fmt.Print("Enter your monthly salary: ")
		fmt.Scanf("%f", &monthlySalary)

		if monthlySalary >= 0 {
			break
		} else {
			fmt.Printf("Invalid input. Please enter a positive number.\n\n")
			fmt.Scanln()
		}
	}

	sssContribution = computeSss(monthlySalary)
	pagIbigContribution = computePagIbig(monthlySalary)
	philHealthContribution = computePhilHealth(monthlySalary)
	monthlyContributions = computeMonthlyContributions(monthlySalary)
	taxableIncome = monthlySalary - monthlyContributions
	incomeTax = computeIncomeTax(taxableIncome)
	netPayAfterTax = monthlySalary - incomeTax
	totalDeductions = monthlyContributions + incomeTax
	netPayAfterDeductions = monthlySalary - totalDeductions

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
	
	switch {
		case monthlySalary < 10000:
			return 225.0
		case monthlySalary < 90000:
			return monthlySalary * 0.045/2
		default:
			return 4050.0
	}

	/*
	if monthlySalary < 10000 {
		return 225.0
	} else if monthlySalary < 90000 {
		return monthlySalary * 0.045 / 2
	} else {
		return 4050.0
	}
	*/
}

// computePagIbig() computes the Pag-Ibig contribution for a given salary
func computePagIbig(monthlySalary float64) float64 {
	
	switch {
		case monthlySalary < 1500:
			return monthlySalary * .01
		case monthlySalary < 5000:
			return monthlySalary * .02
		default:
			return 100.0
	}

	/*
	if monthlySalary < 1500 {
		return monthlySalary * .01
	} else if monthlySalary < 5000 {
		return monthlySalary * .02
	} else {
		return 100.0
	} 
	*/
}

// computeIncomeTax() computes the taxable income for a given salary
func computeIncomeTax(taxableIncome float64) float64 {
	// TODO implement this function

	switch {
		case taxableIncome < 20833:
			return 0
		case taxableIncome < 33333:
			return 0 + (taxableIncome - 20833) * 0.15
		case taxableIncome < 66667:
			return 0 + (taxableIncome - 33333) * 0.20
		case taxableIncome < 166667:
			return 0 + (taxableIncome - 66667) * 0.25
		case taxableIncome < 666667:
			return 0 + (taxableIncome - 166667) * 0.30
		default:
			return 183541.80 + (taxableIncome - 666667) * 0.35
	}

	/*
	if taxableIncome < 20833 {
		return 0
	} else if taxableIncome < 33333 {
		return 0 + (taxableIncome - 20833) * 0.15
	} else if taxableIncome < 66667 {
		return 1875 + (taxableIncome - 33333) * 0.2
	} else if taxableIncome < 166667 {
		return 8541.80 + (taxableIncome - 66667) * 0.25
	} else if taxableIncome < 666667 {
		return 33541.80 + (taxableIncome - 166667) * 0.30
	} else {
		return 183541.80 + (taxableIncome - 666667) * 0.35
	}
	*/
}
