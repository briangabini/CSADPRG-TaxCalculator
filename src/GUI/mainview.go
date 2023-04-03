package main

import (
	"fmt"
	"strconv"
	"taxcalc-gui/model"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tax Calculator")

	w.SetFixedSize(true)

	monthlySalary := widget.NewEntry()
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

		sss := ComputeSss(salary)
		pagIbig := ComputePagIbig(salary)
		philHealth := ComputePhilHealth(salary)
		contributions := ComputeMonthlyContributions(salary)
		tax := ComputeIncomeTax(salary - contributions)

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

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Monthly Salary", Widget: monthlySalary},
			{Text: "SSS Contribution", Widget: sssContribution},
			{Text: "Pag-Ibig Contribution", Widget: pagIbigContribution},
			{Text: "PhilHealth Contribution", Widget: philHealthContribution},
			{Text: "Total Monthly Contributions", Widget: totalMonthlyContributions},
			{Text: "Taxable Income", Widget: taxableIncome},
			{Text: "Income Tax", Widget: incomeTax},
			{Text: "Net Pay After Tax", Widget: netPayAfterTax},
			{Text: "Total Deductions", Widget: totalDeductions},
			{Text: "Net Pay After Deductions", Widget: netPayAfterDeductions},
		},
	}

	w.SetContent(container.NewVBox(
		form,         // this is a container
		calculateBtn, // this is a button
	))

	w.ShowAndRun()
}
