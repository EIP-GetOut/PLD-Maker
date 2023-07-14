package pld

import "strconv"

func (cli *Client) SetHeader(left, center, right string) {
	fn := func() {
		//Default Value

		cli.Pdf.SetXY((cli.Width-cli.HeaderWidth)/2, 10)
		cli.SetFont("Arial", "B", 12)

		cli.SetDrawColor(0, 0, 0)
		cli.SetFillColor(255, 255, 255)
		cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, left, "", 0, "LM", true, 0, "")
		cli.SetDrawColor(0, 0, 0)
		cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, center, "", 0, "CM", true, 0, "")
		cli.SetDrawColor(0, 0, 0)
		cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, right, "", 0, "RM", true, 0, "")

		cli.Ln(10)
	}
	cli.SetHeaderFunc(fn)
}

func (cli *Client) SetFooter(left, center, right string, pageNumber bool) {
	fn := func() {
		cli.Pdf.SetXY((cli.Width-cli.HeaderWidth)/2, cli.Height-10)
		cli.SetFont("Arial", "B", 12)

		cli.SetDrawColor(0, 0, 0)
		cli.SetFillColor(255, 255, 255)
		cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, left, "", 0, "LM", true, 0, "")
		cli.SetDrawColor(0, 0, 0)
		cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, center, "", 0, "CM", true, 0, "")
		cli.SetDrawColor(0, 0, 0)
		//		tmp := cli.GetX()
		if pageNumber {
			cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, strconv.Itoa(cli.Pdf.PageNo()), "", 0, "RM", true, 0, "")
		} else {
			cli.Pdf.CellFormat(cli.HeaderWidth/3, 1, right, "", 0, "RM", true, 0, "")
		}
		//cli.SetX(tmp)

	}
	cli.Pdf.SetFooterFunc(fn)
}
