package fpdfwrapper

import (
	"github.com/go-pdf/fpdf"
)

func (cli *Client) NewFile(filename string) {
	cli.fileName = filename
	cli.Pdf = fpdf.New("P", "mm", "A4", "")

	cli.Width, cli.Height = cli.Pdf.GetPageSize()
	if cli.Pdf == nil {
		panic("fpdf.New failed")
	}

	cli.HeaderWidth = 190
	cli.FooterWidth = 190

	cli.TableWidth = 180
	cli.TextWidth = 180

	cli.translator = cli.Pdf.UnicodeTranslatorFromDescriptor("")
}

func (cli *Client) CloseFile() {
	err := cli.Pdf.OutputFileAndClose(cli.fileName + ".pdf")
	if err != nil {
		panic(err)
	}
}

func (cli *Client) NewPage() {
	cli.Pdf.AddPage()
	cli.Pdf.SetXY(10, 20)
}
