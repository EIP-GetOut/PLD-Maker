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
}

func NewClient() (*Client, error) {
	var cli Client

	//if err := json.Unmarshal(conf, &cli); err != nil { return nil, err }
	cli.Pdf = fpdf.New("P", "mm", "A4", "")
	cli.Width = 210.0
	//cli.Left = (cli.Width - 4*40) / 2
	cli.PercentColors = []Color{{255, 76, 66}, {255, 176, 97}, {122, 255, 160}}
	return &cli, nil
}
