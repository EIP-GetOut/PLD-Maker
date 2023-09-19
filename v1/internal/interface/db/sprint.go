package db

type Sprint struct {
	Id     string   `json:"id"`
	Title  string   `json:"Name"`
	Cards  []string `json:"Card"`
	Status string   `json:"Status"`
	Number int      `json:"Number"`
}
