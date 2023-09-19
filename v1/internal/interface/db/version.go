package db

type Version struct {
	Id       string `json:"id"`
	Date     string
	Version  string
	Author   string
	Sections string
	Comments string
}
