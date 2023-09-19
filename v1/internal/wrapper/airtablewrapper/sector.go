package airtablewrapper

import (
	"net/url"
	"pld-maker/v1/internal/interface/db"
)

type SectorFields struct {
	Name string `json:"Name"`
}

type Sector struct {
	Id          string       `json:"id"`
	CreatedTime string       `json:"createdTime"`
	Fields      SectorFields `json:"fields"`
}

type Sectors struct {
	Sectors []Sector `json:"records"`
	Offset  string   `json:"offset"`
}

// Interface & Generic Function
func (s Sectors) GetOffset() string {
	return s.Offset
}

func (s Sectors) SetOffset(str string) {
	s.Offset = str
}

func SectorsAppender(tmp Sectors) []db.Sector {
	var result []db.Sector

	for _, v := range tmp.Sectors {
		result = append(result, SectorTranslater(v))
	}
	return result
}

func SectorTranslater(tmp Sector) db.Sector {
	return db.Sector{
		Id:   tmp.Id,
		Name: tmp.Fields.Name,
	}
}

// Sectors
func (cli *Client) ListSectors(params url.Values) ([]db.Sector, error) {
	return ListItems[Sectors, db.Sector](cli, "Sector", params, SectorsAppender)
}
func (cli *Client) GetSector(id string) (db.Sector, error) {
	return GetItem[Sector, db.Sector](cli, "Sector", id, SectorTranslater)
}
