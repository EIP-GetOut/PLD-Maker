package db

type Category struct {
	Name  string   `json:"Name"`
	Cards []string `json:"Card"`
}
