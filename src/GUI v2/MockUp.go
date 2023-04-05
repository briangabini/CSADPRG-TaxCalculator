package main

import (
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

	sssContribution := canvas.NewText("0.00", color.White)
	pagIbigContribution := canvas.NewText("0.00", color.White)
	philHealthContribution := canvas.NewText("0.00", color.White)
	totalMonthlyContributions := canvas.NewText("0.00", color.White)
	incomeTax := canvas.NewText("0.00", color.White)
	netPayAfterTax := canvas.NewText("0.00", color.White)
	totalDeductions := canvas.NewText("0.00", color.White)
	netPayAfterDeductions := canvas.NewText("0.00", color.White)

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
	headingAesthetic := canvas.NewText("////", color.NRGBA{R: 100, G: 100, B: 100, A: 255})
	
	totalMonthlyText.Alignment = fyne.TextAlignCenter
	totalDeductionLabel.Alignment = fyne.TextAlignCenter
	netPayDeductionLabel.Alignment = fyne.TextAlignCenter

	styles := []*canvas.Text {
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

	fonts := []*canvas.Text {
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

	headers := []*canvas.Text {
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
		//TODO implement input function
	})
	
	clearBtn := widget.NewButton("Clear", func() {
		//TODO implement clear function
	})

	sssRow := container.New(layout.NewHBoxLayout(), sssText, layout.NewSpacer(), sssContribution)
	philHealthRow := container.New(layout.NewHBoxLayout(), philHealthText, layout.NewSpacer(), philHealthContribution)
	pagIbigRow := container.New(layout.NewHBoxLayout(), pagIbigText, layout.NewSpacer(), pagIbigContribution)
	totalMonthlyCol := container.New(layout.NewVBoxLayout(), totalMonthlyText, totalMonthlyContributions)

	incomeTaxCol := container.New(layout.NewVBoxLayout(), incomeTaxText, incomeTax)
	netPayTaxCol := container.New(layout.NewVBoxLayout(), netPayTaxText, netPayAfterTax)

	totalDeductionCol := container.New(layout.NewVBoxLayout(), totalDeductionLabel, totalDeductions)
	netPayDeductionCol := container.New(layout.NewVBoxLayout(), netPayDeductionLabel, netPayAfterDeductions)

	line := canvas.NewRectangle(color.NRGBA{R:100, G:100, B:100, A:255})
	monthlyTable := container.New(layout.NewVBoxLayout(), sssRow, philHealthRow, pagIbigRow, line, totalMonthlyCol)
	taxTable := container.New(layout.NewVBoxLayout(), incomeTaxCol, netPayTaxCol)
	finalComputeTable := container.New(layout.NewVBoxLayout(), totalDeductionCol, netPayDeductionCol)

	headingContainer := container.New(layout.NewHBoxLayout(),headingLabel, layout.NewSpacer(), headingAesthetic)
	buttonContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), calculateBtn, clearBtn, layout.NewSpacer())
	
	leftcontainer := container.New(layout.NewVBoxLayout(), monthlySalary, buttonContainer, taxComputationLabel, taxTable)
	rightcontainer := container.New(layout.NewVBoxLayout(), monthlyContributionsLabel, monthlyTable)

	computationResult := container.New(layout.NewGridLayout(2), leftcontainer, rightcontainer)
	finalComputationContainer := container.New(layout.NewCenterLayout(), finalComputeTable)

	content := container.New(layout.NewVBoxLayout(), headingContainer, computationResult, finalComputationContainer)

	w.SetContent(content)

	w.ShowAndRun()
}

//TODO return functionality