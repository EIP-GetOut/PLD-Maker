package pdf

type TableParams struct {
	Background *Color
	TextColor  *Color
	TableColor *Color
}

type Cell struct {
	Str     string
	Percent float64
	ZtoO    bool
	Params  *TableParams
}

type Row struct {
	Cells  []Cell
	Params *TableParams
}

type Table struct {
	Rows   []Row
	Params *TableParams
}
