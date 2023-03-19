package main

import "fmt"

func main() {
	var salary float64 = 0

	fmt.Print("Enter your salary: ")
	fmt.Scanf("%f", &salary)

	fmt.Printf("SSS: %.2f\n", computeSss(salary))
	fmt.Printf("Pag-ibig: %.2f\n", computePagIbig(salary))
	fmt.Printf("PhilHealth: %.2f\n", computePhilHealth(salary))
	fmt.Printf("Monthly Contributions: %.2f\n", computeMonthlyContributions(salary))
}

// computeMonthlyContributions() computes the monthly contributions for a given salary
func computeMonthlyContributions(monthlySalary float64) float64 {
	return computePagIbig(monthlySalary) + computePhilHealth(monthlySalary) + computeSss(monthlySalary)
}

// computeSss() computes the SSS contribution for a given salary
func computeSss(monthlySalary float64) float64 {
	if monthlySalary < 4250 {
		return 180
	} else if monthlySalary < 4750 {
		return 202.5
	} else if monthlySalary < 5250 {
		return 225
	} else if monthlySalary < 5750 {
		return 247.5
	} else if monthlySalary < 6250 {
		return 270
	} else if monthlySalary < 6750 {
		return 292.50
	} else if monthlySalary < 7250 {
		return 315
	} else if monthlySalary < 7750 {
		return 337.5
	} else if monthlySalary < 8250 {
		return 360
	} else if monthlySalary < 8750 {
		return 382.5
	} else if monthlySalary < 9250 {
		return 405
	} else if monthlySalary < 9750 {
		return 427.50
	} else if monthlySalary < 10250 {
		return 450
	} else if monthlySalary < 10750 {
		return 472.5
	} else if monthlySalary < 11250 {
		return 495
	} else if monthlySalary < 11750 {
		return 517.5
	} else if monthlySalary < 12250 {
		return 540
	} else if monthlySalary < 12750 {
		return 562.5
	} else if monthlySalary < 13250 {
		return 585
	} else if monthlySalary < 13750 {
		return 607.5
	} else if monthlySalary < 14250 {
		return 630
	} else if monthlySalary < 14750 {
		return 652.5
	} else if monthlySalary < 15250 {
		return 675
	} else if monthlySalary < 15750 {
		return 697.5
	} else if monthlySalary < 16250 {
		return 720
	} else if monthlySalary < 16750 {
		return 742.5
	} else if monthlySalary < 17250 {
		return 765
	} else if monthlySalary < 17750 {
		return 787.5
	} else if monthlySalary < 18250 {
		return 810
	} else if monthlySalary < 18750 {
		return 832.5
	} else if monthlySalary < 19250 {
		return 855
	} else if monthlySalary < 19750 {
		return 877.5
	} else if monthlySalary < 20250 {
		return 900
	} else if monthlySalary < 20750 {
		return 922.5
	} else if monthlySalary < 21250 {
		return 945
	} else if monthlySalary < 21750 {
		return 967.5
	} else if monthlySalary < 22250 {
		return 990
	} else if monthlySalary < 22750 {
		return 1012.5
	} else if monthlySalary < 23250 {
		return 1035
	} else if monthlySalary < 23750 {
		return 1057.5
	} else if monthlySalary < 24250 {
		return 1080
	} else if monthlySalary < 24750 {
		return 1102.5
	} else if monthlySalary < 25250 {
		return 1125
	} else if monthlySalary < 25750 {
		return 1147.5
	} else if monthlySalary < 26250 {
		return 1170
	} else if monthlySalary < 26750 {
		return 1192.5
	} else if monthlySalary < 27250 {
		return 1215
	} else if monthlySalary < 27750 {
		return 1237.5
	} else if monthlySalary < 28250 {
		return 1260
	} else if monthlySalary < 28750 {
		return 1282.5
	} else if monthlySalary < 29250 {
		return 1305
	} else if monthlySalary < 29750 {
		return 1327.5
	} else {
		return 1350
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
