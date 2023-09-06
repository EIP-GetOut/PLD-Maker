package db

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
