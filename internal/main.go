package main

import (
	"fmt"
	"net/url"
	"os"
	"pld-maker/internal/airtable"
	"pld-maker/internal/epitech"
	"pld-maker/internal/pld"
	"pld-maker/internal/tools"
	"strconv"
)

func main() {
	credential := tools.Must(os.ReadFile("./conf/credential.json"))
	airtableCli := tools.Must(airtable.NewClient(credential))
	//Request Current Sprint
	currentSprints, currentSectors, currentCategories, currentCards := epitech.GetCurrentData(*airtableCli)
	PrintTable(currentSectors, currentCategories, currentCards)
	//Request Previous Sprints
	previousSectors, previousCategories, previousCards := epitech.GetPreviousData(*airtableCli)
	PrintTable(previousSectors, previousCategories, previousCards)

	paramVersions := url.Values{"filterByFormula": {""}, "sort[0][field]": {"Date"}, "sort[0][direction]": {"asc"}}
	versions := epitech.AirtableToPldVersion(tools.Must(airtableCli.ListVersions(paramVersions)).Versions)

	// PldClient
	// You can use it to build a pdf.
	//PDF
	cli := tools.Must(pld.NewClient())
	epitech.HeaderFooter(cli)
	epitech.FirstPage(cli, currentSprints.Sprints[0].Fields.Number)
	cli.AddPage()
	cli.AddDescription("Project Log Document", "PLD Getout du sprint numéro "+strconv.Itoa(currentSprints.Sprints[0].Fields.Number), "Groupe Getout", "getout_2025@labeip.epitech.eu", "2025", "24 avril 2023", versions[len(versions)-1].Version)
	cli.AddVersions(versions...)

	//	epitech.GetSectors()
	//deliveryCards := epitech.GetDeliveryCardsInfo([](map[string]airtable.Cards){currentCards, previousCards})
	arraySectors, SectorsCards := epitech.ArrayMapCardsToMapArrayCard([](map[string]airtable.Cards){currentCards, previousCards})
	deliveryCardsInfo, userStoriesInfo := epitech.GetSectorsInfo(SectorsCards)

	fmt.Println("--------------------------------")
	for key := range arraySectors {
		fmt.Println(key)
	}
	cli.AddSummary(1, 4, arraySectors, deliveryCardsInfo, userStoriesInfo)
	cli.AddPage()
	cli.AddCard("1.1.1", "CreateAccount", 20, "Utilisateur de la plateforme de type a et de context or of type of", "pouvoir me connecter", "I am myself\nyou are yourself\nhe is himself\nwe are ourselves\nyou are yourselves\nthey are themselves", "*definition of done*", 4, []string{"*assignee*"}, false)
	cli.AddCard("1.1.2", "Handler", 55, "Admin", "ajouter des livres", "*description*\n*description*", "*definition of done*", 1.5, []string{"perry", "erwan"}, false)
	cli.AddCard("1.1.3", "Info", 100, "Presse", "ajouter des pubs", "*description*\n*description*", "*definition of done*", 1, []string{"inès"}, false)
	cli.AddPage()
	cli.AddCard("1.1.4", "Test OF Size Page", 49, "Business", "ajouter des pubs", "*description*\n*description*", "*definition of done*", 1, []string{"alexandre"}, false)
	cli.AddPage()
	cli.AddCard("2.2.5", "Info", 0, "Presse", "ajouter des pubs", "*description*\n*description*", "*definition of done*", 1, []string{"inès"}, false)
	cli.AddPage()
	cli.AddCard("2.2.5", "Info", 100, "Presse", "ajouter des pubs", "*description*\n*description*", "*definition of done*", 1, []string{"inès"}, true)
	cli.AddCard("2.2.5", "Info", 99, "Presse", "ajouter des pubs", "*description*\n*description*", "*definition of done*", 1, []string{"inès"}, true)

	err := cli.OutputFileAndClose("hello.pdf")
	fmt.Println("error: ", err)
}

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

//func HeaderFooter(cli *pld.Client) {
//	cli.SetHeader("", "", "EPITECH INNOVATIVE PROJECT - PROJECT LOG DOCUMENT")
//	cli.SetFooter("", "", "", true, false)
//}
