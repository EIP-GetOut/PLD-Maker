package epitech

import "pld-maker/internal/pld"

func HeaderFooter(cli *pld.Client) {
	cli.SetHeader("", "", "EPITECH INNOVATIVE PROJECT - PROJECT LOG DOCUMENT")
	cli.SetFooter("", "", "", true, false)
}
