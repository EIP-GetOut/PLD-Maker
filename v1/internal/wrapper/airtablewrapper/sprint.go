package airtablewrapper

import (
	"net/url"
	"pld-maker/v1/internal/interface/db"
)

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

// Interface & Generic Function
func (s Sprints) GetOffset() string {
	return s.Offset
}

func (s Sprints) SetOffset(str string) {
	s.Offset = str
}

func SprintsAppender(tmp Sprints) []db.Sprint {
	var result []db.Sprint

	for _, v := range tmp.Sprints {
		result = append(result, SprintTranslater(v))
	}
	return result
}

func SprintTranslater(tmp Sprint) db.Sprint {
	return db.Sprint{
		Id:     tmp.Id,
		Title:  tmp.Fields.Title,
		Cards:  tmp.Fields.Cards,
		Status: tmp.Fields.Status,
		Number: tmp.Fields.Number,
	}
}

// Sprints
func (cli *Client) ListSprints(params url.Values) ([]db.Sprint, error) {
	return ListItems[Sprints, db.Sprint](cli, "Sprint", params, SprintsAppender)
}
func (cli *Client) GetSprint(id string) (db.Sprint, error) {
	return GetItem[Sprint, db.Sprint](cli, "Sprint", id, SprintTranslater)
}
