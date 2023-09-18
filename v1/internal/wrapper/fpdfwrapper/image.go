package fpdfwrapper

import (
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
			{f: func() { x = (cli.Width - w) / 2 }, b: params.X == 0},
			//	X,Y set float64
			{f: func() { x = params.X }, b: params.X != 0 && !params.XPercent},
			{f: func() { y = params.Y }, b: params.Y != 0},
			//	XPercent, YPercent bool
			{f: func() { x = (cli.Width - w) * params.X }, b: params.X != 0 && params.XPercent},
		} {
			if v.b {
				v.f()
			}
		}
	}
	return x, y
}

func (cli *Client) Image(image pdf.Image) {
	var (
		opt    fpdf.ImageOptions
		fl     *os.File = tools.Must(os.Open(image.Filepath))
		idxDot int      = strings.LastIndex(image.Filepath, ".")
		//idxSlash int = strings.LastIndex(filepath, "/")
	)

	if idxDot == -1 {
		panic("image?? extension not specified")
	}
	//name := filepath[(idxSlash + 1):(idxDot)]
	extension := image.Filepath[(idxDot + 1):]
	opt.ImageType = extension
	opt.AllowNegativePosition = true

	cli.pdf.RegisterImageOptionsReader(image.Filepath, opt, fl)
	fl.Close()
	cli.pdf.SetX((cli.Width - image.Width) / 2)
	x, y := cli.pdf.GetXY()
	x, y = cli.setImageParams(x, y, image.Width, image.Height, image.Params)
	cli.pdf.ImageOptions(image.Filepath, x, y, image.Width, image.Height, false, opt, 0, "")
	cli.pdf.SetY(y + image.Height)
}
