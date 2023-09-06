package pdf

// Text
type TextParams struct {
	Bold       bool
	Italic     bool
	Underline  bool
	Overline   bool
	Size       *int
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
