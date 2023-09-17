package pdf

// Same as Text
type CellParams struct {
	Bold       bool
	Italic     bool
	Underline  bool
	Overline   bool
	Size       float64
	TextColor  *Color
	Background *Color
	Align      HorizontalPosition
}

type RowParams struct {
	*CellParams

	RowHeight float64 //rowHeight
}

type TableParams struct {
	*RowParams

	Width     float64 //width
	DrawColor *Color
}

type Cell struct {
	Str     string  //external
	Percent float64 //external
	ZtoO    bool    //external
	Params  *CellParams
}

type Row struct {
	Cells  []Cell
	Params *RowParams
}

type Table struct {
	Rows   []Row
	Params *TableParams
}
