package epitech

import (
	"fmt"
	"pld-maker/v0/internal/airtable"
	"pld-maker/v0/internal/tools"
)

func ArrayMapCardsToMapArrayCard(airtableCards [](map[string]airtable.Cards)) ([]string, map[string][]airtable.Card) {
	var sortMapCards = make(map[string][]airtable.Card)
	var sortArray []string
	for _, cardMaps := range airtableCards {
		for k, v := range cardMaps {
			sortArray = append(sortArray, k)
			sortMapCards[k] = append(sortMapCards[k], v.Cards...)
		}
	}
	return tools.FilterUniqueArray(sortArray), sortMapCards
}

func GetSectorsInfo(airtableArrayCard map[string][]airtable.Card) (map[string]int, map[string]int) {
	var result1 = make(map[string]int)
	var result2 = make(map[string]int)

	for k, v := range airtableArrayCard {
		result1[k] = (len(v) / 30) + 1
		result2[k] = len(v)
		fmt.Println(k, len(v))
	}
	return result1, result2
}
