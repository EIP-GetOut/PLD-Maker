package pldwrapper

import (
	"errors"
	"pld-maker/v1/internal/interface/db"
	"pld-maker/v1/internal/interface/pdf"
)

type Client struct {
	PdfClient     *pdf.Client
	PercentColors []pdf.Color
}

func NewClient(pdfCli *pdf.Client) (*Client, error) {
	if pdfCli != nil {
		return &Client{PdfClient: pdfCli}, nil
	} else {
		return nil, errors.New("error: pdfCli *pdf.Client is nil")
	}
}
func (cli *Client) NewFile(str string) {

	cli.PercentColors = []pdf.Color{
		{R: 234, G: 153, B: 153}, // Red
		{R: 255, G: 229, B: 153}, // Yellow
		{R: 182, G: 215, B: 168}, // Green
		{R: 200, G: 200, B: 200}, // Gray
	}
	(*cli.PdfClient).NewFile(str)
}

func (cli *Client) CloseFile() {
	(*cli.PdfClient).CloseFile()
}

// FirstPage
func (cli *Client) FirstPage(imageFilepath, title, lowTitle string) {

}

// Description
func (cli *Client) Description(title, object, author, e_mail, promo, last_update, version string) {

}

// Version
func (cli *Client) Versions([]db.Version) {
}

// Summary
func (cli *Client) Summary(versions []db.Version, schemas []db.Schema, sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card) {

}

// Display Schema Images
func (cli *Client) Schemas(schemas []db.Schema) {

}

// List Cards by Categories and by Sectors
func (cli *Client) ListCards(sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card) {

}

// Show Cards
func (cli *Client) Cards(sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card) {

}

// Show reports
func (cli *Client) Report([]db.Report) {

}
