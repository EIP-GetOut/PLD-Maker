package fpdfwrapper

import (
	"pld-maker/v1/internal/interface/pdf"
	"strconv"
)

func (cli *Client) Header(left, center, right string) {
	fn := func() {
		//Default Value

		cli.Pdf.SetXY((cli.Width-cli.HeaderWidth)/2, 10)
		cli.Pdf.SetFont("Arial", "", 12)
		cli.Pdf.SetTextColor(0, 0, 0)
		cli.Pdf.SetFillColor(255, 255, 255)

		cli.Pdf.SetDrawColor(0, 0, 0)
		cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, left, "", 0, "LM", false, 0, "")
		cli.Pdf.SetDrawColor(0, 0, 0)
		cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, center, "", 0, "CM", false, 0, "")
		cli.Pdf.SetDrawColor(0, 0, 0)
		cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, right, "", 0, "RM", false, 0, "")

		cli.Pdf.Ln(10)
	}
	cli.Pdf.SetHeaderFunc(fn)
}

func (cli *Client) Footer(left, center, right string, footerParams *pdf.FooterParams) {
	fn := func() {
		cli.Pdf.SetXY((cli.Width-cli.HeaderWidth)/2, cli.Height-10)
		cli.Pdf.SetFont("Arial", "", 12)

		cli.Pdf.SetTextColor(0, 0, 0)
		cli.Pdf.SetDrawColor(0, 0, 0)
		cli.Pdf.SetFillColor(255, 255, 255)
		cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, left, "", 0, "LM", true, 0, "")
		cli.Pdf.SetDrawColor(0, 0, 0)
		cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, center, "", 0, "CM", true, 0, "")
		if footerParams.PageNo && (cli.Pdf.PageNo() != 1 || footerParams.FirstPageNo) {
			cli.Pdf.SetDrawColor(0, 0, 0)
			cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, right+strconv.Itoa(cli.Pdf.PageNo()), "", 0, "RM", true, 0, "")
		} else {
			cli.Pdf.SetDrawColor(0, 0, 0)
			cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, right, "", 0, "RM", true, 0, "")
		}
	}
	cli.Pdf.SetFooterFunc(fn)
}
