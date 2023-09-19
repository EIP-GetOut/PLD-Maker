package airtablewrapper

import (
	"net/url"
	"pld-maker/v1/internal/interface/db"
)

type File struct {
	Id     string `json:"id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Url    string `json:"url"`
	Type   string `json:"type"`
}

type SchemaFields struct {
	Files []File `json:"File"`
	Title string `json:"Name"`
}

type Schema struct {
	Id          string       `json:"id"`
	CreatedTime string       `json:"createdTime"`
	Fields      SchemaFields `json:"fields"`
}

type Schemas struct {
	Schemas []Schema `json:"records"`
	Offset  string   `json:"offset"`
}

// Interface & Generic Function
func (s Schemas) GetOffset() string {
	return s.Offset
}

func (s Schemas) SetOffset(str string) {
	s.Offset = str
}

func SchemasAppender(tmp Schemas) []db.Schema {
	var result []db.Schema

	for _, v := range tmp.Schemas {
		result = append(result, SchemaTranslater(v))
	}
	return result
}

func SchemaTranslater(tmp Schema) db.Schema {
	return db.Schema{
		Id:     tmp.Id,
		Width:  tmp.Fields.Files[0].Width,
		Height: tmp.Fields.Files[0].Height,
		Url:    tmp.Fields.Files[0].Url,
		Type:   tmp.Fields.Files[0].Type,
		Title:  tmp.Fields.Title,
	}
}

// Schema
func (cli *Client) ListSchemas(params url.Values) ([]db.Schema, error) {
	return ListItems[Schemas, db.Schema](cli, "Schema", params, SchemasAppender)
}

func (cli *Client) GetSchema(id string) (db.Schema, error) {
	return GetItem[Schema, db.Schema](cli, "Schema", id, SchemaTranslater)
}
