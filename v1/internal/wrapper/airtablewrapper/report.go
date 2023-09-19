package airtablewrapper

import (
	"net/url"
	"pld-maker/v1/internal/interface/db"
)

type ReportFields struct {
	Name  string `json:"Name"`
	Notes string `json:"Notes"`
}

type Report struct {
	Id          string       `json:"id"`
	CreatedTime string       `json:"createdTime"`
	Fields      ReportFields `json:"fields"`
}

type Reports struct {
	Reports []Report `json:"records"`
	Offset  string   `json:"offset"`
}

// Interface & Generic Function
func (s Reports) GetOffset() string {
	return s.Offset
}

func (s Reports) SetOffset(str string) {
	s.Offset = str
}

func ReportsAppender(tmp Reports) []db.Report {
	var result []db.Report

	for _, v := range tmp.Reports {
		result = append(result, ReportTranslater(v))
	}
	return result
}

func ReportTranslater(tmp Report) db.Report {
	return db.Report{
		Id:    tmp.Id,
		Name:  tmp.Fields.Name,
		Notes: tmp.Fields.Notes,
	}
}

// Reports
func (cli *Client) ListReports(params url.Values) ([]db.Report, error) {
	return ListItems[Reports, db.Report](cli, "Report", params, ReportsAppender)
}

func (cli *Client) GetReport(id string) (db.Report, error) {
	return GetItem[Report, db.Report](cli, "Report", id, ReportTranslater)
}
