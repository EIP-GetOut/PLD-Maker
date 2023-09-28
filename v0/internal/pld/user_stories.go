package pld

import (
	"fmt"
	"pld-maker/v0/internal/airtable"
)

// Type
type CardInfo struct {
	Card     airtable.Card
	Archived bool
}

type CategoriesCards map[string]([]CardInfo)

type UserStories map[string]CategoriesCards

func (cli *Client) PrintUserStories(arraySectors []string, userStories UserStories) {

	for i, sector := range arraySectors {
		fmt.Println("-", sector)

		j := 0
		for categories, cards := range userStories[sector] {
			fmt.Println("- -", categories)
			for k, card := range cards {
				fmt.Println("- - -", card.Card.Fields.Title)
				cli.AddCard(fmt.Sprintf("%d,%d,%d", i+1, j+1, k+1), card.Card.Fields.Title, card.Card.Fields.Progress, card.Card.Fields.AsWho, card.Card.Fields.IWant, card.Card.Fields.Description, card.Card.Fields.DefinitionOfDone, card.Card.Fields.Jh, AssigneeToString(card.Card.Fields.Assignee), card.Archived)
			}
			j++
		}
	}
}

func AssigneeToString(assignees []airtable.Assignee) []string {
	var result []string
	for _, assignee := range assignees {
		result = append(result, assignee.Name)
	}
	return result
}
