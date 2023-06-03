package main

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	_ "os/exec"
	"time"
)

func generatePDF() error {
	// Create a new PDF object
	pdf := gofpdf.New("P", "mm", "A4", "")

	// Add a new page
	pdf.AddPage()

	// Set font and size
	pdf.SetFont("Arial", "B", 16)

	// Set content
	// Add a logo image
	logoFile := "Logo_Main.png" // Path to your logo image file
	pdf.Image(logoFile, 20, 10, 15, 0, false, "", 0, "")
	pdf.Ln(12)
	pdf.Cell(40, 10, "Apnader Service")
	pdf.Ln(20)
	pdf.SetFont("Arial", "", 14)
	pdf.Cell(40, 5, "Date:")
	pdf.Cell(40, 5, time.Now().UTC().Format("1/2/2006"))
	pdf.Ln(7)
	pdf.Cell(40, 5, "Name:")
	pdf.Cell(40, 5, "Shabrul Islam")
	pdf.Ln(7)
	pdf.Cell(40, 5, "Address:")
	pdf.Cell(40, 5, "Shabrul Islam")
	pdf.Ln(10)
	// Set table headers
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(40, 10, "Product name", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, "Unit price", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, "Quantity", "1", 0, "", false, 0, "")
	pdf.CellFormat(40, 10, "Sub-total", "1", 0, "", false, 0, "")
	pdf.Ln(-1)

	// Add table rows
	pdf.SetFont("Arial", "", 12)
	for i := 0; i < 10; i++ {
		pdf.CellFormat(40, 10, fmt.Sprintf("Row %d, Column 1", i+1), "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("Row %d, Column 2", i+1), "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("Row %d, Column 3", i+1), "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 10, fmt.Sprintf("Row %d, Column 3", i+1), "1", 0, "M", false, 0, "")
		pdf.Ln(-1)
	}
	pdf.Ln(7)
	pdf.Cell(40, 5, "Sub-Total:")
	pdf.Cell(40, 5, "100")

	// Save the PDF to a temporary file
	tempFile := "output.pdf"
	err := pdf.OutputFileAndClose(tempFile)
	if err != nil {
		return err
	}

	fmt.Println("PDF generated successfully.")
	return nil
}

func downloadPDF(w http.ResponseWriter, r *http.Request) {
	// Generate the PDF
	err := generatePDF()
	if err != nil {
		log.Fatal("Error generating PDF:", err)
	}

	// Read the PDF file content
	filePath := "output.pdf"
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading PDF file:", err)
	}

	// Set the appropriate headers
	w.Header().Set("Content-Disposition", "attachment; filename=output.pdf")
	w.Header().Set("Content-Type", "application/pdf")

	// Write the PDF content to the response
	_, err = w.Write(fileContent)
	if err != nil {
		log.Fatal("Error writing PDF content:", err)
	}

	// Remove the temporary file
	err = os.Remove(filePath)
	if err != nil {
		log.Fatal("Error removing temporary file:", err)
	}
}

func main() {
	http.HandleFunc("/download", downloadPDF)

	fmt.Println("Server started on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
