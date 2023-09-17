package fpdfwrapper

import (
	"pld-maker/v1/internal/interface/pdf"
)

func (cli *Client) setTextDefault(size float64) {
	cli.pdf.SetDrawColor(0, 0, 0)
	cli.pdf.SetFillColor(255, 255, 255)
	cli.pdf.SetTextColor(0, 0, 0)
	cli.pdf.SetFont("Arial", "", size)
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
			{f: func() { cli.pdf.SetFontStyle(styleStr) }, b: true},
			{f: func() { cli.pdf.SetFillColor(params.Background.R, params.Background.G, params.Background.B) }, b: params.Background != nil},
			{f: func() { cli.pdf.SetTextColor(params.TextColor.R, params.TextColor.G, params.TextColor.B) }, b: params.TextColor != nil},
			{f: func() { cli.pdf.SetFontSize(params.Size) }, b: params.Size != 0},
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

func (cli *Client) Title(text pdf.Text) {
	var alignStr string = ""

	if text.Params != nil {
		alignStr = alignToStr(text.Params.Align)
	}
	cli.setTextDefault(30)
	cli.TextParams(text.Params)
	cli.pdf.MultiCell(cli.HeaderWidth, 10, cli.translator(text.Data), "", alignStr, true)
}

func (cli *Client) SubTitle(text pdf.Text) {
	var alignStr string = ""

	cli.setTextDefault(24)
	cli.TextParams(text.Params)
	if text.Params != nil {
		alignStr = alignToStr(text.Params.Align)
	}
	cli.pdf.MultiCell(cli.HeaderWidth, 9, cli.translator(text.Data), "", alignStr, true)
}

func (cli *Client) Heading1(text pdf.Text) {
	var alignStr string = ""

	cli.setTextDefault(20)
	cli.TextParams(text.Params)
	if text.Params != nil {
		alignStr = alignToStr(text.Params.Align)
	}
	cli.pdf.MultiCell(cli.HeaderWidth, 8, cli.translator(text.Data), "", alignStr, true)
}

func (cli *Client) Heading2(text pdf.Text) {
	var alignStr string = ""

	cli.setTextDefault(16)
	cli.TextParams(text.Params)
	if text.Params != nil {
		alignStr = alignToStr(text.Params.Align)
	}
	cli.pdf.MultiCell(cli.HeaderWidth, 7, cli.translator(text.Data), "", alignStr, true)
}

func (cli *Client) Text(text pdf.Text) {
	var alignStr string = ""

	cli.setTextDefault(12)
	cli.TextParams(text.Params)
	if text.Params != nil {
		alignStr = alignToStr(text.Params.Align)
	}
	cli.pdf.MultiCell(cli.HeaderWidth, 5, cli.translator(text.Data), "", alignStr, true)
}
