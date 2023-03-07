package invoice

import (
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
)

func GenerateInvoicePDF(userId string, stockName string, quantity float64, price float64, totalCost float64) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Order Buy Confirmation")

	outputQty := fmt.Sprintf("Quantity: %.0f\n Lots", quantity)
	outputPrice := fmt.Sprintf("Price: %.0f\n", price)
	outputTotalCost := fmt.Sprintf("Total Cost: %.0f\n", totalCost)

	pdf.Ln(20)
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(40, 10, fmt.Sprintf("User ID: %s", userId))
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Stock Name: %s", stockName))
	pdf.Ln(10)
	pdf.Cell(40, 10, outputQty)
	pdf.Ln(10)
	pdf.Cell(40, 10, outputPrice)
	pdf.Ln(10)
	pdf.Cell(40, 10, outputTotalCost)
	pdf.Ln(10)
	pdf.Cell(40, 10, fmt.Sprintf("Date: %s", time.Now().Format("2006-01-02")))

	err := pdf.OutputFileAndClose(fmt.Sprintf("%s_confirmationOrder.pdf", userId))
	if err != nil {
		return err
	}

	return nil
}
