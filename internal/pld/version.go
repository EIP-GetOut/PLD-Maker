package pld

import (
	"pld-maker/internal/tools"
	"strings"
)

type Version struct {
	Date     string
	Version  string
	Author   string
	Sections string
	Comments string
}

func (cli *Client) AddVersions(versions ...Version) {
	tr := cli.UnicodeTranslatorFromDescriptor("")

	cli.AddPage()
	cli.Pdf.SetTextColor(0, 0, 0)
	cli.Pdf.SetFont("Arial", "B", 10)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	cli.Pdf.MultiCell(cli.CardWith, 7, tr("Tableau des révisions"), "1", "", false)
	cli.addVersionHeader()
	for _, item := range versions {
		if cli.Pdf.GetY() > cli.Height-50 {
			cli.AddPage()
		}
		cli.addVersionRow(item.Date, item.Version, item.Author, item.Sections, item.Comments)
	}
}

func (cli *Client) addVersionHeader() {
	tr := cli.UnicodeTranslatorFromDescriptor("")

	cli.Pdf.SetDrawColor(255, 255, 255)
	cli.Pdf.SetTextColor(0, 0, 0)
	cli.Pdf.SetFont("Arial", "", 10)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	for i, item := range []string{"Date (yyyy/mm/dd)", "Version", "Auteur", "Section(s)", "Commentaires"} {
		x := cli.Pdf.GetX()
		y := cli.Pdf.GetY()
		cli.Pdf.SetFillColor(60, 120, 216)
		cli.Pdf.SetTextColor(255, 255, 255)
		if i == 0 {
			cli.Pdf.MultiCell(35-1, 7, tr(item), "1", "", true)
		} else if i == 1 {
			cli.Pdf.MultiCell(20-1, 7, tr(item), "1", "", true)
		} else {
			cli.Pdf.MultiCell(((cli.CardWith-55)/3)-1, 7, tr(item), "1", "", true)
		}

		if i < 4 {
			if i == 0 {
				cli.Pdf.SetXY(x+35-1, y)
			} else if i == 1 {
				cli.Pdf.SetXY(x+20-1, y)
			} else {
				cli.Pdf.SetXY(x+((cli.CardWith-55)/3)-1, y)
			}
		}

	}
}

func (cli *Client) addVersionRow(date, version, author, sections, comments string) {
	tr := cli.UnicodeTranslatorFromDescriptor("")

	// WrapText
	author = WrapText(author, 28)
	sections = WrapText(sections, 28)
	comments = WrapText(comments, 28)

	// Diff
	authorCount := strings.Count(author, "\n")
	sectionsCount := strings.Count(sections, "\n")
	commentsCount := strings.Count(comments, "\n")
	maxCount := tools.Max(authorCount, sectionsCount, commentsCount)

	// Calcul
	date += strings.Repeat("\n ", maxCount)
	version += strings.Repeat("\n ", maxCount)
	if authorCount < maxCount {
		author += strings.Repeat("\n ", maxCount-authorCount)
	}
	if sectionsCount < maxCount {
		sections += strings.Repeat("\n ", maxCount-sectionsCount)
	}
	if commentsCount < maxCount {
		comments += strings.Repeat("\n ", maxCount-commentsCount)
	}
	//
	cli.Pdf.SetDrawColor(255, 255, 255)
	cli.Pdf.SetFillColor(164, 194, 244)
	cli.Pdf.SetTextColor(0, 0, 0)
	cli.Pdf.SetFont("Arial", "", 8)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)

	for i, item := range []string{date, version, author, sections, comments} {
		x := cli.Pdf.GetX()
		y := cli.Pdf.GetY()
		if i == 0 {
			cli.Pdf.MultiCell(35-1, 7, tr(item), "1", "", true)
		} else if i == 1 {
			cli.Pdf.MultiCell(20-1, 7, tr(item), "1", "", true)
		} else {
			cli.Pdf.MultiCell(((cli.CardWith-55)/3)-1, 7, tr(item), "1", "", true)
		}
		if i < 4 {
			if i == 0 {
				cli.Pdf.SetXY(x+35-1, y)
			} else if i == 1 {
				cli.Pdf.SetXY(x+20-1, y)
			} else {
				cli.Pdf.SetXY(x+((cli.CardWith-55)/3)-1, y)
			}
		}

	}
}
