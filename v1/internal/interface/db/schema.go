package db

type Schema struct {
	Id     string `json:"id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Url    string `json:"url"`
	Title  string `json:"Name"`
	Type   string `json:"type"`
}
