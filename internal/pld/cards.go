package pld

import (
	"strconv"
	"strings"
)

type Cards struct {
	Number           string
	Title            string
	Progress         int
	AsWho            string
	IWant            string
	Description      string
	DefinitionOfDone string
	Jh               float64
	Assignee         string
}

//45 char

func (cli *Client) AddCards(cards ...Cards) {
	for _, it := range cards {
		cli.AddCard(it.Number, it.Title, it.Progress, it.AsWho, it.IWant, it.Description, it.DefinitionOfDone, it.Jh, it.Assignee)
	}
}

func (cli *Client) AddCard(number string, title string, progress int, asWho string, iWant string, description string, definitionOfDone string, jh float64, assignee string) {
	cli.SetDrawColor(0, 0, 0)
	cli.SetTextColor(0, 0, 0)

	//style
	div := 100 / (len(cli.PercentColors) - 1)
	color := cli.PercentColors[progress/div]
	tr := cli.Pdf.UnicodeTranslatorFromDescriptor("")

	//row_1
	cli.SetFillColor(color.R, color.G, color.B)
	cli.Pdf.SetX((cli.Width - 3*50) / 2)
	for _, str := range []string{number + " " + tr(title), "Progres", strconv.Itoa(progress) + " %"} {
		cli.Pdf.CellFormat(50, 10, str, "1", 0, "", true, 0, "")
	}
	cli.Pdf.Ln(-1)

	//row_2
	cli.Pdf.SetX((cli.Width - 75*2) / 2)
	for _, str := range []string{"en tant que:", "je veux:"} {
		cli.CellFormat(75, 5, str, "1", 0, "", true, 0, "")
	}
	cli.Pdf.Ln(-1)

	//row_3
	//add lines to keep cards rect
	diffLen := strings.Count(asWho, "\n") - strings.Count(iWant, "\n")
	if diffLen < 0 {
		asWho += strings.Repeat("\n ", -diffLen)
	} else {
		iWant += strings.Repeat("\n ", diffLen)
	}
	//display lines
	cli.SetFillColor(255, 255, 255)
	cli.SetX((cli.Width - 75*2) / 2)
	for i, str := range []string{tr(asWho), tr(iWant)} {
		x := cli.GetX()
		y := cli.GetY()
		cli.MultiCell(75, 5, str, "1", "", true)
		if i == 0 {
			cli.SetXY(x+75, y)
		}
	}

	//row_4-7
	for i, str := range []string{"Description", tr(description), "Definition of done", tr(definitionOfDone)} {
		cli.Pdf.SetX((cli.Width - 150) / 2)
		cli.SetFillColor(color.R, color.G, color.B)
		if i%2 == 0 {
			//cli.Pdf.CellFormat(150, 5, str, "1", 0, "", true, 0, "")
			cli.Pdf.MultiCell(150, 5, str, "1", "", true)
		} else {
			cli.MultiCell(150, 5, str+"\n ", "1", "", false)
		}
	}

	//row_8
	cli.Pdf.SetX((cli.Width - 75*2) / 2)
	for i, str := range []string{tr("Charge Estimée (J/H) :"), strconv.FormatFloat(jh, 'f', -1, 64), tr("Assignés (J/H) :"), tr(assignee)} {
		if i%2 == 0 {
			cli.SetFillColor(color.R, color.G, color.B)
		} else {
			cli.SetFillColor(255, 255, 255)
		}
		if i < 1 || i > 2 {
			cli.CellFormat(50, 5, str, "1", 0, "", true, 0, "")
		} else {
			cli.CellFormat(25, 5, str, "1", 0, "", true, 0, "")
		}
	}
	cli.Pdf.Ln(-1)
	//row_4
	cli.Pdf.Ln(10)
}
