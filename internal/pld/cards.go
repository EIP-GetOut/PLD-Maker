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
	//style
	div := 100 / (len(cli.PercentColors) - 1)
	color := cli.PercentColors[progress/div]

	//row_1
	cli.SetFillColor(color.R, color.G, color.B)
	cli.Pdf.SetX((cli.Width - 3*50) / 2)
	for _, str := range []string{number + " " + title, "Progres", strconv.Itoa(progress) + " %"} {
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
	for i, str := range []string{asWho, iWant} {
		x := cli.GetX()
		y := cli.GetY()
		cli.MultiCell(75, 5, str, "1", "", true)
		if i == 0 {
			cli.SetXY(x+75, y)
		}
	}

	//row_4-7
	for i, str := range []string{"Description", description, "Definition of done", definitionOfDone} {
		cli.Pdf.SetX((cli.Width - 150) / 2)
		if i%2 == 0 {
			cli.SetFillColor(color.R, color.G, color.B)
			//cli.Pdf.CellFormat(150, 5, str, "1", 0, "", true, 0, "")
			cli.Pdf.MultiCell(150, 5, str, "1", "", true)
		} else {
			cli.SetFillColor(255, 255, 255)
			cli.MultiCell(150, 5, str+"\n ", "1", "", true)
		}
	}

	//row_8
	cli.SetFillColor(color.R, color.G, color.B)
	cli.Pdf.SetX((cli.Width - 75*2) / 2)
	for i, str := range []string{"Charge Estimée (J/H) :", strconv.FormatFloat(jh, 'f', -1, 64), "Assignés (J/H) :", assignee} {
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
