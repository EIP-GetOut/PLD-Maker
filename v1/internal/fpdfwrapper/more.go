package fpdfwrapper

import (
	"fmt"
	"os"
	"pld-maker/v1/internal/interface/pdf"
	"pld-maker/v1/internal/tools"
	"strings"

	"github.com/go-pdf/fpdf"
)

func (cli *Client) Image(y, w, h float64, filepath string) {
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

func (cli *Client) Image2(filepath string, w, h float64, params *pdf.ImageParams) {
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
	cli.Pdf.ImageOptions("logo", (cli.Width-w)/2, 10, w, h, false, opt, 0, "")
	//cli.Pdf.SetY(y + h)
}

func (cli *Client) Table(data [][]string, tableParams *pdf.TableParams) {

}
