package fpdfwrapper

import (
	"fmt"
	"pld-maker/v1/internal/interface/pdf"
)

func (cli *Client) setTextParams(params *pdf.TextParams) {

}

func (cli *Client) Title(str string, params *pdf.TextParams) {
	fmt.Println("Title")
}

func (cli *Client) SubTitle(str string, params *pdf.TextParams) {
	fmt.Println("SubTitle")
}

func (cli *Client) Heading1(str string, params *pdf.TextParams) {
	fmt.Println("Heading1")
}

func (cli *Client) Heading2(str string, params *pdf.TextParams) {
	fmt.Println("Heading2")
}

func (cli *Client) Text(str string, params *pdf.TextParams) {
	var styleStr string = ""

	cli.Pdf.SetDrawColor(0, 0, 0)
	cli.Pdf.SetFillColor(255, 255, 255)
	cli.Pdf.SetTextColor(0, 0, 0)

	//cli.setTextParams(params)

	if params != nil {
		for _, v := range []struct {
			f func()
			b bool
		}{
			{f: func() { styleStr += "B" }, b: params.Bold},
			{f: func() { styleStr += "I" }, b: params.Italic},
			{f: func() { styleStr += "U" }, b: params.Underline},
			{f: func() { styleStr += "S" }, b: params.Overline},
			{f: func() { cli.Pdf.SetFillColor(params.Background.R, params.Background.G, params.Background.B) }, b: params.Background != nil},
			{f: func() { cli.Pdf.SetTextColor(params.TextColor.R, params.TextColor.G, params.TextColor.B) }, b: params.TextColor != nil},
			{f: func() { cli.Pdf.SetFontSize(*params.Size) }, b: params.Size != nil},
		} {
			if v.b {
				v.f()
			}
		}
		cli.Pdf.SetFont("Arial", styleStr, 12)

		//		for _, value := range []struct {
		//			f func(r, g, b int)
		//			c *pdf.Color
		//		}{
		//			{f: cli.Pdf.SetFillColor, c: params.Background},
		//			{f: cli.Pdf.SetTextColor, c: params.TextColor},
		//		} {
		//			if color := value.c; color != nil {
		//				value.f(color.R, color.G, color.B)
		//			}
		//		}
		//		//Size
		//		if params.Size != nil {
		//			cli.Pdf.SetFont("Arial", styleStr, *params.Size)
		//		}
	}
	cli.Pdf.MultiCell(cli.HeaderWidth, 7, str, "", "", true)
}
