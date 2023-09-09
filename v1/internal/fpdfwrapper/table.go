package fpdfwrapper

import (
	"fmt"
	"pld-maker/v1/internal/interface/pdf"
)

func (cli *Client) setTableParams(params *pdf.TableParams) {
}

func (cli *Client) Table(table pdf.Table) {
	cli.setTableParams(table.Params)
	for _, rows := range table.Rows {
		for _, cell := range rows.Cells {
			fmt.Print("\t" + cell.Str)
		}
		fmt.Println("")
	}
}
