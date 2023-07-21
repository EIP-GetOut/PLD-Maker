package epitech

import (
	"fmt"
	"pld-maker/internal/airtable"
)

func GetSectorsInfo(airtableCards [](map[string]airtable.Cards)) (map[string]int, map[string]int) {
	var result1 = make(map[string]int)
	var result2 = make(map[string]int)
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
		result1[k] = (len(v) / 30) + 1
		result2[k] = len(v)
		fmt.Println(k, len(v))
	}
	fmt.Println("-------------------------------------------")

	return result1, result2
}
