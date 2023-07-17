package pld

type keyValue struct {
	key   string
	value string
}

func (cli *Client) AddDescription(title, object, author, e_mail, promo, last_update, version string) {
	tr := cli.UnicodeTranslatorFromDescriptor("")

	cli.Pdf.SetDrawColor(255, 255, 255)
	cli.Pdf.SetTextColor(0, 0, 0)
	//	cli.Pdf.SetFontSize(8)
	cli.Pdf.SetFont("Arial", "", 8)
	for _, item := range []keyValue{{key: "Titre", value: title}, {key: "Objet", value: object}, {key: "Auteur", value: author}, {key: "E-mail", value: e_mail}, {key: "Promo", value: promo}, {key: "Mise à jour", value: last_update}, {key: "Version", value: version}} {
		for i, str := range []string{tr(item.key), tr(item.value)} {
			x := cli.Pdf.GetX()
			y := cli.Pdf.GetY()
			if i < 1 {
				cli.Pdf.SetFillColor(60, 120, 216)
				cli.Pdf.SetTextColor(255, 255, 255)
				cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
				cli.Pdf.MultiCell(cli.CardWith/4, 5, str, "1", "", true)
				cli.Pdf.SetXY(x+(cli.CardWith/4), y)
			} else {
				cli.Pdf.SetFillColor(164, 194, 244)
				cli.Pdf.SetTextColor(0, 0, 0)
				cli.Pdf.MultiCell((cli.CardWith/4)*3, 5, str, "1", "", true)
			}
		}
	}
}
