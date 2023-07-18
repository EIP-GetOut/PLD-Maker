package epitech

import (
	"pld-maker/internal/pld"
	"strconv"
)

func FirstPage(cli *pld.Client, sprintNumber int) {
	cli.AddPage()
	cli.AddImage("./conf/epitech.png", 50, 160, 50)
	cli.AddTitle1("")
	cli.AddTitle1("EPITECH INNOVATIVE PROJECT")
	cli.AddTitle1("PROJECT LOG DOCUMENT")
	cli.AddTitle1("SPRINT NUMERO " + strconv.Itoa(sprintNumber))
	cli.Pdf.Ln(-1)
	cli.AddTitle1B("PROMO 2025")
}
