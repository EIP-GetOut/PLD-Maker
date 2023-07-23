package epitech

import (
	"fmt"
	"pld-maker/internal/airtable"
	"pld-maker/internal/pld"
)

// Type
type CardInfo struct {
	Card     airtable.Card
	Archived bool
}

type CategoriesCards map[string]([]CardInfo)

type UserStories map[string]CategoriesCards

// Func
func GetUserStories(arraySectors []string, previousCategories map[string]airtable.Categories, currentCategories map[string]airtable.Categories, previousCards map[string]airtable.Cards, currentCards map[string]airtable.Cards) UserStories {
	result := make(UserStories)

	for i, sector := range arraySectors {
		categoriesCards := make(CategoriesCards)

		fmt.Println(i+1, sector)
		for j, category := range previousCategories[sector].Categories {
			if len(previousCards[sector].Cards) == 0 {
				continue // no card linked to sector and then to categories
			}
			fmt.Printf("%d.%d %s\n", i+1, j+1, category.Fields.Name)
			var cardInfos []CardInfo
			for k, card := range previousCards[sector].Cards {
				if card.Fields.Category != nil && category.Id == card.Fields.Category[0] {
					fmt.Printf("%d.%d.%d %s\n", i+1, j+1, k+1, card.Fields.Title)
					cardInfos = append(cardInfos, CardInfo{Card: card, Archived: true})
				}
			}
			categoriesCards[category.Fields.Name] = append(categoriesCards[category.Fields.Name], cardInfos...)
		}
		for j, category := range currentCategories[sector].Categories {
			if len(currentCards[sector].Cards) == 0 {
				continue // no card linked to sector and then to categories
			}
			var cardInfos []CardInfo
			fmt.Printf("%d.%d %s\n", i+1, j+1, category.Fields.Name)
			for k, card := range currentCards[sector].Cards {
				if card.Fields.Category != nil && category.Id == card.Fields.Category[0] {
					fmt.Printf("%d.%d.%d %s\n", i+1, j+1, k+1, card.Fields.Title)
					cardInfos = append(cardInfos, CardInfo{Card: card, Archived: false})
				}
			}
			categoriesCards[category.Fields.Name] = append(categoriesCards[category.Fields.Name], cardInfos...)
		}
		result[sector] = categoriesCards
	}

	return result
}

func AssigneeToString(assignees []airtable.Assignee) []string {
	var result []string
	for _, assignee := range assignees {
		result = append(result, assignee.Name)
	}
	return result
}

func AddUserStories(cli *pld.Client, arraySectors []string, userStories UserStories) {
	tr := cli.UnicodeTranslatorFromDescriptor("")
	for i, sector := range arraySectors {
		fmt.Println("-", sector)
		cli.AddPage()
		if i == 0 {
			cli.AddTitle1("3. User Stories")
		}
		cli.Pdf.SetDrawColor(255, 255, 255)
		cli.Pdf.SetTextColor(0, 0, 0)
		cli.Pdf.SetFont("Arial", "B", 10)
		cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
		cli.Pdf.MultiCell(cli.CardWith, 7, tr(sector), "1", "", false)

		j := 0
		alreadyHere := false
		for categories, cards := range userStories[sector] {
			fmt.Println("- -", categories)
			for k, card := range cards {
				fmt.Println("- - -", card.Card.Fields.Title)
				if j != 0 || alreadyHere == true {
					cli.AddPage()
				}
				cli.AddCard(fmt.Sprintf("%d.%d.%d", i+1, j+1, k+1), card.Card.Fields.Title, card.Card.Fields.Progress*100, card.Card.Fields.AsWho, card.Card.Fields.IWant, card.Card.Fields.Description, card.Card.Fields.DefinitionOfDone, card.Card.Fields.Jh, AssigneeToString(card.Card.Fields.Assignee), card.Archived)
				alreadyHere = true
			}
			j++
		}
	}
}

func AddDeliveryCards(cli *pld.Client, arraySectors []string, userStories UserStories) {
	tr := cli.UnicodeTranslatorFromDescriptor("")

	for i, sector := range arraySectors {
		fmt.Println("-", sector)
		if i != 0 {
			cli.AddPage()
		}
		cli.Pdf.SetDrawColor(255, 255, 255)
		cli.Pdf.SetTextColor(0, 0, 0)
		cli.Pdf.SetFont("Arial", "B", 10)
		cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
		cli.Pdf.MultiCell(cli.CardWith, 7, tr(fmt.Sprintf("%d %s", i+1, sector)), "1", "", false)

		j := 0
		alreadyHere := false
		for categories, cards := range userStories[sector] {
			fmt.Println("- -", categories)
			for k, card := range cards {
				fmt.Println("- - -", card.Card.Fields.Title)
				if j != 0 || alreadyHere == true {
					if k%31 == 30 {
						cli.AddPage()
					}
				}
				cli.Pdf.SetDrawColor(255, 255, 255)
				cli.Pdf.SetTextColor(0, 0, 0)
				cli.Pdf.SetFont("Arial", "B", 10)
				cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
				cli.Pdf.MultiCell(cli.CardWith, 7, tr(fmt.Sprintf("%d.%d.%d %s", i+1, j+1, k+1, card.Card.Fields.Title)), "1", "", false)

				alreadyHere = true
			}
			j++
		}
	}
}
