package fpdfwrapper

func (cli *Client) UnicodeTranslator(str string) string {
	return cli.translator(str)
}
