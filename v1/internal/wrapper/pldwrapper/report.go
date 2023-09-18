package pldwrapper

import (
	"pld-maker/v1/internal/interface/db"
	"pld-maker/v1/internal/interface/pdf"
)

func reportTableTemplate(name string, notes string) pdf.Table {
	return pdf.Table{
		Rows: []pdf.Row{
			{
				Cells: []pdf.Cell{
					{
						Str:     name,
						Percent: 100,
						Params: &pdf.CellParams{
							Background: &pdf.Color{R: 164, G: 194, B: 244},
							TextColor:  &pdf.Color{R: 0, G: 0, B: 0},
							Size:       15,
							Bold:       true,
						},
					},
				},
			},
			{
				Cells: []pdf.Cell{
					{
						Str:     notes,
						Percent: 100,
						Params: &pdf.CellParams{
							Background: &pdf.Color{R: 255, G: 255, B: 255},
							TextColor:  &pdf.Color{R: 0, G: 0, B: 0},
						},
					},
				},
			},
		},
	}
}

// Show reports
func (cli *Client) Report(reports []db.Report) {
	var (
		global          pdf.Table
		problem         pdf.Table
		commment        pdf.Table
		individualTable pdf.Table = pdf.Table{Rows: []pdf.Row{
			{
				Cells: []pdf.Cell{
					{
						Str:     "Avancement individuel",
						Percent: 100,
						Params: &pdf.CellParams{
							Background: &pdf.Color{R: 164, G: 194, B: 244},
							TextColor:  &pdf.Color{R: 0, G: 0, B: 0},
							Size:       15,
							Bold:       true,
						},
					},
				},
			},
		}}
	)

	//Sort
	for _, report := range reports {
		switch report.Name {
		case "Global":
			global = reportTableTemplate("Avancement global pour ce rendez-vous", report.Notes)
		case "Problems":
			problem = reportTableTemplate("Points bloquants", report.Notes)
		case "Comments":
			commment = reportTableTemplate("Commentaire général", report.Notes)
		default:
			individualTable.Rows = append(individualTable.Rows, []pdf.Row{
				{
					Cells: []pdf.Cell{
						{
							Str:     report.Name,
							Percent: 100,
							Params: &pdf.CellParams{
								Background: &pdf.Color{R: 201, G: 218, B: 248},
								TextColor:  &pdf.Color{R: 0, G: 0, B: 0},
								Bold:       true,
							},
						},
					},
				},
				{
					Cells: []pdf.Cell{
						{
							Str:     report.Notes,
							Percent: 100,
							Params: &pdf.CellParams{
								Background: &pdf.Color{R: 255, G: 255, B: 255},
								TextColor:  &pdf.Color{R: 0, G: 0, B: 0},
							},
						},
					},
				},
			}...)
		}
	}

	(*cli.PdfClient).NewPage()
	(*cli.PdfClient).Text(pdf.Text{Data: "4. Rapport d'avancement:", Params: &pdf.TextParams{Bold: true}})
	(*cli.PdfClient).NewLine()
	//Global
	(*cli.PdfClient).Table(global)
	(*cli.PdfClient).NewLine()
	//Individual
	(*cli.PdfClient).Table(individualTable)
	(*cli.PdfClient).NewLine()
	// Problems
	(*cli.PdfClient).Table(problem)
	(*cli.PdfClient).NewLine()
	// Comments
	(*cli.PdfClient).Table(commment)
}
