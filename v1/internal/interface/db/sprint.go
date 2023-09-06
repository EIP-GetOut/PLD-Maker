package db

type SprintFields struct {
	Title  string   `json:"Name"`
	Cards  []string `json:"Card"`
	Status string   `json:"Status"`
	Number int      `json:"Number"`
}

type Sprint struct {
	Id          string       `json:"id"`
	CreatedTime string       `json:"createdTime"`
	Fields      SprintFields `json:"fields"`
}

type Sprints struct {
	Sprints []Sprint `json:"records"`
	Offset  string   `json:"offset"`
}
