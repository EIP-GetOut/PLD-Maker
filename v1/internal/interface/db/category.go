package db

type Category struct {
	Id      string   `json:"id"`
	Name    string   `json:"Name"`
	Cards   []string `json:"Card"`
	Sprints []string `json:"Sprint"`
	Sectors []string `json:"Sector"`
}
