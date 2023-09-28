package pldwrapper

import (
	"fmt"
	"pld-maker/v1/internal/interface/db"
	"pld-maker/v1/internal/interface/pdf"
	"pld-maker/v1/internal/tools"
	"strings"
)

//func

// GetCategories
func getSectorCategories(sector db.Sector, categories []db.Category) []db.Category {
	var result []db.Category

	for _, category := range categories {
		for _, categorySector := range category.Sectors {
			if sector.Id == categorySector {
				result = append(result, category)
				break
			}
		}
	}
	return result
}

// get Cards by Sector & Category
func getSectorAndCategoryCard(sector db.Sector, category db.Category, card db.Card) *db.Card {
	for _, cardCategory := range card.Category {
		//fmt.Println(tools.Red(fmt.Sprint(category.Id, cardCategory)))
		if category.Id == cardCategory {
			//fmt.Println(tools.Green(fmt.Sprint(category.Id, cardCategory)))
			for _, cardSector := range card.Sector {
				//fmt.Println(tools.Cyan(fmt.Sprint(sector.Id, cardSector)))
				if sector.Id == cardSector {
					return &card
				}
			}
		}
	}
	return nil
}

func getSectorAndCategoryCards(sector db.Sector, category db.Category, cards []db.Card) []db.Card {
	var result []db.Card

	for _, card := range cards {
		if card := getSectorAndCategoryCard(sector, category, card); card != nil {
			result = append(result, *card)
		}
	}
	return result
}

/*
 * List Cards by Categories by Sectors
 *
 */
func (cli *Client) ListCards(currentSprint db.Sprint, sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card) {
	var cardIdx int = 0

	(*cli.PdfClient).NewPage()
	(*cli.PdfClient).Text(pdf.Text{Data: "2. Cartes des livrables:", Params: &pdf.TextParams{Bold: true}})
	(*cli.PdfClient).NewLine()

	for _, sector := range sectors {
		fmt.Println(tools.Magenta(sector.Name + " - " + sector.Id))
		(*cli.PdfClient).Heading2(pdf.Text{Data: sector.Name})
		//category of sector
		for _, category := range getSectorCategories(sector, categories) {
			fmt.Println(tools.Cyan("\t" + category.Name))
			(*cli.PdfClient).Text(pdf.Text{Data: "    " + category.Name, Params: &pdf.TextParams{Bold: true}})
			//card of category
			for _, card := range getSectorAndCategoryCards(sector, category, cards) {
				var color pdf.Color
				cardIdx++

				fmt.Println(tools.BgCyan(fmt.Sprint(card.Sprint, currentSprint.Id, card.Progress)))

				if card.Progress == 0 {
					fmt.Println(tools.Red(fmt.Sprint("\t\t"+card.Title+" - ", card.Progress*100)))
					color = cli.PercentColors[0]
				} else if card.Progress < 1 {
					if card.Sprint[0] != currentSprint.Id {
						fmt.Println(tools.Red(fmt.Sprint("\t\t"+card.Title+" - ", card.Progress*100)))
						color = cli.PercentColors[0]
					} else {
						fmt.Println(tools.Yellow(fmt.Sprint("\t\t"+card.Title+" - ", card.Progress*100)))
						color = cli.PercentColors[1]
					}
				} else {
					if card.Sprint[0] != currentSprint.Id {
						fmt.Println(tools.Grey(fmt.Sprint("\t\t"+card.Title+" - ", card.Progress*100)))
						color = cli.PercentColors[3]
					} else {
						fmt.Println(tools.Green(fmt.Sprint("\t\t"+card.Title+" - ", card.Progress*100)))
						color = cli.PercentColors[2]
					}
				}
				(*cli.PdfClient).Text(pdf.Text{Data: "        " + card.Title, Params: &pdf.TextParams{TextColor: &color}})
				if cardIdx%20 == 0 {
					(*cli.PdfClient).NewPage()
				}
			}
		}
	}
}

// Show Cards

// format assignee: concate name or email prefix of full email
func formatAssignee(assignees []db.Assignee) string {
	var result string
	for _, assignee := range assignees {
		if assignee.Name != "" {
			result += assignee.Name + ",\n"
		} else if idx := strings.Index(assignee.Email, "@"); idx != -1 {
			result += assignee.Email[0:idx] + ",\n"
		} else {
			result += assignee.Email + ",\n"
		}
	}
	if len(result) >= 2 {
		result = result[:len(result)-2]
	}
	return result
}

func (cli *Client) Cards(currentSprint db.Sprint, sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card) {
	(*cli.PdfClient).NewPage()
	(*cli.PdfClient).Text(pdf.Text{Data: "3. User Stories:", Params: &pdf.TextParams{Bold: true}})
	(*cli.PdfClient).NewLine()

	for _, sector := range sectors {
		//fmt.Println(tools.Magenta(sector.Name + " - " + sector.Id))
		(*cli.PdfClient).Heading2(pdf.Text{Data: sector.Name})
		//category of sector
		for _, category := range getSectorCategories(sector, categories) {
			//fmt.Println(tools.Cyan("\t" + category.Name))
			(*cli.PdfClient).Text(pdf.Text{Data: "    " + category.Name, Params: &pdf.TextParams{Bold: true}})
			//card of category
			for _, card := range getSectorAndCategoryCards(sector, category, cards) {
				var color pdf.Color

				//fmt.Println(tools.BgCyan(fmt.Sprint(card.Sprint, currentSprint.Id, card.Progress)))

				if card.Progress == 0 {
					//fmt.Println(tools.Red(fmt.Sprint("\t\t"+card.Title+" - ", card.Progress*100)))
					color = cli.PercentColors[0]
				} else if card.Progress < 1 {
					if card.Sprint[0] != currentSprint.Id {
						//fmt.Println(tools.Red(fmt.Sprint("\t\t"+card.Title+" - ", card.Progress*100)))
						color = cli.PercentColors[0]
					} else {
						//fmt.Println(tools.Yellow(fmt.Sprint("\t\t"+card.Title+" - ", card.Progress*100)))
						color = cli.PercentColors[1]
					}
				} else {
					if card.Sprint[0] != currentSprint.Id {
						//fmt.Println(tools.Grey(fmt.Sprint("\t\t"+card.Title+" - ", card.Progress*100)))
						color = cli.PercentColors[3]
					} else {
						//fmt.Println(tools.Green(fmt.Sprint("\t\t"+card.Title+" - ", card.Progress*100)))
						color = cli.PercentColors[2]
					}
				}
				(*cli.PdfClient).Table(pdf.Table{
					Rows: []pdf.Row{
						{
							Cells: []pdf.Cell{
								{
									Str:     card.Title,
									Percent: 60,
									Params: &pdf.CellParams{
										Background: &color,
									},
								},
								{
									Str:     "Progress",
									Percent: 20,
									Params: &pdf.CellParams{
										Background: &color,
									},
								},
								{
									Str:     fmt.Sprint(card.Progress*100) + "%",
									Percent: 20,
									Params: &pdf.CellParams{
										Background: &color,
									},
								},
							},
						},
						{
							Cells: []pdf.Cell{
								{
									Str:     "En tant que:",
									Percent: 50,
									Params: &pdf.CellParams{
										Background: &color,
									},
								},
								{
									Str:     "Je veux:",
									Percent: 50,
									Params: &pdf.CellParams{
										Background: &color,
									},
								},
							},
						},
						{
							Cells: []pdf.Cell{
								{
									Str:     card.AsWho,
									Percent: 50,
								},
								{
									Str:     card.IWant,
									Percent: 50,
								},
							},
						},
						{
							Cells: []pdf.Cell{
								{
									Str:     "Description:",
									Percent: 100,
									Params: &pdf.CellParams{
										Background: &color,
									},
								},
							},
						},
						{
							Cells: []pdf.Cell{
								{
									Str:     card.Description,
									Percent: 100,
								},
							},
						},
						{
							Cells: []pdf.Cell{
								{
									Str:     "Definition Of Done:",
									Percent: 100,
									Params: &pdf.CellParams{
										Background: &color,
									},
								},
							},
						},
						{
							Cells: []pdf.Cell{
								{
									Str:     card.DefinitionOfDone,
									Percent: 100,
								},
							},
						},
						{
							Cells: []pdf.Cell{
								{
									Str:     "Charge Estimée (J/H):",
									Percent: 30,
									Params: &pdf.CellParams{
										Background: &color,
									},
								},
								{
									Str:     fmt.Sprint(card.Jh),
									Percent: 20,
								},
								{
									Str:     "Assignés (J/H):",
									Percent: 20,
									Params: &pdf.CellParams{
										Background: &color,
									},
								},
								{
									Str:     formatAssignee(card.Assignees),
									Percent: 30,
								},
							},
						},
					},
					Params: &pdf.TableParams{
						RowParams: &pdf.RowParams{
							CellParams: &pdf.CellParams{
								Background: &pdf.Color{R: 255, G: 255, B: 255},
							},
						},
					},
				})
				(*cli.PdfClient).NewPage()
			}
		}
	}
}
