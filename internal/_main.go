package main

import (
	"fmt"
	"os"
	"pld-maker/internal/airtable"
	"pld-maker/internal/gdoc"
	"pld-maker/internal/tools"

	"github.com/go-pdf/fpdf"
)

func main() {
	credential := tools.Must(os.ReadFile("./internal/conf/credential.json"))

	airtableCli := tools.Must(airtable.NewClient(credential))
	fmt.Println("airtableCli.APIpath:", airtableCli.APIpath)
	fmt.Println("airtableCli.Token:", airtableCli.Token)

	gdocCli := tools.Must(gdoc.NewClient(credential))
	fmt.Println("gdocCli.APIpath:", gdocCli.APIpath)
	fmt.Println("gdocCli.Token:", gdocCli.Token)

	// client := &http.Client{Timeout: 10 * time.Second}
	// data, err := tools.RequestGet(client, "https://pkg.go.dev/net/http", url.Values{})
	//
	//	if err == nil {
	//		print(string(data))
	//	} else {
	//
	//		print(err.Error())
	//	}
	pdf := fpdf.New("P", "mm", "A4", "")

	type countryType struct {
		nameStr, capitalStr, areaStr, popStr string
	}
	countryList := make([]countryType, 0, 8)
	header := []string{"Country", "Capital", "Area (sq km)", "Pop. (thousands)"}
	left := (210.0 - 4*40) / 2
	pdf.SetX(left)

	basicTable := func() {
		left := (210.0 - 4*40) / 2
		pdf.SetX(left)
		for _, str := range header {
			pdf.CellFormat(40, 7, str, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
		for _, c := range countryList {
			pdf.SetX(left)
			pdf.CellFormat(40, 6, c.nameStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.capitalStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.areaStr, "1", 0, "", false, 0, "")
			pdf.CellFormat(40, 6, c.popStr, "1", 0, "", false, 0, "")
			pdf.Ln(-1)
		}
	}

	basicTable()

	/*	pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(40, 10, "Hello, world")
		pdf.AddPage()
		pdf.Cell(40, 10, "Hi, world")

		//////////
		pdf.AddPage()
		pdf.SetFont("Arial", "", 15)
		pdf.Write(8, "This line doesn't belong to any layer.\n")*/

	pdf.OutputFileAndClose("hello.pdf")
}
