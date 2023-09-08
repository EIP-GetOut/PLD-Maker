package pdf

// Text
type TextParams struct {
	Bold       bool
	Italic     bool
	Underline  bool
	Overline   bool
	Size       *float64
	TextColor  *Color
	Background *Color
}

// Table
type TableParams struct {
	Top  bool
	Left bool
	Head *TextParams
	Body *TextParams
}

// if you setup footer pageNo you won't display right string on your pdf
type FooterParams struct {
	PageNo      bool
	FirstPageNo bool
}
