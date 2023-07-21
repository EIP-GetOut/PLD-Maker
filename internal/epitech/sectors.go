package epitech

import (
	"fmt"
	"pld-maker/internal/airtable"
)

func GetSectors(airtableCards [](map[string]airtable.Cards)) map[string]int {
	var result = make(map[string]int)
	var sortMapCards = make(map[string][]airtable.Card)

	fmt.Println("-------------------------------------------")
	for _, cardMaps := range airtableCards {
		for k, v := range cardMaps {
			sortMapCards[k] = append(sortMapCards[k], v.Cards...)
		}
	}
	for k, v := range sortMapCards {
		fmt.Println(k, ": {")
		for _, card := range v {
			fmt.Println("  ", card.Fields.Title)
		}
		fmt.Println("}")
	}
	fmt.Println("-------------------------------------------")
	for k, v := range sortMapCards {
		result[k] = len(v)
		fmt.Println(k, len(v))
	}
	fmt.Println("-------------------------------------------")

	return result
}
