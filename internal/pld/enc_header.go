package pld

func (cli *Client) SetHeaderFunc(fn func()) {
	cli.Pdf.SetHeaderFunc(fn)
}
