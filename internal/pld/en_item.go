package pld

func (cli *Client) AddPage() {
	cli.Pdf.AddPage()
}

func (cli *Client) CellFormat(w, h float64, txtStr, borderStr string, ln int, alignStr string, fill bool, link int, linkStr string) {
	cli.Pdf.CellFormat(w, h, txtStr, borderStr, ln, alignStr, fill, link, linkStr)
}

func (cli *Client) MultiCell(w, h float64, txtStr, borderStr, alignStr string, fill bool) {
	cli.Pdf.MultiCell(w, h, txtStr, borderStr, alignStr, fill)
}
