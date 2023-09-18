package pldwrapper

import "pld-maker/v1/internal/interface/pdf"

func (cli *Client) Header(left, center, right string) {
	(*cli.PdfClient).Header(left, center, right)
}

func (cli *Client) Footer(left, center, right string, footerParams *pdf.FooterParams) {
	(*cli.PdfClient).Footer(left, center, right, footerParams)
}
