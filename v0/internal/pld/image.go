package pld

import (
	"os"
	"pld-maker/v0/internal/tools"
	"strings"

	"github.com/go-pdf/fpdf"
)

func (cli *Client) AddImage(filepath string, y float64, w float64, h float64) {
	//Image
	var (
		opt fpdf.ImageOptions
		fl  *os.File
	)

	fl = tools.Must(os.Open(filepath))
	idx := strings.LastIndex(filepath, ".")
	if idx == -1 {
		panic("image?? extension not specified")
	}
	extension := filepath[(idx + 1):]
	opt.ImageType = extension
	opt.AllowNegativePosition = true
	_ = cli.Pdf.RegisterImageOptionsReader("logo", opt, fl)
	fl.Close()
	cli.Pdf.ImageOptions("logo", (cli.Width-w)/2, y, w, h, false, opt, 0, "")
	cli.Pdf.SetY(y + h)
}
