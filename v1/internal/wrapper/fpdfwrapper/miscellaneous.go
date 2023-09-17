package fpdfwrapper

func (cli *Client) UnicodeTranslator(str string) string {
	return cli.translator(str)
}

func (cli *Client) PageNo() int {
	return cli.pdf.PageNo()
}
