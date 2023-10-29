package main

import (
	"net/url"
	"os"
	"pld-maker/v1/internal/interface/db"
	"pld-maker/v1/internal/interface/pdf"
	"pld-maker/v1/internal/interface/pld"
	"strconv"

	"pld-maker/v1/internal/wrapper/airtablewrapper"
	"pld-maker/v1/internal/wrapper/fpdfwrapper"
	"pld-maker/v1/internal/wrapper/pldwrapper"

	"pld-maker/v1/internal/tools"
)

func main() {
	//dbCli := db.Client(tools.Must(airtable.NewClient(credential)))
	credential := tools.Must(os.ReadFile("./conf/credential.json"))
	dbCli := db.Client(tools.Must(airtablewrapper.NewClient(credential)))
	pdfCli := pdf.Client(tools.Must(fpdfwrapper.NewClient()))
	pldCli := pld.Client(tools.Must(pldwrapper.NewClient(&pdfCli)))

	// Database
	// VERSION, SCHEMA, REPORT
	versions := tools.Must(dbCli.ListVersions(url.Values{"filterByFormula": {""}, "sort[0][field]": {"Date"}, "sort[0][direction]": {"asc"}}))
	schemas := tools.Must(dbCli.ListSchemas(nil))
	reports := tools.Must(dbCli.ListReports(nil))
	// CARD
	currentSprints := tools.Must(dbCli.ListSprints(url.Values{"filterByFormula": {"FIND(\"In progress\", {Status})"}, "sort[0][field]": {"Number"}, "sort[0][direction]": {"desc"}}))[0]
	sprints := tools.Must(dbCli.ListSprints(url.Values{"filterByFormula": {"NOT({Status} = \"Todo\")"}}))
	sectors := tools.Must(dbCli.ListSectors(nil))
	categories := tools.Must(dbCli.ListCategories(url.Values{"filterByFormula": {"NOT(SprintStatus = 'Todo')"}}))
	cards := tools.Must(dbCli.ListCards(url.Values{"filterByFormula": {"NOT(SprintStatus = 'Todo')"}}))

	//Pld  d
	pldCli.NewFile("2025_PLD_GETOUT")
	pldCli.Header("", "", "EPITECH INNOVATIVE PROJECT - PROJECT LOG DOCUMENT")
	pldCli.Footer("", "", "", &pdf.FooterParams{PageNo: true, FirstPageNo: false})

	//FirstPage
	pldCli.FirstPage("../conf/epitech.png", "EPITECH INNOVATIVE PROJECT\n\nPROJECT LOG DOCUMENT\n\nSPRINT NUMERO "+strconv.Itoa(currentSprints.Number), "Promo 2025")
	//Description
	pldCli.Description("Project Log Document", "PLD Getout du sprint numéro "+strconv.Itoa(currentSprints.Number), "Groupe Getout", "getout_2025@labeip.epitech.eu", "2025", "24 avril 2023", versions[len(versions)-1].Version)
	//Version
	pldCli.Versions(versions)
	//Summary
	pldCli.TableOfContent(versions, schemas, sprints, sectors, categories, cards)
	pldCli.Schemas(schemas)
	//DeliveryCards
	pldCli.ListCards(currentSprints, sprints, sectors, categories, cards)
	//UserStories
	pldCli.Cards(currentSprints, sprints, sectors, categories, cards)
	//Report
	pldCli.Report(reports)
	pldCli.CloseFile()
}
