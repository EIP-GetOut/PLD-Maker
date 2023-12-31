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
	//Airtable
	credential := tools.Must(os.ReadFile("./conf/credential.json"))
	airtableCli := tools.Must(airtable.NewClient(credential))
	//Request Current Sprint
	currentSprints, _, currentCategories, currentCards := airtableCli.GetCurrentData()
	//Request Previous Sprints
	_, previousCategories, previousCards := airtableCli.GetPreviousData()
	paramVersions := url.Values{"filterByFormula": {""}, "sort[0][field]": {"Date"}, "sort[0][direction]": {"asc"}}
	versions := epitech.AirtableToPldVersion(tools.Must(airtableCli.ListVersions(paramVersions)).Versions)

	//PLD x Epitech
	cli := tools.Must(pld.NewClient())
	epitech.HeaderFooter(cli)
	epitech.FirstPage(cli, currentSprints.Sprints[0].Fields.Number)
	cli.AddPage()
	cli.AddDescription("Project Log Document", "PLD Getout du sprint numéro "+strconv.Itoa(currentSprints.Sprints[0].Fields.Number), "Groupe Getout", "getout_2025@labeip.epitech.eu", "2025", "24 avril 2023", versions[len(versions)-1].Version)
	cli.AddVersions(versions...)

	arraySectors, SectorsCards := epitech.ArrayMapCardsToMapArrayCard([](map[string]airtable.Cards){currentCards, previousCards})
	//Sumary
	deliveryCardsInfo, userStoriesInfo := epitech.GetSectorsInfo(SectorsCards)
	cli.AddSummary(1, 1, arraySectors, deliveryCardsInfo, userStoriesInfo)

	//UserStories
	userStories := epitech.GetUserStories(arraySectors, previousCategories, currentCategories, previousCards, currentCards)
	cli.AddPage()
	cli.AddTitle1("1. Shema Fonctionel")
	cli.AddPage()
	cli.AddTitle1("2. Carte des Livrables")
	epitech.AddDeliveryCards(cli, arraySectors, userStories)
	epitech.AddUserStories(cli, arraySectors, userStories)

	//Rapport D'avancement
	cli.AddPage()
	cli.AddTitle1("4. Rapport d'Avancement")

	fmt.Println("___________________________________________")
	paramProgress := url.Values{"filterByFormula": {""}, "sort[0][field]": {"Name"}, "sort[0][direction]": {"asc"}}
	progressReport := epitech.AirtableToPldReports(tools.Must(airtableCli.ListReports(paramProgress)).Reports)
	cli.AddProgressReport(progressReport)

	err := cli.OutputFileAndClose("hello.pdf")
	fmt.Println("error: ", err)
}

//func PrintTable(sectors airtable.Sectors, categories map[string]airtable.Categories, cards map[string]airtable.Cards) {
//	for i, sector := range sectors.Sectors {
//		if len(categories[sector.Fields.Name].Categories) == 0 {
//			continue //no categories linked to sector
//		}
//		fmt.Printf("%d %s\n", i+1, sector.Fields.Name)
//		for j, category := range categories[sector.Fields.Name].Categories {
//			if len(cards[sector.Fields.Name].Cards) == 0 {
//				continue // no card linked to sector and then to categories
//			}
//			fmt.Printf("%d.%d %s\n", i+1, j+1, category.Fields.Name)
//			for k, card := range cards[sector.Fields.Name].Cards {
//				if card.Fields.Category != nil && category.Id == card.Fields.Category[0] {
//					fmt.Printf("%d.%d.%d %s\n", i+1 /*j+*/, 1, k+1, card.Fields.Title)
//				}
//			}
//		}
//	}
//}
