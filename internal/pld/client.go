package pld

import (
	"github.com/go-pdf/fpdf"
)

type Color struct {
	R int
	G int
	B int
}

type Client struct {
	Pdf           *fpdf.Fpdf
	Width         float64
	Height        float64
	HeaderWidth   float64
	CardWith      float64
	PercentColors []Color
}

func NewClient() (*Client, error) {
	var cli Client

	//if err := json.Unmarshal(conf, &cli); err != nil { return nil, err }
	cli.Pdf = fpdf.New("P", "mm", "A4", "")
	cli.Width, cli.Height = cli.Pdf.GetPageSize()
	cli.HeaderWidth = 190
	cli.CardWith = 180

	//cli.Left = (cli.Width - 4*40) / 2
	cli.PercentColors = []Color{{234, 153, 153}, {255, 229, 153}, {182, 215, 168}}
	//	a4c2f4	blue
	return &cli, nil
}

// OutputFileAndClose creates or truncates the file specified by fileStr and writes the PDF document to it.
// This method will close f and the newly written file, even if an error is detected and no document is produced.
//
// Most examples demonstrate the use of this method.
func (cli *Client) OutputFileAndClose(fileStr string) error {
	return cli.Pdf.OutputFileAndClose(fileStr)
}

// UnicodeTranslatorFromDescriptor returns a function that can be used to translate, where possible,
// utf-8 strings to a form that is compatible with the specified code page. See UnicodeTranslator for more details.
//
// cpStr identifies a code page. A descriptor file in the font directory, set with the fontDirStr argument
// in the call to New(), should have this name plus the extension ".map".
// If cpStr is empty, it will be replaced with "cp1252", the gofpdf code page default.
//
// If an error occurs reading the descriptor, the returned function is valid but does not perform any rune translation.
//
// The CellFormat_codepage example demonstrates this method.
func (cli *Client) UnicodeTranslatorFromDescriptor(cpStr string) (rep func(string) string) {
	return cli.Pdf.UnicodeTranslatorFromDescriptor(cpStr)
}
