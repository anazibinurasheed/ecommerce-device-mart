package helper

import (
	"bytes"

	"github.com/jung-kurt/gofpdf"
)

func GenerateInvoicePDF(invoiceData map[string]interface{}) []byte {
	// Create a new PDF
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Set initial position and left margin
	x, y, leftMargin := 10.0, 10.0, 10.0

	// Add the title "Device mart"
	pdf.SetXY(x, y)
	pdf.Cell(0, 10, "Device mart")
	y += 10

	// Add a line separator
	pdf.SetLineWidth(0.2)
	pdf.Line(leftMargin, y, 200-leftMargin, y)
	y += 10

	// Add invoice details to the PDF
	for key, value := range invoiceData {
		// Use SetLeftMargin to control the left margin for the text
		pdf.SetLeftMargin(leftMargin)

		// Use SetXY to set the position
		pdf.SetXY(x, y)

		// Use MultiCell to wrap text and control alignment
		pdf.MultiCell(0, 10, key+": "+value.(string), "0", "L", false)
		y += 10

		// You can customize the layout and design of the invoice here
	}

	// Save the PDF to a buffer
	var pdfBuf bytes.Buffer
	pdf.Output(&pdfBuf)

	return pdfBuf.Bytes()
}
