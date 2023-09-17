package fpdfwrapper

import (
	"pld-maker/v1/internal/interface/pdf"
	"strconv"
)

func (cli *Client) Header(left, center, right string) {
	fn := func() {
		//Default Value

		cli.pdf.SetXY((cli.Width-cli.HeaderWidth)/2, 10)
		cli.pdf.SetFont("Arial", "", 12)
		cli.pdf.SetTextColor(0, 0, 0)
		cli.pdf.SetFillColor(255, 255, 255)

		cli.pdf.SetDrawColor(0, 0, 0)
		cli.pdf.CellFormat(cli.HeaderWidth/3, 1, left, "", 0, "LM", false, 0, "")
		cli.pdf.SetDrawColor(0, 0, 0)
		cli.pdf.CellFormat(cli.HeaderWidth/3, 1, center, "", 0, "CM", false, 0, "")
		cli.pdf.SetDrawColor(0, 0, 0)
		cli.pdf.CellFormat(cli.HeaderWidth/3, 1, right, "", 0, "RM", false, 0, "")

		cli.pdf.Ln(10)
	}
	cli.pdf.SetHeaderFunc(fn)
}

func (cli *Client) Footer(left, center, right string, footerParams *pdf.FooterParams) {
	fn := func() {
		cli.pdf.SetXY((cli.Width-cli.HeaderWidth)/2, cli.Height-10)
		cli.pdf.SetFont("Arial", "", 12)

		cli.pdf.SetTextColor(0, 0, 0)
		cli.pdf.SetDrawColor(0, 0, 0)
		cli.pdf.SetFillColor(255, 255, 255)
		cli.pdf.CellFormat(cli.HeaderWidth/3, 1, cli.translator(left), "", 0, "LM", true, 0, "")
		cli.pdf.SetDrawColor(0, 0, 0)
		cli.pdf.CellFormat(cli.HeaderWidth/3, 1, cli.translator(center), "", 0, "CM", true, 0, "")
		if footerParams.PageNo && (cli.pdf.PageNo() != 1 || footerParams.FirstPageNo) {
			cli.pdf.SetDrawColor(0, 0, 0)
			cli.pdf.CellFormat(cli.HeaderWidth/3, 1, cli.translator(right+strconv.Itoa(cli.pdf.PageNo())), "", 0, "RM", true, 0, "")
		} else {
			cli.pdf.SetDrawColor(0, 0, 0)
			cli.pdf.CellFormat(cli.HeaderWidth/3, 1, right, "", 0, "RM", true, 0, "")
		}
	}
	cli.pdf.SetFooterFunc(fn)
}
