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

type Text struct {
	Data   string
	Params *TextParams
}
