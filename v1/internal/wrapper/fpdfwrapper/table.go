package fpdfwrapper

import (
	"pld-maker/v1/internal/interface/pdf"
	"pld-maker/v1/internal/tools"
	"strings"
)

// Params
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

// other
func cutStr(str string, max int) string {
	var result string = ""
	var lenght int = len(str)
	// j define the further position, k is used to find previous space when no '\n' is found in max range.
	var i, j, k int
	var skip bool

	// i define start of substring
	for i = 0; i < lenght; i++ {
		skip = false
		for j = i; skip == false && j < lenght && j < i+max; j++ {
			if str[j] == '\n' {
				result += str[i : j+1]
				i = j
				skip = true
			}
		}
		for k = j; skip == false && k < lenght && k > 0 && k > i; k-- {
			if str[k] == ' ' {
				result += str[i:k] + "\n"
				i = k
				skip = true
			}
		}
		if skip == false {
			if j == lenght {
				result += str[i:lenght]
				break
			} else {
				result += str[i:j+1] + "\n"
				i = j
			}
		}
	}
	return result
}

// Cli
func (cli *Client) Table(table pdf.Table) {
	var (
		alignStr string
		width    float64
		height   float64
	)
	if table.Params != nil && table.Params.Width != 0 {
		width = table.Params.Width
	} else {
		width = cli.TableWidth
	}
	if table.Params != nil && table.Params.RowParams != nil && table.Params.RowHeight != 0 {
		height = table.Params.RowHeight
	} else {
		height = 5
	}

	for _, row := range table.Rows {
		cli.pdf.SetX((cli.Width - width) / 2)

		//Edit Cell of Each Rows
		var (
			maxX int
			maxY int
		)
		for i := range row.Cells {
			var cell *pdf.Cell = &row.Cells[i]
			//180 == 80
			maxX = int((width / 2) * tools.Ternary(cell.ZtoO, cell.Percent, cell.Percent/100))
			cell.Str = cutStr(cell.Str, maxX)
			if maxY < strings.Count(cell.Str, "\n") {
				maxY = strings.Count(cell.Str, "\n")
			}
		}

		for i := range row.Cells {
			var cell *pdf.Cell = &row.Cells[i]
			y := strings.Count(cell.Str, "\n")
			cell.Str += strings.Repeat("\n ", maxY-y)
		}

		for i, cell := range row.Cells {

			//Define default value
			cli.setTextDefault(12)
			percentSize := tools.Ternary(cell.ZtoO, cell.Percent, cell.Percent/100)

			//Define Table Preferences
			if table.Params != nil {
				cli.setTableParams(table.Params)
				if table.Params.RowParams != nil && table.Params.RowParams.CellParams != nil {
					alignStr = alignToStr(table.Params.Align)
				}
			}
			//Define Row Preferences
			if row.Params != nil {
				cli.setRowParams(row.Params)
				if row.Params.CellParams != nil {
					alignStr = alignToStr(row.Params.CellParams.Align)
				}
			}
			//Define Cell Preferences
			if cell.Params != nil {
				cli.setCellParams(cell.Params)
				alignStr = alignToStr(cell.Params.Align)
			}

			x, y := cli.pdf.GetXY()
			cli.pdf.MultiCell((width)*percentSize, height, cli.translator(cell.Str), "1", alignStr, true)
			if i < len(row.Cells)-1 {
				cli.pdf.SetXY(x+(width*percentSize), y)
			}
		}
	}
}
