package main

import (
	"fmt"
	"net/url"
	"os"
	"pld-maker/internal/airtable"
	"pld-maker/internal/pld"
	"pld-maker/internal/tools"
	"strconv"
)

func main() {
	credential := tools.Must(os.ReadFile("./conf/credential.json"))
	airtableCli := tools.Must(airtable.NewClient(credential))
	//fmt.Println(airtableCli.Token)

	//Request Sprint
	paramSprints := url.Values{"filterByFormula": {"FIND(\"In progress\", {Status})"}}
	sprints := tools.Must(airtableCli.ListSprints(paramSprints))
	//fmt.Println(sprints)
	//airtableCli.PrintSprints(sprints.Sprints, "")
	if len(sprints.Sprints) < 1 {
		panic("no sprint in progress")
	} else if len(sprints.Sprints) != 1 {
		panic(fmt.Sprintf("Error: %s", "multiple sprint In progress"))
	}
	//Request Sector
	paramSectors := url.Values{"filterByFormula": {""}, "sort[0][field]": {"Name"}, "sort[0][direction]": {"asc"}}
	sectors := tools.Must(airtableCli.ListSectors(paramSectors))
	//airtableCli.PrintSectors(sectors.Sectors, "")
	var categories = make(map[string]airtable.Categories)
	var cards = make(map[string]airtable.Cards)
	for _, sector := range sectors.Sectors {
		//Request Categories
		paramCategories := url.Values{"filterByFormula": {fmt.Sprintf("AND(FIND(\"%s\",CONCATENATE(\"\",{Sprint})),FIND(\"%s\",CONCATENATE(\"\",{Sector})))", sprints.Sprints[0].Fields.Title, sector.Fields.Name)}}
		categories[sector.Fields.Name] = tools.Must(airtableCli.ListCategories(paramCategories))
		//		airtableCli.PrintCategories(categories[sector.Fields.Name].Categories, "%")

		//Request Cards
		paramCards := url.Values{"filterByFormula": {fmt.Sprintf("AND(FIND(\"%s\",CONCATENATE(\"\",{Sprint})),FIND(\"%s\",CONCATENATE(\"\",{Sector})))", sprints.Sprints[0].Fields.Title, sector.Fields.Name)}}
		cards[sector.Fields.Name] = tools.Must(airtableCli.ListCards(paramCards))
		//		airtableCli.PrintCards(cards[sector.Fields.Name].Cards, "*")
	}
	PrintTable(sectors, categories, cards)

	// PldClient
	// You can use it to build a pdf.
	//PDF
	cli := tools.Must(pld.NewClient())
	HeaderFooter(cli)
	FirstPage(cli, sprints.Sprints[0].Fields.Number)
	cli.AddPage()
	cli.AddDescription("Project Log Document", "PLD Getout du sprint numéro "+strconv.Itoa(sprints.Sprints[0].Fields.Number), "Groupe Getout", "getout_2025@labeip.epitech.eu", "2025", "24 avril 2023", "1.0")
	cli.Ln(-1)
	cli.AddVersionHeader()
	cli.AddVersionRow("17/07/2023", "1.0.0", "Groupe Getout", "toutes", "Première version")
	cli.AddPage()
	cli.AddCard("1.1.1", "CreateAccount", 20, "Utilisateur de la plateforme de type a et de context or of type of", "pouvoir me connecter", "I am myself\nyou are yourself\nhe is himself\nwe are ourselves\nyou are yourselves\nthey are themselves", "*definition of done*", 4, []string{"*assignee*"})
	cli.AddCard("1.1.2", "Handler", 55, "Admin", "ajouter des livres", "*description*\n*description*", "*definition of done*", 1.5, []string{"perry", "erwan"})
	cli.AddCard("1.1.3", "Info", 100, "Presse", "ajouter des pubs", "*description*\n*description*", "*definition of done*", 1, []string{"inès"})
	cli.AddPage()
	cli.AddCard("1.1.4", "Test OF Size Page", 49, "Business", "ajouter des pubs", "*description*\n*description*", "*definition of done*", 1, []string{"alexandre"})
	cli.AddPage()
	cli.AddCard("2.2.5", "Info", 0, "Presse", "ajouter des pubs", "*description*\n*description*", "*definition of done*", 1, []string{"inès"})

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

func HeaderFooter(cli *pld.Client) {
	cli.SetHeader("", "", "EPITECH INNOVATIVE PROJECT - PROJECT LOG DOCUMENT")
	cli.SetFooter("", "", "", true, false)
}

func FirstPage(cli *pld.Client, sprintNumber int) {
	cli.AddPage()
	cli.AddImage("./conf/epitech.png", 50, 160, 50)
	cli.AddTitle1("")
	cli.AddTitle1("EPITECH INNOVATIVE PROJECT")
	cli.AddTitle1("PROJECT LOG DOCUMENT")
	cli.AddTitle1("SPRINT NUMERO " + strconv.Itoa(sprintNumber))
	cli.Pdf.Ln(-1)
	cli.AddTitle1B("PROMO 2025")
}
