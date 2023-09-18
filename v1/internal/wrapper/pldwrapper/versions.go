package pldwrapper

import (
	"pld-maker/v1/internal/interface/db"
	"pld-maker/v1/internal/interface/pdf"
)

func head() {

}

// Version
func (cli *Client) Versions(versions []db.Version) {
	var table pdf.Table = pdf.Table{
		Rows: []pdf.Row{},
		Params: &pdf.TableParams{
			RowParams: &pdf.RowParams{
				CellParams: &pdf.CellParams{
					Size:       7,
					Background: &pdf.Color{R: 164, G: 194, B: 244},
					TextColor:  &pdf.Color{R: 0, G: 0, B: 0},
				},
			},
			DrawColor: &pdf.Color{R: 255, G: 255, B: 255},
		},
	}

	(*cli.PdfClient).NewPage()
	(*cli.PdfClient).Text(pdf.Text{Data: "Tableau des révisions:", Params: &pdf.TextParams{Bold: true}})
	(*cli.PdfClient).NewLine()
	(*cli.PdfClient).Table(pdf.Table{
		Rows: []pdf.Row{
			{
				Cells: []pdf.Cell{
					{
						Str:     "Date",
						Percent: 20,
					},
					{
						Str:     "Version",
						Percent: 10,
					}, {
						Str:     "Author",
						Percent: 20,
					}, {
						Str:     "Sections",
						Percent: 20,
					}, {
						Str:     "Comments",
						Percent: 30,
					},
				},
			},
		},
		Params: &pdf.TableParams{
			RowParams: &pdf.RowParams{
				CellParams: &pdf.CellParams{
					//					Bold:       true,
					Background: &pdf.Color{R: 60, G: 120, B: 216},
					TextColor:  &pdf.Color{R: 255, G: 255, B: 255},
					Size:       7,
				},
			},
			DrawColor: &pdf.Color{R: 255, G: 255, B: 255},
		},
	})
	//Split & Display
	for i, version := range versions {
		//Create
		table.Rows = append(table.Rows, pdf.Row{
			Cells: []pdf.Cell{
				{
					Str:     version.Date,
					Percent: 20,
				},
				{
					Str:     version.Version,
					Percent: 10,
				}, {
					Str:     version.Author,
					Percent: 20,
				}, {
					Str:     version.Sections,
					Percent: 20,
				}, {
					Str:     version.Comments,
					Percent: 30,
				},
			},
		})
		(*cli.PdfClient).Table(table)
		table.Rows = []pdf.Row{}
		//+2 used to display 39version row on first page cause header row take one
		if (i+2)%40 == 0 {
			(*cli.PdfClient).NewPage()
		}
	}
}
