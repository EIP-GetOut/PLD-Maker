package pld

import (
	"strconv"
)

func (cli *Client) SetFooter(left, center, right string, pageNo bool, firstPageNo bool) {
	fn := func() {
		cli.Pdf.SetXY((cli.Width-cli.HeaderWidth)/2, cli.Height-10)
		cli.Pdf.SetFont("Arial", "", 12)

		cli.Pdf.SetTextColor(0, 0, 0)
		cli.Pdf.SetDrawColor(0, 0, 0)
		cli.Pdf.SetFillColor(255, 255, 255)
		cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, left, "", 0, "LM", true, 0, "")
		cli.Pdf.SetDrawColor(0, 0, 0)
		cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, center, "", 0, "CM", true, 0, "")
		if pageNo && (cli.Pdf.PageNo() != 1 || firstPageNo) {
			cli.Pdf.SetDrawColor(0, 0, 0)
			cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, strconv.Itoa(cli.Pdf.PageNo()), "", 0, "RM", true, 0, "")
		} else {
			cli.Pdf.SetDrawColor(0, 0, 0)
			cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, right, "", 0, "RM", true, 0, "")
		}
	}
	cli.Pdf.SetFooterFunc(fn)
}
