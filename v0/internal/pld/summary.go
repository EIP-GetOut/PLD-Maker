package pld

import (
	"fmt"
	"strconv"
)

func (cli *Client) AddSummary(summary int, schema int, arraySectors []string, deliveryCards map[string]int, userStories map[string]int) {
	var y float64
	tr := cli.UnicodeTranslatorFromDescriptor("")
	tmpPage := cli.Pdf.PageNo()

	cli.AddPage()
	tmpPage++

	cli.Pdf.SetTextColor(0, 0, 0) //black
	// Summary
	y = cli.Pdf.GetY()
	cli.Pdf.SetFont("Arial", "B", 10)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, "Sommaire", "1", "", false)
	cli.Pdf.SetXY(((cli.Width-cli.CardWith)/2)+cli.CardWith-20, y)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, strconv.Itoa(tmpPage), "1", "", false)

	//ShémaFonctionnel
	tmpPage += summary
	y = cli.Pdf.GetY()
	cli.Pdf.SetFont("Arial", "B", 10)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, tr("1. Schéma Fonctionnel"), "1", "", false)
	cli.Pdf.SetXY(((cli.Width-cli.CardWith)/2)+cli.CardWith-20, y)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, strconv.Itoa(tmpPage), "1", "", false)

	//DeliveryCards
	cli.Pdf.SetFont("Arial", "B", 10)
	tmpPage += schema
	y = cli.Pdf.GetY()
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, tr("2. Cartes des livrables"), "1", "", false)
	cli.Pdf.SetXY(((cli.Width-cli.CardWith)/2)+cli.CardWith-20, y)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, strconv.Itoa(tmpPage), "1", "", false)

	cli.Pdf.SetFont("Arial", "", 10)
	sfIdx := 1
	for _, sector := range arraySectors {
		y = cli.Pdf.GetY()
		cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
		cli.Pdf.MultiCell(cli.CardWith-20, 7, tr("2."+strconv.Itoa(sfIdx)+" "+sector), "1", "", false)
		cli.Pdf.SetXY(((cli.Width-cli.CardWith)/2)+cli.CardWith-20, y)
		cli.Pdf.MultiCell(cli.CardWith-20, 7, strconv.Itoa(tmpPage), "1", "", false)
		fmt.Println("+ ", tmpPage, deliveryCards[sector])
		tmpPage += deliveryCards[sector]
		sfIdx++
	}

	//UserStories
	y = cli.Pdf.GetY()
	cli.Pdf.SetFont("Arial", "B", 10)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, tr("3. User Stories"), "1", "", false)
	cli.Pdf.SetXY(((cli.Width-cli.CardWith)/2)+cli.CardWith-20, y)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, strconv.Itoa(tmpPage), "1", "", false)

	cli.Pdf.SetFont("Arial", "", 10)
	idx := 1
	for _, sector := range arraySectors {
		y = cli.Pdf.GetY()
		cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
		cli.Pdf.MultiCell(cli.CardWith-20, 7, tr("3."+strconv.Itoa(idx)+" "+sector), "1", "", false)
		cli.Pdf.SetXY(((cli.Width-cli.CardWith)/2)+cli.CardWith-20, y)
		cli.Pdf.MultiCell(cli.CardWith-20, 7, strconv.Itoa(tmpPage), "1", "", false)
		fmt.Println("> ", tmpPage, userStories[sector])
		tmpPage += userStories[sector]
		idx++

	}

	//ProgressReport
	y = cli.Pdf.GetY()
	cli.Pdf.SetFont("Arial", "B", 10)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, tr("4. Rapport d'avancement"), "1", "", false)
	cli.Pdf.SetXY(((cli.Width-cli.CardWith)/2)+cli.CardWith-20, y)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, strconv.Itoa(tmpPage), "1", "", false)
}
