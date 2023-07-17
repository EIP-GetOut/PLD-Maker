package pld

import "strings"

type keyValue struct {
	key   string
	value string
}

func (cli *Client) AddDescription(title, object, author, e_mail, promo, last_update, version string) {
	tr := cli.UnicodeTranslatorFromDescriptor("")

	cli.Pdf.SetDrawColor(255, 255, 255)
	cli.Pdf.SetTextColor(0, 0, 0)
	//	cli.Pdf.SetFontSize(8)
	cli.Pdf.SetFont("Arial", "B", 10)
	for i, item := range []keyValue{{key: "Titre", value: title}, {key: "Objet", value: object}, {key: "Auteur", value: author}, {key: "E-mail", value: e_mail}, {key: "Promo", value: promo}, {key: "Mise à jour", value: last_update}, {key: "Version", value: version}} {
		item.value = WrapText(item.value, 60)
		item.key += strings.Repeat("\n ", strings.Count(item.value, "\n"))
		for j, str := range []string{tr(item.key), tr(item.value)} {
			x := cli.Pdf.GetX()
			y := cli.Pdf.GetY()
			if j < 1 {
				cli.Pdf.SetFillColor(60, 120, 216)
				cli.Pdf.SetTextColor(255, 255, 255)
				cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
				cli.Pdf.MultiCell(cli.CardWith/4, 7, str, "1", "", true)
				cli.Pdf.SetXY(x+(cli.CardWith/4), y)
			} else if i%2 == 0 {
				cli.Pdf.SetFillColor(164, 194, 244)
				cli.Pdf.SetTextColor(0, 0, 0)
				cli.Pdf.MultiCell((cli.CardWith/4)*3, 7, str, "1", "", true)
			} else if i%2 == 1 {
				cli.Pdf.SetFillColor(201, 218, 248)
				cli.Pdf.SetTextColor(0, 0, 0)
				cli.Pdf.MultiCell((cli.CardWith/4)*3, 7, str, "1", "", true)
			}
		}
	}
}
