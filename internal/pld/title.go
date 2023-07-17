package pld

func (cli *Client) AddTitle1(str string) {
	cli.Pdf.SetFillColor(255, 255, 255)
	cli.Pdf.SetDrawColor(255, 255, 255)
	cli.Pdf.SetTextColor(17, 85, 204)
	cli.Pdf.SetFontSize(25)
	cli.Pdf.SetFontStyle("")
	cli.Pdf.SetX((cli.Width - cli.TitleWith) / 2)
	cli.Pdf.MultiCell(cli.TitleWith, 30, str, "1", "CM", false)
	cli.Pdf.SetY(cli.Pdf.GetY() - 10)
}

func (cli *Client) AddTitle1B(str string) {
	cli.Pdf.SetFillColor(255, 255, 255)
	cli.Pdf.SetDrawColor(255, 255, 255)
	cli.Pdf.SetTextColor(17, 85, 204)
	cli.Pdf.SetFontSize(25)
	cli.Pdf.SetFontStyle("B")
	cli.Pdf.SetX((cli.Width - cli.TitleWith) / 2)
	cli.Pdf.MultiCell(cli.TitleWith, 30, str, "1", "CM", false)
	cli.Pdf.SetY(cli.Pdf.GetY() - 10)
}
