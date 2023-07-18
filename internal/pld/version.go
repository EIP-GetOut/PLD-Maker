package pld

type Version struct {
	Date     string
	Version  string
	Author   string
	Sections string
	Comments string
}

func (cli *Client) AddVersions(versions ...Version) {
	cli.addVersionHeader()
	for _, item := range versions {
		cli.addVersionRow(item.Date, item.Version, item.Author, item.Sections, item.Comments)
	}
}

func (cli *Client) addVersionHeader() {
	tr := cli.UnicodeTranslatorFromDescriptor("")

	cli.Pdf.SetDrawColor(255, 255, 255)
	cli.Pdf.SetTextColor(0, 0, 0)
	cli.Pdf.SetFont("Arial", "", 10)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	for i, item := range []string{"Date", "Version", "Auteur", "Section(s)", "Commentaires"} {
		x := cli.Pdf.GetX()
		y := cli.Pdf.GetY()
		cli.Pdf.SetFillColor(60, 120, 216)
		cli.Pdf.SetTextColor(255, 255, 255)
		if i == 0 {
			cli.Pdf.MultiCell(30-1, 7, tr(item), "1", "", true)
		} else if i == 1 {
			cli.Pdf.MultiCell(20-1, 7, tr(item), "1", "", true)
		} else {
			cli.Pdf.MultiCell(((cli.CardWith-50)/3)-1, 7, tr(item), "1", "", true)
		}

		if i < 4 {
			if i == 0 {
				cli.Pdf.SetXY(x+30-1, y)
			} else if i == 1 {
				cli.Pdf.SetXY(x+20-1, y)
			} else {
				cli.Pdf.SetXY(x+((cli.CardWith-50)/3)-1, y)
			}
		}

	}
}

func (cli *Client) addVersionRow(date, version, author, sections, comments string) {
	tr := cli.UnicodeTranslatorFromDescriptor("")

	cli.Pdf.SetDrawColor(255, 255, 255)
	cli.Pdf.SetFillColor(164, 194, 244)
	cli.Pdf.SetTextColor(0, 0, 0)
	cli.Pdf.SetFont("Arial", "", 8)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)

	for i, item := range []string{date, version, author, sections, comments} {
		x := cli.Pdf.GetX()
		y := cli.Pdf.GetY()
		if i == 0 {
			cli.Pdf.MultiCell(30-1, 7, tr(item), "1", "", true)
		} else if i == 1 {
			cli.Pdf.MultiCell(20-1, 7, tr(item), "1", "", true)
		} else {
			cli.Pdf.MultiCell(((cli.CardWith-50)/3)-1, 7, tr(item), "1", "", true)
		}
		if i < 4 {
			if i == 0 {
				cli.Pdf.SetXY(x+30-1, y)
			} else if i == 1 {
				cli.Pdf.SetXY(x+20-1, y)
			} else {
				cli.Pdf.SetXY(x+((cli.CardWith-50)/3)-1, y)
			}
		}

	}
}
