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
	PercentColors []Color
	Tr            func(string) string
}

func NewClient() (*Client, error) {
	var cli Client

	//if err := json.Unmarshal(conf, &cli); err != nil { return nil, err }
	cli.Pdf = fpdf.New("P", "mm", "A4", "")
	cli.Width = 210.0
	cli.Tr = cli.Pdf.UnicodeTranslatorFromDescriptor("")

	//cli.Left = (cli.Width - 4*40) / 2
	cli.PercentColors = []Color{{234, 153, 153}, {255, 229, 153}, {182, 215, 168}}
	//	a4c2f4	blue
	return &cli, nil
}
