package pld

func (cli *Client) OutputFileAndClose(fileStr string) error {
	return cli.Pdf.OutputFileAndClose(fileStr)
}
