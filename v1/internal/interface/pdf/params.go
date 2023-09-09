package pdf

type HorizontalPosition int64

const (
	Default HorizontalPosition = iota
	Left
	Center
	Right
)

// Text
type TextParams struct {
	Bold       bool
	Italic     bool
	Underline  bool
	Overline   bool
	Size       float64
	TextColor  *Color
	Background *Color
	Align      HorizontalPosition
}

// Image
type ImageParams struct {
	X              float64
	Y              float64
	XPercent       bool
	YPercent       bool
	TopLeftGravity bool
}

// if you setup footer pageNo you won't display right string on your pdf
type FooterParams struct {
	PageNo      bool
	FirstPageNo bool
}
