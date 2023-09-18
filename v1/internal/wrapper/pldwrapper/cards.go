package pldwrapper

import (
	"pld-maker/v1/internal/interface/db"
	"pld-maker/v1/internal/interface/pdf"
)

// List Cards by Categories and by Sectors
func (cli *Client) ListCards(sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card) {
	(*cli.PdfClient).NewPage()
	(*cli.PdfClient).Text(pdf.Text{Data: "2. Cartes des livrables:", Params: &pdf.TextParams{Bold: true}})
	(*cli.PdfClient).NewLine()
}

// Show Cards
func (cli *Client) Cards(sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card) {
	(*cli.PdfClient).NewPage()
	(*cli.PdfClient).Text(pdf.Text{Data: "3. User Stories:", Params: &pdf.TextParams{Bold: true}})
	(*cli.PdfClient).NewLine()
}
