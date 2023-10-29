package pldwrapper

import (
	"fmt"
	"pld-maker/v1/internal/interface/db"
	"pld-maker/v1/internal/interface/pdf"
	"strconv"
)

// Summary
func (cli *Client) TableOfContent(versions []db.Version, schemas []db.Schema, sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card) {
	(*cli.PdfClient).NewPage()
	(*cli.PdfClient).Text(pdf.Text{Data: "Summary:", Params: &pdf.TextParams{Bold: true}})
	(*cli.PdfClient).NewLine()

	//Calcul index
	var (
		//Info
		sizeFirstPage    int = 1
		sizeDescription  int = 1
		sizeVersion      int = 1 + (len(versions)+1)/30
		sizeSummary      int = 1
		sizeSchema       int = len(schemas)
		sizeDeliveryCard int = (len(cards)) / 20
		sizeUserStories  int = 1 + len(cards)
		//PagesNumber
		posSchema         int = sizeFirstPage + sizeDescription + sizeVersion + sizeSummary + 1
		posDeliveryCard   int = posSchema + sizeSchema
		posUserStories    int = posDeliveryCard + sizeDeliveryCard
		posAdvanceRepport int = posUserStories + sizeUserStories
	)
	//Size
	fmt.Println("sizeFirstPage", sizeFirstPage)
	fmt.Println("sizeDescription", sizeDescription)
	fmt.Println("sizeVersion", sizeVersion)
	fmt.Println("sizeSummary", sizeSummary)
	fmt.Println("sizeSchema", sizeSchema)
	//Pos
	fmt.Println("posSchema", posSchema)
	fmt.Println("posDeliveryCard", posDeliveryCard)
	fmt.Println("posUserStories", posUserStories)
	fmt.Println("posAdvanceRepport", posAdvanceRepport)
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
						Str:     strconv.Itoa(posSchema),
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
						Str:     strconv.Itoa(posDeliveryCard),
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
						Str:     strconv.Itoa(posUserStories),
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
						Str:     strconv.Itoa(posAdvanceRepport),
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
