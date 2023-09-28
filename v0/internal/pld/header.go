package pld

func (cli *Client) SetHeader(left, center, right string) {
	fn := func() {
		//Default Value

		cli.Pdf.SetXY((cli.Width-cli.HeaderWidth)/2, 10)
		cli.Pdf.SetFont("Arial", "", 12)
		cli.Pdf.SetTextColor(0, 0, 0)

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
