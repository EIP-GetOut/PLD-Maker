package fpdfwrapper

import (
	"pld-maker/v1/internal/interface/pdf"
)

var bToI = map[bool]int{
	false: 0,
	true:  1,
}

func (cli *Client) setTableParams(params *pdf.TableParams) {
}

func (cli *Client) Table(table pdf.Table) {
	cli.setTableParams(table.Params)
	for _, rows := range table.Rows {
		cli.Pdf.SetX((cli.Width - cli.TableWidth) / 2)
		for _, cell := range rows.Cells {
			percentSize := cell.Percent / (100 - (99 * float64(bToI[cell.ZtoO])))
			x, y := cli.Pdf.GetXY()
			cli.Pdf.MultiCell((cli.TableWidth)*percentSize, 5, cell.Str, "1", "", true)
			cli.Pdf.SetXY(x+(cli.TableWidth)*percentSize, y)
		}
		cli.Pdf.Ln(-1)
	}
}
