package main

import (
	"fmt"
	"pld-maker/v0/internal/airtable"
)

func PrintTable(sectors airtable.Sectors, categories map[string]airtable.Categories, cards map[string]airtable.Cards) {
	for i, sector := range sectors.Sectors {
		if len(categories[sector.Fields.Name].Categories) == 0 {
			continue //no categories linked to sector
		}
		fmt.Printf("%d %s\n", i+1, sector.Fields.Name)
		for j, category := range categories[sector.Fields.Name].Categories {
			if len(cards[sector.Fields.Name].Cards) == 0 {
				continue // no card linked to sector and then to categories
			}
			fmt.Printf("%d.%d %s\n", i+1, j+1, category.Fields.Name)
			for k, card := range cards[sector.Fields.Name].Cards {
				if card.Fields.Category != nil && category.Id == card.Fields.Category[0] {
					fmt.Printf("%d.%d.%d %s\n", i+1 /*j+*/, 1, k+1, card.Fields.Title)
				}
			}
		}
	}
}
