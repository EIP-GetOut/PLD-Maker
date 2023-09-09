package fpdfwrapper

import (
	"pld-maker/v1/internal/interface/pdf"
)

func (cli *Client) setTextDefault(size float64) {
	cli.Pdf.SetDrawColor(0, 0, 0)
	cli.Pdf.SetFillColor(255, 255, 255)
	cli.Pdf.SetTextColor(0, 0, 0)
	cli.Pdf.SetFont("Arial", "", size)
}

func (cli *Client) TextParams(params *pdf.TextParams) {
	var styleStr string = ""

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
			{f: func() { cli.Pdf.SetFontSize(params.Size) }, b: params.Size != 0},
		} {
			if v.b {
				v.f()
			}
		}
	}
}

func alignToStr(pos pdf.HorizontalPosition) string {
	switch pos {
	case 1:
		return "L"
	case 2:
		return "C"
	case 3:
		return "R"
	default:
		return ""
	}
}

func (cli *Client) Title(str string, params *pdf.TextParams) {
	var alignStr string = ""

	cli.setTextDefault(30)
	cli.TextParams(params)
	if params != nil {
		alignStr = alignToStr(params.Align)
	}
	cli.Pdf.MultiCell(cli.HeaderWidth, 5, cli.translator(str), "", alignStr, true)
	cli.Pdf.Ln(-1)

}

func (cli *Client) SubTitle(str string, params *pdf.TextParams) {
	var alignStr string = ""

	cli.setTextDefault(24)
	cli.TextParams(params)
	if params != nil {
		alignStr = alignToStr(params.Align)
	}
	cli.Pdf.MultiCell(cli.HeaderWidth, 5, cli.translator(str), "", alignStr, true)
	cli.Pdf.Ln(-1)
}

func (cli *Client) Heading1(str string, params *pdf.TextParams) {
	var alignStr string = ""

	cli.setTextDefault(20)
	cli.TextParams(params)
	if params != nil {
		alignStr = alignToStr(params.Align)
	}
	cli.Pdf.MultiCell(cli.HeaderWidth, 5, cli.translator(str), "", alignStr, true)
	cli.Pdf.Ln(-1)
}

func (cli *Client) Heading2(str string, params *pdf.TextParams) {
	var alignStr string = ""

	cli.setTextDefault(16)
	cli.TextParams(params)
	if params != nil {
		alignStr = alignToStr(params.Align)
	}
	cli.Pdf.MultiCell(cli.HeaderWidth, 5, cli.translator(str), "", alignStr, true)
	cli.Pdf.Ln(-1)
}

func (cli *Client) Text(str string, params *pdf.TextParams) {
	var alignStr string = ""

	cli.setTextDefault(12)
	cli.TextParams(params)
	if params != nil {
		alignStr = alignToStr(params.Align)
	}
	cli.Pdf.MultiCell(cli.HeaderWidth, 5, cli.translator(str), "", alignStr, true)
	cli.Pdf.Ln(-1)
}
