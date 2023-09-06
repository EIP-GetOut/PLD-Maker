package db

type CategoryFields struct {
	Name  string   `json:"Name"`
	Cards []string `json:"Card"`
}

type Category struct {
	Id          string         `json:"id"`
	CreatedTime string         `json:"createdTime"`
	Fields      CategoryFields `json:"fields"`
}

type Categories struct {
	Categories []Category `json:"records"`
	Offset     string     `json:"offset"`
}
