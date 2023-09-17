package fpdfwrapper

import (
	"pld-maker/internal/tools"
	"pld-maker/v1/internal/interface/pdf"
)

var bToI = map[bool]int{
	false: 0,
	true:  1,
}

func (cli *Client) setCellParams(params *pdf.CellParams) {
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
			{f: func() { cli.pdf.SetFillColor(params.Background.R, params.Background.G, params.Background.B) }, b: params.Background != nil},
			{f: func() { cli.pdf.SetTextColor(params.TextColor.R, params.TextColor.G, params.TextColor.B) }, b: params.TextColor != nil},
			{f: func() { cli.pdf.SetFontSize(params.Size) }, b: params.Size != 0},
		} {
			if v.b {
				v.f()
			}
		}
	}
	cli.pdf.SetFontStyle(styleStr)
}

func (cli *Client) setRowParams(params *pdf.RowParams) {
	//var styleStr string = ""

	if params != nil {
		for _, v := range []struct {
			f func()
			b bool
		}{
			{f: func() { cli.setCellParams(params.CellParams) }, b: params.CellParams != nil},
		} {
			if v.b {
				v.f()
			}
		}
	}
}

func (cli *Client) setTableParams(params *pdf.TableParams) {
	if params != nil {
		for _, v := range []struct {
			f func()
			b bool
		}{
			{f: func() { cli.pdf.SetDrawColor(params.DrawColor.R, params.DrawColor.G, params.DrawColor.B) }, b: params != nil && params.DrawColor != nil},
			{f: func() { cli.setRowParams(params.RowParams) }, b: params.RowParams != nil},
		} {
			if v.b {
				v.f()
			}
		}
	}
}

func (cli *Client) cellAlign(tableParams *pdf.TableParams, rowParams *pdf.RowParams, cellParams *pdf.CellParams) {
	if tableParams != nil {
		cli.setTableParams(tableParams)
		if rowParams != nil {
			cli.setRowParams(rowParams)
			if cellParams != nil {
				cli.setCellParams(cellParams)
			}
		}
	}
}

func (cli *Client) Table(table pdf.Table) {
	var alignStr string = ""
	var width float64 = tools.Ternary((table.Params != nil && table.Params.Width != 0), table.Params.Width, cli.TableWidth)
	var height float64 = tools.Ternary((table.Params != nil && table.Params.RowHeight != 0), table.Params.RowHeight, 5)

	for _, row := range table.Rows {
		cli.pdf.SetX((cli.Width - width) / 2)
		for _, cell := range row.Cells {
			cli.setTextDefault(12)
			percentSize := tools.Ternary(cell.ZtoO, cell.Percent, cell.Percent/100)
			// alignTextInCell
			if cell.Params != nil {
				cli.setCellParams(cell.Params)
				alignStr = alignToStr(cell.Params.Align)
			} else if row.Params != nil {
				cli.setRowParams(row.Params)
				if row.Params.CellParams != nil {
					alignStr = alignToStr(row.Params.Align)
				}
			} else if table.Params != nil {
				cli.setTableParams(table.Params)
				if table.Params.RowParams.CellParams != nil {
					alignStr = alignToStr(table.Params.Align)
				}
			}
			x, y := cli.pdf.GetXY()
			cli.pdf.MultiCell((width)*percentSize, height, cell.Str, "1", alignStr, true)
			cli.pdf.SetXY(x+(width*percentSize), y)
		}
		cli.pdf.Ln(-1)
	}
}
