package pld

import (
	"pld-maker/internal/tools"
	"strconv"
	"strings"
)

type Card struct {
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

func (cli *Client) AddCards(cards ...Card) {
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
	cli.Pdf.SetDrawColor(0, 0, 0)
	cli.Pdf.SetTextColor(0, 0, 0)
	cli.Pdf.SetFontSize(8)
	//style

	div := 100 / (len(cli.PercentColors) - 1)
	color := cli.PercentColors[progress/div]

	//rows
	if cli.Pdf.GetY() > cli.Height-50 {
		cli.Pdf.AddPage()
	}
	cli.addCardsRow1(color, number, title, progress)
	cli.addCardsRow2(color)
	if cli.Pdf.GetY() > cli.Height-50 {
		cli.Pdf.AddPage()
	}
	cli.addCardsRow3(color, asWho, iWant)
	cli.addCardsRow4_7(color, description, definitionOfDone)
	if cli.Pdf.GetY() > cli.Height-50 {
		cli.Pdf.AddPage()
	}
	cli.addCardsRow8(color, jh, assignee)
	//end
	cli.Pdf.Ln(5)
}

// title should be cut when string len is over 25 without a \n
func (cli *Client) addCardsRow1(color Color, number, title string, progress int) {
	tr := cli.UnicodeTranslatorFromDescriptor("")
	//add \n
	WrapedTitle := WrapText(title, 25)
	diffLen := strings.Count(WrapedTitle, "\n")
	Progress := "Progres" + strings.Repeat("\n ", diffLen)
	ProgressValue := strconv.Itoa(progress) + " %" + strings.Repeat("\n ", diffLen)
	//display
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	for i, str := range []string{number + " " + tr(WrapedTitle), Progress, ProgressValue} {
		x, y := cli.Pdf.GetXY()
		cli.Pdf.SetFillColor(color.R, color.G, color.B)
		cli.Pdf.MultiCell(cli.CardWith/3, 5, str, "1", "", true)
		if i < 2 {
			cli.Pdf.SetXY(x+(cli.CardWith/3), y)
		}
	}
}

func (cli *Client) addCardsRow2(color Color) {
	//tr := cli.UnicodeTranslatorFromDescriptor("")

	cli.Pdf.SetFillColor(color.R, color.G, color.B)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	for _, str := range []string{"en tant que:", "je veux:"} {
		cli.Pdf.CellFormat(cli.CardWith/2, 5, str, "1", 0, "", true, 0, "")
	}
	cli.Pdf.Ln(-1)

}

// asWho should be cut when string len is over 45 without a \n
// iWant should be cut when string len is over 45 without a \n
func (cli *Client) addCardsRow3(color Color, asWho, iWant string) {
	tr := cli.UnicodeTranslatorFromDescriptor("")

	WrapedAsWho := WrapText(asWho, 60)
	WrapedIWant := WrapText(iWant, 60)
	diffLen := strings.Count(WrapedAsWho, "\n") - strings.Count(WrapedIWant, "\n")
	if diffLen < 0 {
		WrapedAsWho += strings.Repeat("\n ", -diffLen)
	} else {
		WrapedIWant += strings.Repeat("\n ", diffLen)
	}
	//display lines
	cli.Pdf.SetFillColor(255, 255, 255)
	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	for i, str := range []string{tr(WrapedAsWho), tr(WrapedIWant)} {
		x := cli.Pdf.GetX()
		y := cli.Pdf.GetY()
		cli.Pdf.MultiCell(cli.CardWith/2, 5, str, "1", "", true)
		if i < 1 {
			cli.Pdf.SetXY(x+(cli.CardWith/2), y)
		}
	}

}

func (cli *Client) addCardsRow4_7(color Color, description, definitionOfDone string) {
	tr := cli.UnicodeTranslatorFromDescriptor("")

	for i, str := range []string{"Description", tr(description), "Definition of done", tr(definitionOfDone)} {
		cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
		cli.Pdf.SetFillColor(color.R, color.G, color.B)
		if i%2 == 0 {
			//cli.Pdf.CellFormat(150, 5, str, "1", 0, "", true, 0, "")
			cli.Pdf.MultiCell(cli.CardWith, 5, str, "1", "", true)
		} else {
			cli.Pdf.MultiCell(cli.CardWith, 5, str+"\n ", "1", "", false)
		}
	}
}

func (cli *Client) addCardsRow8(color Color, jh float64, assignee []string) {
	tr := cli.UnicodeTranslatorFromDescriptor("")
	WrapedAssignee := WrapText(strings.Join(assignee, ", "), 25)
	diffLen := strings.Count(WrapedAssignee, "\n")
	JH := "Charge Estimée (J/H) :" + strings.Repeat("\n ", diffLen)
	JHValue := strconv.FormatFloat(jh, 'f', -1, 64) + strings.Repeat("\n ", diffLen)
	Assignee := "Assignés (J/H) :" + strings.Repeat("\n ", diffLen)

	cli.Pdf.SetX((cli.Width - cli.CardWith) / 2)
	for i, str := range []string{tr(JH), JHValue, tr(Assignee), tr(WrapedAssignee)} {
		x := cli.Pdf.GetX()        //remove partially a bug
		y := cli.Pdf.GetY()        //remove partially a bug
		pageNo := cli.Pdf.PageNo() //remove partially a bug

		if i%2 == 0 {
			cli.Pdf.SetFillColor(color.R, color.G, color.B)
		} else {
			cli.Pdf.SetFillColor(255, 255, 255)
		}
		w := tools.Ternary(i < 1 || i > 2, cli.CardWith/3, cli.CardWith/6)
		cli.Pdf.SetPage(pageNo) //remove partially a bug
		cli.Pdf.MultiCell(w, 5, str, "1", "", true)
		if i < 3 {
			cli.Pdf.SetXY(x+w, y) //remove partially a bug
		}
	}
	cli.Pdf.Ln(-1)
}
