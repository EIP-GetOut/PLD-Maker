package airtablewrapper

import (
	"net/url"
	"pld-maker/v1/internal/interface/db"
)

type VersionFields struct {
	Date     string
	Version  string
	Author   string
	Sections string
	Comments string
}

type Version struct {
	Id          string        `json:"id"`
	CreatedTime string        `json:"createdTime"`
	Fields      VersionFields `json:"fields"`
}

type Versions struct {
	Versions []Version `json:"records"`
	Offset   string    `json:"offset"`
}

func (v Versions) GetOffset() string {
	return v.Offset
}

func (v Versions) SetOffset(str string) {
	v.Offset = str
}

func VersionsAppender(tmp Versions) []db.Version {
	var result []db.Version

	for _, v := range tmp.Versions {
		result = append(result, VersionTranslater(v))
	}
	return result
}

func VersionTranslater(tmp Version) db.Version {
	return db.Version{
		Id:       tmp.Id,
		Date:     tmp.Fields.Date,
		Version:  tmp.Fields.Version,
		Author:   tmp.Fields.Author,
		Sections: tmp.Fields.Sections,
		Comments: tmp.Fields.Comments,
	}
}

// Versions
func (cli *Client) ListVersions(params url.Values) ([]db.Version, error) {
	return ListItems[Versions, db.Version](cli, "Version", params, VersionsAppender)
}

func (cli *Client) GetVersion(id string) (db.Version, error) {
	return GetItem[Version, db.Version](cli, "Version", id, VersionTranslater)
}
