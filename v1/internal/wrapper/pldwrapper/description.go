package pldwrapper

import "pld-maker/v1/internal/interface/pdf"

//darkBlue:60, 120, 216
//blue: 164, 194, 244
//lightBlue: 201, 218, 24

// Description
func (cli *Client) Description(title, object, author, e_mail, promo, last_update, version string) {
	(*cli.PdfClient).NewPage()
	(*cli.PdfClient).Text(pdf.Text{Data: "Description:", Params: &pdf.TextParams{Bold: true}})
	(*cli.PdfClient).NewLine()
	(*cli.PdfClient).Table(pdf.Table{
		Rows: []pdf.Row{
			{
				Cells: []pdf.Cell{
					{
						Str:     "Titre",
						Percent: 25,
					},
					{
						Str:     title,
						Percent: 75,
						Params: &pdf.CellParams{
							Background: &pdf.Color{R: 164, G: 194, B: 244},
							TextColor:  &pdf.Color{R: 0, G: 0, B: 0},
							Bold:       true,
						},
					},
				},
			},
			{
				Cells: []pdf.Cell{
					{Str: "Objet", Percent: 25},
					{
						Str:     object,
						Percent: 75,
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
					{Str: "Auteur", Percent: 25},
					{
						Str:     author,
						Percent: 75,
						Params: &pdf.CellParams{
							Background: &pdf.Color{R: 164, G: 194, B: 244},
							TextColor:  &pdf.Color{R: 0, G: 0, B: 0},
							Bold:       true,
						},
					},
				},
			},
			{
				Cells: []pdf.Cell{
					{Str: "Promo", Percent: 25},
					{
						Str:     promo,
						Percent: 75,
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
					{Str: "Last Update", Percent: 25},
					{
						Str:     last_update,
						Percent: 75,
						Params: &pdf.CellParams{
							Background: &pdf.Color{R: 164, G: 194, B: 244},
							TextColor:  &pdf.Color{R: 0, G: 0, B: 0},
							Bold:       true,
						},
					},
				},
			},
			{
				Cells: []pdf.Cell{
					{Str: "Version", Percent: 25},
					{
						Str:     version,
						Percent: 75,
						Params: &pdf.CellParams{
							Background: &pdf.Color{R: 201, G: 218, B: 248},
							TextColor:  &pdf.Color{R: 0, G: 0, B: 0},
							Bold:       true,
						},
					},
				},
			},
		},

		Params: &pdf.TableParams{
			RowParams: &pdf.RowParams{
				CellParams: &pdf.CellParams{
					Background: &pdf.Color{R: 60, G: 120, B: 216},
					TextColor:  &pdf.Color{R: 255, G: 255, B: 255},
					Bold:       true,
				},
				RowHeight: 8,
			},
			DrawColor: &pdf.Color{R: 255, G: 255, B: 255},
		},
	})
}
