package main

import (
	"fmt"
	"log"
	"os"

	"github.com/signintech/gopdf"
)

func PrependText(inputFilePath string, outputFilePath, text string) {

	// Create a new PDF object
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) // Assuming A4, modify as necessary
	pdf.AddPage()

	// Set the font
	err := pdf.AddTTFFont("arial", "./arial.ttf") // Path to a TTF font file
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("arial", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.SetFillColor(255, 255, 255)

	// Add text
	pdf.SetX(10) // Adjust these coordinates and page number as necessary
	pdf.SetY(10)
	pdf.Text(text)

	// Import pages from existing PDF
	tpl := pdf.ImportPage(inputFilePath, 1, "/MediaBox") // import first page

	pdf.UseImportedTemplate(tpl, 0, 0, 595, 842) // draw imported page on the current page

	// Save the new PDF
	err = pdf.WritePdf(outputFilePath)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	// Check if the input and output file paths are provided
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input-file-path> <output-file-path>")
		return
	}
	//read input file from command line
	inputFilePath := os.Args[1]
	//read output file from command line
	outputFilePath := os.Args[2]
	fmt.Println("PDF generated successfully:", outputFilePath)
	PrependText(inputFilePath, outputFilePath, "Hello, World!")
}
