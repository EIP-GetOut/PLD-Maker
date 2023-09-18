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
	currentSprints := tools.Must(dbCli.ListSprints(url.Values{"filterByFormula": {"FIND(\"In progress\", {Status})"}}))
	sprints := tools.Must(dbCli.ListSprints(nil))
	sectors := tools.Must(dbCli.ListSectors(nil))
	categories := tools.Must(dbCli.ListCategories(nil))
	cards := tools.Must(dbCli.ListCards(nil))

	//Pld  d
	pldCli.NewFile("2025_PLD_GETOUT")
	pldCli.Header("", "", "EPITECH INNOVATIVE PROJECT - PROJECT LOG DOCUMENT")
	pldCli.Footer("", "", "", &pdf.FooterParams{PageNo: true, FirstPageNo: false})

	//FirstPage
	pldCli.FirstPage("../conf/epitech.png", "EPITECH INNOVATIVE PROJECT\n\nPROJECT LOG DOCUMENT\n\nSPRINT NUMERO "+tools.Ternary(len(currentSprints) == 1, strconv.Itoa(currentSprints[0].Number), "???"), "Promo 2025")
	//Description
	pldCli.Description("Project Log Document", "PLD Getout du sprint numéro "+tools.Ternary(len(currentSprints) == 1, strconv.Itoa(currentSprints[0].Number), "???"), "Groupe Getout", "getout_2025@labeip.epitech.eu", "2025", "24 avril 2023", versions[len(versions)-1].Version)
	//Version
	pldCli.Versions(versions)
	//Summary
	pldCli.Summary(versions, schemas, sprints, sectors, categories, cards)
	pldCli.Schemas(schemas)
	//DeliveryCards
	pldCli.ListCards(sprints, sectors, categories, cards)
	//UserStories
	pldCli.Cards(sprints, sectors, categories, cards)
	//Report
	pldCli.Report(reports)
	pldCli.CloseFile()

	// pldCli := pld.Client(tools.Must(epipld.NewClient()))
	/*pdfCli.NewFile("hello")
	pdfCli.Header("", "", "EPITECH INNOVATIVE PROJECT - PROJECT LOG DOCUMENT")
	pdfCli.Footer("", "", "", &pdf.FooterParams{PageNo: true, FirstPageNo: false})
	pdfCli.NewPage()
	//	pdfCli
	pdfCli.Title(pdf.Text{Data: "Title"})
	pdfCli.SubTitle(pdf.Text{Data: "Title"})
	pdfCli.Heading1(pdf.Text{Data: "Title"})
	pdfCli.Heading2(pdf.Text{Data: "Title"})
	pdfCli.Text(pdf.Text{Data: "Text: loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum. loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum. loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum. loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum. loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum loreal ipsum."})
	pdfCli.Image(pdf.Image{Filepath: "../conf/epitech.png", Width: 150, Height: 50, Params: &pdf.ImageParams{X: 0.5, XPercent: true}})
	table := pdf.Table{
		Rows: []pdf.Row{
			{
				Cells: []pdf.Cell{
					{Str: "33.3", Percent: 33.3, Params: nil},
					{Str: "33.3", Percent: 33.3, Params: nil},
					{Str: "33.4", Percent: 33.4, Params: nil},
				},
				Params: &pdf.RowParams{
					CellParams: &pdf.CellParams{
						Background: &pdf.Color{R: 0, G: 200, B: 0},
					},
				},
			}, {
				Cells: []pdf.Cell{
					{Str: "20", Percent: 20, Params: &pdf.CellParams{TextColor: &pdf.Color{R: 100, G: 255, B: 100}}},
					{Str: "20", Percent: 20, Params: nil},
					{Str: "60", Percent: 60, Params: nil},
				},
				Params: &pdf.RowParams{
					CellParams: &pdf.CellParams{
						Bold:      true,
						TextColor: &pdf.Color{R: 255, G: 100, B: 100},
					},
					RowHeight: 20,
				},
			},
		},
		Params: &pdf.TableParams{
			RowParams: &pdf.RowParams{
				RowHeight: 10,
				CellParams: &pdf.CellParams{
					Background: &pdf.Color{R: 255, G: 0, B: 0},
				},
			},
		},
	}
	pdfCli.Table(table)
	pdfCli.CloseFile()

	// a := tools.Tuple[int]([]int{1, 2, 3})
	// fmt.Println(a)*/
}
