package pldwrapper

import (
	"errors"
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
