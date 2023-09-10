package main

import (
	"pld-maker/v1/internal/fpdfwrapper"
	"pld-maker/v1/internal/interface/pdf"
	"pld-maker/v1/internal/tools"
)

func main() {
	//credential := tools.Must(os.ReadFile("./conf/credential.json"))
	//dbCli := db.Client(tools.Must(airtable.NewClient(credential)))
	pdfCli := pdf.Client(tools.Must(fpdfwrapper.NewClient()))
	// pldCli := pld.Client(tools.Must(epipld.NewClient()))
	pdfCli.NewFile("hello")
	pdfCli.Header("", "", "EPITECH INNOVATIVE PROJECT - PROJECT LOG DOCUMENT")
	pdfCli.Footer("", "", "", &pdf.FooterParams{PageNo: true, FirstPageNo: false})
	pdfCli.NewPage()
	//	pdfCli
	pdfCli.Title("Title", nil)
	pdfCli.SubTitle("SubTitle", nil)
	pdfCli.Heading1("Heading1", nil)
	pdfCli.Heading2("Heading2", nil)
	pdfCli.Text("Text: loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum. loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum. loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum. loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum. loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum.", nil)
	pdfCli.Image("../conf/epitech.png", 150, 45, &pdf.ImageParams{X: 0.5, XPercent: true, TopLeftGravity: true})
	table := pdf.Table{
		Rows: []pdf.Row{
			{
				Cells: []pdf.Cell{
					{Str: "___", Percent: 30, Params: nil},
					{Str: "---", Percent: 30, Params: nil},
					{Str: "====", Percent: 40, Params: nil},
				},
				Params: &pdf.TableParams{Background: &pdf.Color{R: 255, G: 255, B: 255}},
			},
			{
				Cells: []pdf.Cell{
					{Str: "__", Percent: 20, Params: nil},
					{Str: "--", Percent: 20, Params: nil},
					{Str: "======", Percent: 60, Params: nil},
				},
			},
		},
		Params: &pdf.TableParams{Background: &pdf.Color{R: 255, G: 255, B: 255}},
	}

	pdfCli.Table(table)
	pdfCli.CloseFile()
}

/*func main() {
	// Create a new PDF document
	pdf := fpdf.New("P", "mm", "A4", "")

	// Add a page to the PDF
	pdf.AddPage()

	// Set the X and Y position for text
	x := 20.0 // X-coordinate
	y := 40.0 // Y-coordinate

	// Set the font after setting the position
	pdf.SetXY(x, y)
	pdf.SetFont("Arial", "B", 16) // Set the font

	// Text to be displayed
	text := "Hello, this is a sample text."

	// Add text at the specified position
	pdf.Cell(0, 0, text)

	// Output the PDF to a file
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		log.Fatal(err)
	}
}*/
