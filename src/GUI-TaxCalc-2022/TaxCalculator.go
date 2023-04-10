package main

import (
	"fmt"
	"math"
	"strconv"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tax Calculator")
	a.Settings().SetTheme(theme.DarkTheme())

	w.SetFixedSize(true)
	size := fyne.NewSize(500, 400)
	w.Resize(size)

	monthlySalary := widget.NewEntry()
	monthlySalary.PlaceHolder = "Enter Salary..."

	sssContribution := canvas.NewText("--", color.White)
	pagIbigContribution := canvas.NewText("--", color.White)
	philHealthContribution := canvas.NewText("--", color.White)
	totalMonthlyContributions := canvas.NewText("--", color.White)
	incomeTax := canvas.NewText("--", color.White)
	netPayAfterTax := canvas.NewText("--", color.White)
	totalDeductions := canvas.NewText("--", color.White)
	netPayAfterDeductions := canvas.NewText("--", color.White)

	sssContribution.Alignment = fyne.TextAlignTrailing
	pagIbigContribution.Alignment = fyne.TextAlignTrailing
	philHealthContribution.Alignment = fyne.TextAlignTrailing
	totalMonthlyContributions.Alignment = fyne.TextAlignCenter
	totalDeductions.Alignment = fyne.TextAlignCenter
	netPayAfterDeductions.Alignment = fyne.TextAlignCenter

	monthlyContributionsLabel := canvas.NewText("Monthly Contributions", color.White)
	sssText := canvas.NewText("SSS", color.White)
	philHealthText := canvas.NewText("PhilHealth", color.White)
	pagIbigText := canvas.NewText("Pag-Ibig", color.White)
	totalMonthlyText := canvas.NewText("Total Contributions", color.White)
	taxComputationLabel := canvas.NewText("Tax Computation", color.White)
	incomeTaxText := canvas.NewText("Income Tax", color.White)
	netPayTaxText := canvas.NewText("Net Pay After Tax", color.White)
	totalDeductionLabel := canvas.NewText("Total Deductions", color.White)
	netPayDeductionLabel := canvas.NewText("Net Pay After Deductions", color.White)
	headingLabel := canvas.NewText("Philippine Tax Calculator", color.White)
	headingAesthetic := canvas.NewText("////", color.NRGBA{R: 0, G: 255, B: 0, A: 255})

	totalMonthlyText.Alignment = fyne.TextAlignCenter
	totalDeductionLabel.Alignment = fyne.TextAlignCenter
	netPayDeductionLabel.Alignment = fyne.TextAlignCenter

	styles := []*canvas.Text{
		sssText,
		philHealthText,
		pagIbigText,
		totalMonthlyText,
		taxComputationLabel,
		incomeTaxText,
		netPayTaxText,
		totalDeductionLabel,
		netPayDeductionLabel,
		headingLabel,
		headingAesthetic,
	}

	for _, text := range styles {
		text.TextStyle.Bold = true
	}

	fonts := []*canvas.Text{
		sssContribution,
		pagIbigContribution,
		philHealthContribution,
		totalMonthlyContributions,
		incomeTax,
		netPayAfterTax,
		totalDeductions,
		netPayAfterDeductions,
		monthlyContributionsLabel,
		sssText,
		philHealthText,
		pagIbigText,
		totalMonthlyText,
		taxComputationLabel,
		incomeTaxText,
		netPayTaxText,
	}

	for _, text := range fonts {
		text.TextSize = 16
	}

	headers := []*canvas.Text{
		totalDeductionLabel,
		totalDeductions,
		netPayDeductionLabel,
		netPayAfterDeductions,
		headingLabel,
		headingAesthetic,
	}

	for _, text := range headers {
		text.TextSize = 18
	}

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

		sssContribution.Text = fmt.Sprintf("%.2f", sss)
		pagIbigContribution.Text = fmt.Sprintf("%.2f", pagIbig)
		philHealthContribution.Text = fmt.Sprintf("%.2f", philHealth)
		totalMonthlyContributions.Text = fmt.Sprintf("%.2f", contributions)

		incomeTax.Text = fmt.Sprintf("%.2f", tax)
		netPayAfterTax.Text = fmt.Sprintf("%.2f", salary-tax)

		totalDeductions.Text = fmt.Sprintf("%.2f", contributions+tax)
		netPayAfterDeductions.Text = fmt.Sprintf("%.2f", salary-contributions-tax)

	})

	clearBtn := widget.NewButton("Clear", func() {
		monthlySalary.SetText("")
		sssContribution.Text = "--"
		pagIbigContribution.Text = "--"
		philHealthContribution.Text = "--"
		totalMonthlyContributions.Text = "--"
		incomeTax.Text = "--"
		netPayAfterTax.Text = "--"
		totalDeductions.Text = "--"
		netPayAfterDeductions.Text = "--"
	})

	sssRow := container.New(layout.NewHBoxLayout(), sssText, layout.NewSpacer(), sssContribution)
	philHealthRow := container.New(layout.NewHBoxLayout(), philHealthText, layout.NewSpacer(), philHealthContribution)
	pagIbigRow := container.New(layout.NewHBoxLayout(), pagIbigText, layout.NewSpacer(), pagIbigContribution)

	totalMonthlySpace := container.New(layout.NewCenterLayout(), totalMonthlyContributions)
	totalMonthlyCol := container.New(layout.NewVBoxLayout(), totalMonthlyText, totalMonthlySpace)

	incomeTaxSpace := container.New(layout.NewHBoxLayout(), incomeTax)
	netPayTaxSpace := container.New(layout.NewHBoxLayout(), netPayAfterTax)
	incomeTaxCol := container.New(layout.NewVBoxLayout(), incomeTaxText, incomeTaxSpace)
	netPayTaxCol := container.New(layout.NewVBoxLayout(), netPayTaxText, netPayTaxSpace)

	totalDeductionSpace := container.New(layout.NewCenterLayout(), totalDeductions)
	netPayDeductionSpace := container.New(layout.NewCenterLayout(), netPayAfterDeductions)
	totalDeductionCol := container.New(layout.NewVBoxLayout(), totalDeductionLabel, totalDeductionSpace)
	netPayDeductionCol := container.New(layout.NewVBoxLayout(), netPayDeductionLabel, netPayDeductionSpace)

	line := canvas.NewRectangle(color.NRGBA{R: 100, G: 100, B: 100, A: 255})
	monthlyTable := container.New(layout.NewVBoxLayout(), sssRow, philHealthRow, pagIbigRow, line, totalMonthlyCol)
	taxTable := container.New(layout.NewVBoxLayout(), incomeTaxCol, netPayTaxCol)
	finalComputeTable := container.New(layout.NewVBoxLayout(), totalDeductionCol, netPayDeductionCol)

	headingContainer := container.New(layout.NewHBoxLayout(), headingLabel, layout.NewSpacer(), headingAesthetic)
	buttonContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), calculateBtn, clearBtn, layout.NewSpacer())

	leftcontainer := container.New(layout.NewVBoxLayout(), monthlySalary, buttonContainer, taxComputationLabel, taxTable)
	rightcontainer := container.New(layout.NewVBoxLayout(), monthlyContributionsLabel, monthlyTable)

	computationResult := container.New(layout.NewGridLayout(2), leftcontainer, rightcontainer)
	finalComputationContainer := container.New(layout.NewCenterLayout(), finalComputeTable)

	content := container.New(layout.NewVBoxLayout(), headingContainer, computationResult, finalComputationContainer)

	w.SetContent(content)

	w.ShowAndRun()
}

// computeMonthlyContributions() computes the monthly contributions for a given salary
func computeMonthlyContributions(monthlySalary float64) float64 {
	return computePagIbig(monthlySalary) + computePhilHealth(monthlySalary) + computeSss(monthlySalary)
}

// computeSss() computes the SSS contribution for a given salary
// changed to 2022 rates
func computeSss(monthlySalary float64) float64 {
	var multiplier float64 = 0    // determines the multiplier for the SSS contribution
	var isWholeNumber bool = true // determines if the multiplier is a whole number

	if monthlySalary >= 24750 {
		return 1125.0
	} else if monthlySalary >= 3250 {
		multiplier = monthlySalary - 3250
		multiplier /= 500

		isWholeNumber = math.Mod(multiplier, 1) == 0

		if isWholeNumber {
			multiplier++
		} else {
			multiplier = math.Ceil(multiplier)
		}

		return 135 + multiplier*22.5
	} else {
		return 135.0
	}
}

// computePhilHealth() computes the PhilHealth contribution for a given salary
// changed to 2022 rates
func computePhilHealth(monthlySalary float64) float64 {

	switch {
	case monthlySalary < 10000:
		return 200.0
	case monthlySalary < 80000:
		return monthlySalary * 0.04 / 2
	default:
		return 3200.0
	}
}

// computePagIbig() computes the Pag-Ibig contribution for a given salary
// unchanged from 2021 rates
func computePagIbig(monthlySalary float64) float64 {
	switch {
	case monthlySalary < 1500:
		return monthlySalary * .01
	case monthlySalary < 5000:
		return monthlySalary * .02
	default:
		return 100.0
	}
}

// computeIncomeTax() computes the taxable income for a given salary
// changed to 2022 rates
func computeIncomeTax(taxableIncome float64) float64 {
	switch {
	case taxableIncome < 20833:
		return 0
	case taxableIncome < 33333:
		return 0 + (taxableIncome-20833)*0.15
	case taxableIncome < 66667:
		return 2500 + (taxableIncome-33333)*0.20
	case taxableIncome < 166667:
		return 10833.33 + (taxableIncome-66667)*0.25
	case taxableIncome < 666667:
		return 40833.33 + (taxableIncome-166667)*0.30
	default:
		return 200833.33 + (taxableIncome-666667)*0.35
	}
}
