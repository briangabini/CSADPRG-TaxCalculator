package main

import (
	"fmt"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tax Calculator")

	w.SetFixedSize(true)
	size := fyne.NewSize(700, 500)
	w.Resize(size)

	monthlySalary := widget.NewEntry()
	monthlySalary.PlaceHolder = "Enter your monthly salary..."

	sssContribution := widget.NewLabel("")
	pagIbigContribution := widget.NewLabel("")
	philHealthContribution := widget.NewLabel("")
	totalMonthlyContributions := widget.NewLabel("")
	taxableIncome := widget.NewLabel("")
	incomeTax := widget.NewLabel("")
	netPayAfterTax := widget.NewLabel("")
	totalDeductions := widget.NewLabel("")
	netPayAfterDeductions := widget.NewLabel("")

	calculateBtn := widget.NewButton("Calculate", func() {
		salary, err := strconv.ParseFloat(monthlySalary.Text, 64)
		if err != nil {
			fmt.Println("Invalid input:", err)
			return
		}

		sss := computeSss(salary)
		pagIbig := computePagIbig(salary)
		philHealth := computePhilHealth(salary)
		contributions := computeMonthlyContributions(salary)
		tax := computeIncomeTax(salary - contributions)

		sssContribution.SetText(fmt.Sprintf("%.2f", sss))
		pagIbigContribution.SetText(fmt.Sprintf("%.2f", pagIbig))
		philHealthContribution.SetText(fmt.Sprintf("%.2f", philHealth))
		totalMonthlyContributions.SetText(fmt.Sprintf("%.2f", contributions))
		taxableIncome.SetText(fmt.Sprintf("%.2f", salary-contributions))
		incomeTax.SetText(fmt.Sprintf("%.2f", tax))
		netPayAfterTax.SetText(fmt.Sprintf("%.2f", salary-tax))
		totalDeductions.SetText(fmt.Sprintf("%.2f", contributions+tax))
		netPayAfterDeductions.SetText(fmt.Sprintf("%.2f", salary-contributions-tax))
	})

	clearBtn := widget.NewButton("Clear", func() {
		monthlySalary.SetText("")
		sssContribution.SetText("")
		pagIbigContribution.SetText("")
		philHealthContribution.SetText("")
		totalMonthlyContributions.SetText("")
		taxableIncome.SetText("")
		incomeTax.SetText("")
		netPayAfterTax.SetText("")
		totalDeductions.SetText("")
		netPayAfterDeductions.SetText("")
	})

	// TODO: initialize labels here
	headingLabel := widget.NewLabel("Tax Calculator Philippines 2023")
	// headingLabel.Alignment = fyne.TextAlignCenter
	headingLabel.TextStyle.Bold = true

	subheadingLabel := widget.NewLabel("Computation Result")
	subheadingLabel.TextStyle.Bold = true

	taxComputationLabel := widget.NewLabel("Tax Computation")
	monthlyContributionsLabel := widget.NewLabel("Monthly Contributions")
	// TODO increase the size of this label if possible

	// TODO: initialize the containers here
	taxComputationForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Income Tax", Widget: incomeTax},
			{Text: "Net Pay After Tax", Widget: netPayAfterTax},
		},
	}

	monthlyContributionsForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "SSS Contribution", Widget: sssContribution},
			{Text: "PhilHealth Contribution", Widget: philHealthContribution},
			{Text: "Pag-Ibig Contribution", Widget: pagIbigContribution},
			{Text: "Total Monthly Contributions", Widget: totalMonthlyContributions},
		},
	}

	finalComputationForm := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Total Deductions", Widget: totalDeductions},
			{Text: "Net Pay After Deductions", Widget: netPayAfterDeductions},
		},
	}

	// arrange the containers here
	headingContainer := container.New(layout.NewCenterLayout(), headingLabel)
	buttonContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), calculateBtn, clearBtn, layout.NewSpacer())
	subheadingContainer := container.New(layout.NewCenterLayout(), subheadingLabel)

	// TODO create two vbox containers for the tax computation and the monthly contributions; and place in a grid container
	leftcontainer := container.New(layout.NewVBoxLayout(), taxComputationLabel, taxComputationForm)
	rightcontainer := container.New(layout.NewVBoxLayout(), monthlyContributionsLabel, monthlyContributionsForm)
	computationResultGrid := container.New(layout.NewGridLayout(2), leftcontainer, rightcontainer)

	// container2 := container.New(layout.NewHBoxLabout(), layout.NewSpacer(), layout.newSpacer())
	finalComputationContainer := container.New(layout.NewCenterLayout(), finalComputationForm)

	content := container.New(layout.NewVBoxLayout(), headingContainer, monthlySalary, buttonContainer, subheadingContainer, computationResultGrid, finalComputationContainer)

	w.SetContent(content)

	w.ShowAndRun()
}

func computeMonthlyContributions(monthlySalary float64) float64 {
	return computePagIbig(monthlySalary) + computePhilHealth(monthlySalary) + computeSss(monthlySalary)
}

// computeSss() computes the SSS contribution for a given salary
func computeSss(monthlySalary float64) float64 {
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
func computePhilHealth(monthlySalary float64) float64 {

	switch {
	case monthlySalary < 10000:
		return 225.0
	case monthlySalary < 90000:
		return monthlySalary * 0.045 / 2
	default:
		return 4050.0
	}
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
		return 0 + (taxableIncome-20833)*0.15
	case taxableIncome < 66667:
		return 1875 + (taxableIncome-33333)*0.20
	case taxableIncome < 166667:
		return 8541.80 + (taxableIncome-66667)*0.25
	case taxableIncome < 666667:
		return 33541.80 + (taxableIncome-166667)*0.30
	default:
		return 183541.80 + (taxableIncome-666667)*0.35
	}
}
