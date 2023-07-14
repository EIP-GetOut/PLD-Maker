package pld

func (cli *Client) SetAutoPageBreak(auto bool, margin float64) {
	cli.Pdf.SetAutoPageBreak(auto, margin)
}

func (cli *Client) Ln(h float64) {
	cli.Pdf.Ln(h)
}

func (cli *Client) OutputFileAndClose(fileStr string) error {
	return cli.Pdf.OutputFileAndClose(fileStr)
}
