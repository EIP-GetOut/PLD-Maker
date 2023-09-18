package fpdfwrapper

import (
	"github.com/go-pdf/fpdf"
)

func (cli *Client) NewFile(filename string) {
	cli.fileName = filename
	cli.pdf = fpdf.New("P", "mm", "A4", "")

	cli.Width, cli.Height = cli.pdf.GetPageSize()
	if cli.pdf == nil {
		panic("fpdf.New failed")
	}

	cli.HeaderWidth = 190
	cli.FooterWidth = 190
	cli.ImageWidth = 200

	cli.TableWidth = 190
	cli.TextWidth = 180

	cli.translator = cli.pdf.UnicodeTranslatorFromDescriptor("")
}

func (cli *Client) CloseFile() {
	err := cli.pdf.OutputFileAndClose(cli.fileName + ".pdf")
	if err != nil {
		panic(err)
	}
}

func (cli *Client) NewPage() {
	cli.pdf.AddPage()
	cli.pdf.SetXY(10, 20)
}

func (cli *Client) NewLine() {
	cli.pdf.Ln(-1)
}
