package main

import (
	"github.com/go-pdf/fpdf"
)

func main() {
	pdf := fpdf.New("P", "mm", "A4", "")

	//	countryList := make([]countryType, 0, 8)
	header1 := []string{"Country", "Capital", "Area (sq km)", "Pop. (thousands) treergerev"}
	left := (210.0 - 4*40) / 2
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)
	pdf.SetFillColor(100, 250, 100)
	generateCard := func() {
		pdf.SetX(left)
		for _, str := range header1 {
			pdf.CellFormat(40, 7, str, "1", 0, "", true, 0, "")
		}
		pdf.Ln(-1)
	}
	generateCard()

	/*for _, c := range header1 {
		pdf.SetX(left)
		pdf.CellFormat(40, 6, c+"0", "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 6, c+"1", "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 6, c+"2", "1", 0, "", false, 0, "")
		pdf.CellFormat(40, 6, c+"3", "1", 0, "", false, 0, "")
		pdf.Ln(-1)
	}

		fancyTable := func() {
			// Colors, line width and bold font
			pdf.SetFillColor(255, 0, 0)
			pdf.SetTextColor(255, 255, 255)
			pdf.SetDrawColor(128, 0, 0)
			pdf.SetLineWidth(.3)
			pdf.SetFont("", "B", 0)
			// 	Header
			w := []float64{40, 35, 40, 45}
			wSum := 0.0
			for _, v := range w {
				wSum += v
			}
			left := (210 - wSum) / 2
			pdf.SetX(left)
			for j, str := range header {
				pdf.CellFormat(w[j], 7, str, "1", 0, "C", true, 0, "")
			}
			pdf.Ln(-1)
			// Color and font restoration
			pdf.SetFillColor(224, 235, 255)
			pdf.SetTextColor(0, 0, 0)
			pdf.SetFont("", "", 0)
			// 	Data
			fill := false
			for _, c := range countryList {
				pdf.SetX(left)
				pdf.CellFormat(w[0], 6, c.nameStr, "LR", 0, "", fill, 0, "")
				pdf.CellFormat(w[1], 6, c.capitalStr, "LR", 0, "", fill, 0, "")
				pdf.CellFormat(w[2], 6, strDelimit(c.areaStr, ",", 3),
					"LR", 0, "R", fill, 0, "")
				pdf.CellFormat(w[3], 6, strDelimit(c.popStr, ",", 3),
					"LR", 0, "R", fill, 0, "")
				pdf.Ln(-1)
				fill = !fill
			}
			pdf.SetX(left)
			pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
		}
		//loadData(example.TextFile("countries.txt"))
		pdf.SetFont("Arial", "", 14)
		pdf.AddPage()
		basicTable()
		pdf.AddPage()
		improvedTable()
		pdf.AddPage()
		fancyTable()*/
	//fileStr := example.Filename("Fpdf_CellFormat_tables")
	pdf.OutputFileAndClose("hello.pdf")
	//example.SummaryCompare(err, fileStr)
}
