package pld

import "strconv"

func (cli *Client) AddSummary(summary int, schema int, deliveriesCard int, sectors map[string]int, userStories map[string]int, advanceReport int) {
	var y float64
	tr := cli.UnicodeTranslatorFromDescriptor("")
	tmpPage := cli.Pdf.PageNo()

	cli.AddPage()
	tmpPage++

	cli.Pdf.SetTextColor(0, 0, 0) //black
	cli.Pdf.SetFont("Arial", "B", 10)
	// Sommaire
	y = cli.Pdf.GetY()
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, "Sommaire", "1", "", false)
	cli.Pdf.SetXY(((cli.Width-cli.CardWith)/2)+cli.CardWith-20, y)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, strconv.Itoa(tmpPage), "1", "", false)

	//Shéma fonctionnel
	tmpPage += summary
	y = cli.Pdf.GetY()
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, tr("1. Schéma Fonctionnel"), "1", "", false)
	cli.Pdf.SetXY(((cli.Width-cli.CardWith)/2)+cli.CardWith-20, y)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, strconv.Itoa(tmpPage), "1", "", false)

	//Shéma fonctionnel
	tmpPage += schema
	y = cli.Pdf.GetY()
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, tr("1. Schéma Fonctionnel"), "1", "", false)
	cli.Pdf.SetXY(((cli.Width-cli.CardWith)/2)+cli.CardWith-20, y)
	cli.Pdf.MultiCell(cli.CardWith-20, 7, strconv.Itoa(tmpPage), "1", "", false)

	// cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	// cli.Pdf.MultiCell(cli.CardWith, 7, tr("1. Schéma Fonctionnel"), "1", "", false)
	//
	// cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	// cli.Pdf.MultiCell(cli.CardWith, 7, "2. Carte des livrables", "1", "", false)
	//
	// //--------------------- Realm ---------------------
	// cli.Pdf.SetTextColor(60, 120, 216) //blue
	// cli.Pdf.SetFont("Arial", "", 14)
	// cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	// cli.Pdf.MultiCell(cli.CardWith, 7, "Sommaire", "1", "", false)
	//
	// //Shéma fonctionnel
	// cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	// cli.Pdf.MultiCell(cli.CardWith, 7, tr("1. Schéma Fonctionnel"), "1", "", false)
	//
	// cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	// cli.Pdf.MultiCell(cli.CardWith, 7, "2. Carte des livrables", "1", "", false)
}
