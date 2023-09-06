package fpdfwrapper

import (
	"github.com/go-pdf/fpdf"
)

type Client struct {
	//Global
	Pdf      *fpdf.Fpdf
	fileName string
	Width    float64
	Height   float64
	//Specific
	HeaderWidth float64
	FooterWidth float64
	TableWidth  float64
	TextWidth   float64
	//Miscellaneous
	translator func(string) string
}

func NewClient() (*Client, error) {
	return &Client{}, nil
}
