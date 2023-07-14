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
	Assignee         []string
}

//45 char

func (cli *Client) AddCards(cards ...Cards) {
	for _, it := range cards {
		cli.AddCard(it.Number, it.Title, it.Progress, it.AsWho, it.IWant, it.Description, it.DefinitionOfDone, it.Jh, it.Assignee)
	}
}

// Get beginingX
// Get beginingPage
// Get endingPage
// if  beginingPage != endingPage
// setX as beginingX
// Add Page
func (cli *Client) AddCard(number string, title string, progress int, asWho string, iWant string, description string, definitionOfDone string, jh float64, assignee []string) {
	cli.SetDrawColor(0, 0, 0)
	cli.SetTextColor(0, 0, 0)

	//style
	div := 100 / (len(cli.PercentColors) - 1)
	color := cli.PercentColors[progress/div]

	//row_1
	cli.addCardsRow1(color, number, title, progress)
	//row_2
	cli.addCardsRow2(color)
	//row_3
	cli.addCardsRow3(color, asWho, iWant)
	//row_4-7
	cli.addCardsRow4_7(color, description, definitionOfDone)
	//row_8
	cli.addCardsRow8(color, jh, assignee)
	//end
	cli.Pdf.Ln(5)
}

func (cli *Client) addCardsRow1(color Color, number, title string, progress int) {
	tr := cli.UnicodeTranslatorFromDescriptor("")

	cli.SetFillColor(color.R, color.G, color.B)
	cli.SetX((cli.Width - 3*50) / 2)
	for _, str := range []string{number + " " + tr(title), "Progres", strconv.Itoa(progress) + " %"} {
		cli.Pdf.CellFormat(50, 10, str, "1", 0, "", true, 0, "")
	}
	cli.Pdf.Ln(-1)

}

func (cli *Client) addCardsRow2(color Color) {
	//tr := cli.UnicodeTranslatorFromDescriptor("")

	cli.SetFillColor(color.R, color.G, color.B)
	cli.SetX((cli.Width - 2*75) / 2)
	for _, str := range []string{"en tant que:", "je veux:"} {
		cli.CellFormat(75, 5, str, "1", 0, "", true, 0, "")
	}
	cli.Pdf.Ln(-1)

}

func (cli *Client) addCardsRow3(color Color, asWho, iWant string) {
	tr := cli.UnicodeTranslatorFromDescriptor("")

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

}

func (cli *Client) addCardsRow4_7(color Color, description, definitionOfDone string) {
	tr := cli.UnicodeTranslatorFromDescriptor("")

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
}

func (cli *Client) addCardsRow8(color Color, jh float64, assignee []string) {
	tr := cli.UnicodeTranslatorFromDescriptor("")

	cli.Pdf.SetX((cli.Width - 75*2) / 2)
	for i, str := range []string{tr("Charge Estimée (J/H) :"), strconv.FormatFloat(jh, 'f', -1, 64), tr("Assignés (J/H) :"), tr(strings.Join(assignee, ", "))} {
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
}
