package main

import (
	"fmt"
	"net/url"
	"os"
	"pld-maker/internal/airtable"
	"pld-maker/internal/pld"
	"pld-maker/internal/tools"
)

func main() {
	credential := tools.Must(os.ReadFile("./conf/credential.json"))
	cli := tools.Must(pld.NewClient())
	airtableCli := tools.Must(airtable.NewClient(credential))

	//Request Sprint
	paramSprints := url.Values{"filterByFormula": {"FIND(\"In progress\", {Status})"}}
	sprints := tools.Must(airtableCli.ListSprints(&paramSprints))
	fmt.Println(sprints)
	airtableCli.PrintSprints(sprints.Sprints, "")
	if len(sprints.Sprints) < 1 {
		panic("no sprint in progress")
	} else if len(sprints.Sprints) != 1 {
		fmt.Printf("Error: %s", "multiple sprint In progress")
	}
	fmt.Println(sprints.Sprints[0].Fields.Title)
	//Request Cards
	paramCategories := url.Values{"filterByFormula": {fmt.Sprintf("AND(FIND(\"%s\",CONCATENATE(\"\",{Sprint})),FIND(\"Backend\",CONCATENATE(\"\",{Secteur})))", sprints.Sprints[0].Fields.Title)}}
	categories := tools.Must(airtableCli.ListCategories(&paramCategories))
	fmt.Println(categories)
	airtableCli.PrintCategories(categories.Categories, "")

	paramCards := url.Values{"filterByFormula": {fmt.Sprintf("AND(FIND(\"%s\",CONCATENATE(\"\",{Sprint})),FIND(\"Backend\",CONCATENATE(\"\",{Secteur})))", sprints.Sprints[0].Fields.Title)}}
	cards := tools.Must(airtableCli.ListCards(&paramCards))
	fmt.Println(cards)
	airtableCli.PrintCards(cards.Cards, "")
	//	airtableCli.PrintCard(cards.Cards[1], "")
	fmt.Println(airtableCli.Token)
	fmt.Println(len("CreateAccount OF type TOOO"))

	//	cli.SetHeader("", "", "EPITECH INNOVATIVE PROJECT - PROJECT LOG DOCUMENT")
	cli.SetHeader("", "", "EPITECH INNOVATIVE PROJECT - PROJECT LOG DOCUMENT")
	cli.SetFooter("", "", "", true, false)
	//	weight
	cli.AddPage()
	cli.AddImage("./conf/epitech.png", 50, 160, 50)
	cli.AddPage()
	cli.AddCard("1.1.1", "CreateAccount", 20, "Utilisateur de la plateforme de type a et de context or of type of", "pouvoir me connecter", "I am myself\nyou are yourself\nhe is himself\nwe are ourselves\nyou are yourselves\nthey are themselves", "*definition of done*", 4, []string{"*assignee*"})
	cli.AddCard("1.1.2", "Handler", 55, "Admin", "ajouter des livres", "*description*\n*description*", "*definition of done*", 1.5, []string{"perry", "erwan"})
	cli.AddCard("1.1.3", "Info", 100, "Presse", "ajouter des pubs", "*description*\n*description*", "*definition of done*", 1, []string{"inès"})
	cli.AddPage()
	cli.AddCard("1.1.4", "Test OF Size Page", 49, "Business", "ajouter des pubs", "*description*\n*description*", "*definition of done*", 1, []string{"alexandre"})
	err := cli.OutputFileAndClose("hello.pdf")
	fmt.Println("error: ", err)
}
