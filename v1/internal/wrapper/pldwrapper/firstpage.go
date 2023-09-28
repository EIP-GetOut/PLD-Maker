package pldwrapper

import "pld-maker/v1/internal/interface/pdf"

// FirstPage
func (cli *Client) FirstPage(imageFilepath, title, lowTitle string) {
	(*cli.PdfClient).NewPage()
	(*cli.PdfClient).Image(pdf.Image{Filepath: "./conf/epitech.png", Width: 150, Height: 50, Params: &pdf.ImageParams{Y: 50}})
	(*cli.PdfClient).Title(pdf.Text{})
	(*cli.PdfClient).NewLine()
	(*cli.PdfClient).NewLine()
	(*cli.PdfClient).NewLine()
	(*cli.PdfClient).Title(pdf.Text{Data: title, Params: &pdf.TextParams{Align: pdf.Center, TextColor: &pdf.Color{R: 17, G: 85, B: 204}}})
	(*cli.PdfClient).NewLine()
	(*cli.PdfClient).NewLine()
	(*cli.PdfClient).NewLine()
	(*cli.PdfClient).NewLine()
	(*cli.PdfClient).Title(pdf.Text{Data: lowTitle, Params: &pdf.TextParams{Align: pdf.Center, TextColor: &pdf.Color{R: 17, G: 85, B: 204}, Bold: true}})
}
