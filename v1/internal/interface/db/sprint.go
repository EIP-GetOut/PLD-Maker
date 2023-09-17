package db

type Sprint struct {
	Title  string   `json:"Name"`
	Cards  []string `json:"Card"`
	Status string   `json:"Status"`
	Number int      `json:"Number"`
}
