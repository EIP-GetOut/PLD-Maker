package main

import (
	"fmt"
	"os"
	"pld-maker/internal/airtable"
	"pld-maker/internal/pld"
	"pld-maker/internal/tools"
)

func main() {
	cli := tools.Must(pld.NewClient())
	credential := tools.Must(os.ReadFile("./conf/credential.json"))

	airtableCli := tools.Must(airtable.NewClient(credential))
	fmt.Println(airtableCli.Token)

	//	cli.SetHeader("", "", "EPITECH INNOVATIVE PROJECT - PROJECT LOG DOCUMENT")
	cli.SetHeader("", "", "EPITECH INNOVATIVE PROJECT - PROJECT LOG DOCUMENT")
	cli.SetFooter("", "", "", true)
	cli.SetFont("Arial", "", 10)
	cli.SetFillColor(255, 255, 255)
	cli.SetDrawColor(255, 255, 255)
	cli.SetTextColor(0, 0, 0)
	cli.AddPage()
	cli.SetX((cli.Width - cli.HeaderWidth) / 2)
	cli.CellFormat(50, 10, "LinkToPage0", "1", 0, "", true, cli.LinkToPage(0), "")
	cli.Ln(-1)
	cli.CellFormat(50, 10, "LinkToPage1", "1", 0, "", true, cli.LinkToPage(1), "")
	cli.Ln(-1)
	cli.CellFormat(50, 10, "LinkToPage2", "1", 0, "", true, cli.LinkToPage(2), "")
	cli.AddPage()
	cli.AddCard("1.1.0", "CreateAccount OF type TOOOOOOOOO Long", 20, "Utilisateur de la plateforme de type a et de\ncontext tooooooooooooooooooooooooo lonnnnnnng", "pouvoir me connecter", "I am myself\nyou are yourself\nhe is himself\nwe are ourselves\nyou are yourselves\nthey are themselves but I'm tooooooooooooooo lonnnnnnnnnnnnnnnnng", "*definition of done*", 4, []string{"perry", "erwan", "alexandre", "inès", "théo"})
	cli.AddCard("1.1.1", "CreateAccount", 20, "Utilisateur de la plateforme de type a et de\ncontext", "pouvoir me connecter", "I am myself\nyou are yourself\nhe is himself\nwe are ourselves\nyou are yourselves\nthey are themselves", "*definition of done*", 4, []string{"*assignee*"})
	cli.AddCard("1.1.2", "Handler", 55, "Admin", "ajouter des livres", "*description*\n*description*", "*definition of done*", 1.5, []string{"perry", "erwan"})
	cli.AddCard("1.1.3", "Info", 100, "Presse", "ajouter des pubs", "*description*\n*description*", "*definition of done*", 1, []string{"inès"})
	cli.AddPage()
	cli.AddCard("1.1.4", "Test OF Size Page", 49, "Business", "ajouter des pubs", "*description*\n*description*", "*definition of done*", 1, []string{"alexandre"})
	err := cli.OutputFileAndClose("hello.pdf")
	fmt.Println("error: ", err)
}
