package epitech

import (
	"fmt"
	"pld-maker/internal/airtable"
)

func PrintCards(sectors []string, categories map[string]airtable.Categories, previousCards map[string]airtable.Cards, currentCards map[string]airtable.Cards) {
	for i, sector := range sectors {
		if len(categories[sector].Categories) == 0 {
			continue //no categories linked to sector
		}
		fmt.Printf("%d %s\n", i+1, sector)
		for j, category := range categories[sector].Categories {
			//Current Cards
			if len(currentCards[sector].Cards) == 0 {
				continue // no card linked to sector and then to categories
			}
			fmt.Printf("%d.%d %s\n", i+1, j+1, category.Fields.Name)
			for k, card := range currentCards[sector].Cards {
				if card.Fields.Category != nil && category.Id == card.Fields.Category[0] {
					fmt.Printf("%d.%d.%d %s\n", i+1 /*j+*/, 1, k+1, card.Fields.Title)
				}
			}
			//Previous Cards
			if len(previousCards[sector].Cards) == 0 {
				continue // no card linked to sector and then to categories
			}
			fmt.Printf("%d.%d %s\n", i+1, j+1, category.Fields.Name)
			for k, card := range previousCards[sector].Cards {
				if card.Fields.Category != nil && category.Id == card.Fields.Category[0] {
					fmt.Printf("%d.%d.%d %s\n", i+1 /*j+*/, 1, k+1, card.Fields.Title)
				}
			}
		}
	}
}
