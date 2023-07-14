package pld

func (cli *Client) LinkToPage(page int) int {
	link := cli.Pdf.AddLink()
	cli.Pdf.SetLink(link, 0, link)
	return link
}

func (cli *Client) LinkToPagePosition(y float64, page int) int {
	link := cli.Pdf.AddLink()
	cli.Pdf.SetLink(link, y, page)
	return link
}
