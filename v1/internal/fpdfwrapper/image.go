package fpdfwrapper

import (
	"fmt"
	"os"
	"pld-maker/v1/internal/interface/pdf"
	"pld-maker/v1/internal/tools"
	"strings"

	"github.com/go-pdf/fpdf"
)

func (cli *Client) setImageParams(x, y, w, h float64, params *pdf.ImageParams) (float64, float64) {
	if params != nil {
		for _, v := range []struct {
			f func()
			b bool
		}{
			{f: func() { w = 0; h = 0 }, b: params.TopLeftGravity},
			//
			{f: func() { x = (cli.ImageWidth - w) * params.X }, b: params.X != 0 && params.XPercent},
			{f: func() { y = (cli.Height - h) * params.Y }, b: params.Y != 0 && params.YPercent},
			{f: func() { x = params.X - (w / 2) }, b: params.X != 0 && !params.XPercent},
			{f: func() { y = params.Y - (h / 2) }, b: params.Y != 0 && !params.YPercent},
		} {
			if v.b {
				v.f()
			}
		}
	}
	return x, y
}

func (cli *Client) Image(filepath string, w, h float64, params *pdf.ImageParams) {
	fmt.Println(filepath)
	var (
		opt    fpdf.ImageOptions
		fl     *os.File = tools.Must(os.Open(filepath))
		idxDot int      = strings.LastIndex(filepath, ".")
		//idxSlash int = strings.LastIndex(filepath, "/")
	)

	if idxDot == -1 {
		panic("image?? extension not specified")
	}
	//name := filepath[(idxSlash + 1):(idxDot)]
	extension := filepath[(idxDot + 1):]
	opt.ImageType = extension
	opt.AllowNegativePosition = true

	cli.Pdf.RegisterImageOptionsReader("logo", opt, fl)
	fl.Close()
	x, y := cli.Pdf.GetXY()
	x, y = cli.setImageParams(x, y, w, h, params)

	cli.Pdf.ImageOptions("logo", x, y, w, h, false, opt, 0, "")
}
