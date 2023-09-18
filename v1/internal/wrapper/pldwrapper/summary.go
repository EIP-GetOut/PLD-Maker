package pldwrapper

import (
	"fmt"
	"pld-maker/v1/internal/interface/db"
	"pld-maker/v1/internal/interface/pdf"
	"strconv"
)

// Summary
func (cli *Client) Summary(versions []db.Version, schemas []db.Schema, sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card) {
	(*cli.PdfClient).NewPage()
	(*cli.PdfClient).Text(pdf.Text{Data: "Summary:", Params: &pdf.TextParams{Bold: true}})
	(*cli.PdfClient).NewLine()

	//Calcul index
	var (
		beginLen   int = 2 //firstPage+Description
		versionLen int = 1 + (len(versions)+1)/30
		summaryLen int = 1
		schemaLen  int = len(schemas)
		//		deliveryCard int = 1
		//		userStoryNb int = schemaNo + len(schemas)
		//		userStoriesNb []int = []int{}
		//		cardNb        int   = userStoryNb + cards
		//		cardsNb       []int = []int{}
	)

	fmt.Println("begin", beginLen)
	fmt.Println("version", beginLen+versionLen)
	fmt.Println("schema", beginLen+versionLen+schemaLen)
	//Display table
	(*cli.PdfClient).Table(pdf.Table{
		Rows: []pdf.Row{
			{
				Cells: []pdf.Cell{
					{
						Str:     "1. Schéma Fonctionnel",
						Percent: 90,
					},
					{
						Str:     strconv.Itoa(1 + beginLen + versionLen + summaryLen),
						Percent: 10,
					},
				},
			},
			{
				Cells: []pdf.Cell{
					{
						Str:     "2. Cartes des livrables",
						Percent: 90,
					},
					{
						Str:     strconv.Itoa(1 + beginLen + versionLen + summaryLen + schemaLen),
						Percent: 10,
					},
				},
			},
			{
				Cells: []pdf.Cell{
					{
						Str:     "3. User Stories",
						Percent: 90,
					},
					{
						Str:     strconv.Itoa((len(versions) + 1) / 30),
						Percent: 10,
					},
				},
			},
			{
				Cells: []pdf.Cell{
					{
						Str:     "4. Rapport d'avancement",
						Percent: 90,
					},
					{
						Str:     strconv.Itoa((len(versions) + 1) / 30),
						Percent: 10,
					},
				},
			},
		},
		Params: &pdf.TableParams{
			DrawColor: &pdf.Color{R: 255, G: 255, B: 255},
			RowParams: &pdf.RowParams{
				RowHeight: 9,
				CellParams: &pdf.CellParams{
					Bold: true,
				},
			},
		},
	})
}
