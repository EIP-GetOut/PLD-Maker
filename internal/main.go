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

	//	countryList := make([]countryType, 0, 8)
	cli.AddPage()
	//	cli.Pdf.RegisterImageOptions()
	cli.SetFont("Arial", "", 10)
	cli.AddCard("1.1.1", "CreateAccount", 20, "Utilisateur de la plateforme de type a et de\ncontext", "pouvoir me connecter", "I am myself\nyou are yourself\nhe is himself\nwe are ourselves\nyou are yourselves\nthey are themselves", "*definition of done*", 4, "*assignee*")
	cli.AddCard("1.1.2", "Handler", 55, "Admin", "ajouter des livres", "*description*\n*description*", "*definition of done*", 1.5, "*assignee*")
	cli.AddCard("1.1.3", "Info", 100, "Presse", "ajouter des pubs", "*description*\n*description*", "*definition of done*", 1, "*assignee*")
	cli.OutputFileAndClose("hello.pdf")
}
