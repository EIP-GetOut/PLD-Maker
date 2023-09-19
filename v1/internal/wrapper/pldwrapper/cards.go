package pldwrapper

import (
	"fmt"
	"pld-maker/v1/internal/interface/db"
	"pld-maker/v1/internal/interface/pdf"
	"pld-maker/v1/internal/tools"
)

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

// List Cards by Categories and by Sectors
func (cli *Client) ListCards(sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card) {
	(*cli.PdfClient).NewPage()
	(*cli.PdfClient).Text(pdf.Text{Data: "2. Cartes des livrables:", Params: &pdf.TextParams{Bold: true}})
	(*cli.PdfClient).NewLine()

	for _, sector := range sectors {
		fmt.Println(tools.Red(sector.Name + " - " + sector.Id))
		(*cli.PdfClient).Heading2(pdf.Text{Data: sector.Name})
		//category of sector
		for _, category := range getSectorCategories(sector, categories) {
			fmt.Println(tools.Yellow("\t" + category.Name))
			(*cli.PdfClient).Text(pdf.Text{Data: "    " + category.Name, Params: &pdf.TextParams{Bold: true}})
			//card of category
			for _, card := range getSectorAndCategoryCards(sector, category, cards) {
				fmt.Println(tools.Green("\t\t" + card.Title))
				(*cli.PdfClient).Text(pdf.Text{Data: "        " + card.Title})
			}
		}
	}
}

// Show Cards
func (cli *Client) Cards(sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card) {
	(*cli.PdfClient).NewPage()
	(*cli.PdfClient).Text(pdf.Text{Data: "3. User Stories:", Params: &pdf.TextParams{Bold: true}})
	(*cli.PdfClient).NewLine()
}
