package db

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
