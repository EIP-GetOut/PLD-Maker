package pdf

type Cell struct {
	data    string
	percent float64
}

type Row struct {
	cells []Cell
}

type Table struct {
	rows []Row
	data string
}
