package db

import (
	"net/url"
)

type Client interface {
	//Version
	ListVersions(params url.Values) ([]Version, error)
	GetVersion(id string) (Version, error)

	//Schema
	ListSchemas(params url.Values) ([]Schema, error)
	GetSchema(id string) (Schema, error)

	//Sprints
	ListSprints(params url.Values) ([]Sprint, error)
	GetSprint(id string) (Sprint, error)

	//Sectors
	ListSectors(params url.Values) ([]Sector, error)
	GetSector(id string) (Sector, error)

	//Categories
	ListCategories(params url.Values) ([]Category, error)
	GetCategory(id string) (Category, error)

	//Report
	ListCards(params url.Values) ([]Card, error)
	GetCard(id string) (Card, error)

	//Report
	ListReports(params url.Values) ([]Report, error)
	GetReport(id string) (Report, error)
}

//	PrintVersions(versions []Version)
//	PrintVersion(version Version)
//	PrintSchemas(schemas []Schema)
//	PrintSchema(schema Schema)
//	PrintSprints(sprints []Sprint)
//	PrintSprint(sprint Sprint)
//	PrintSectors(sectors []Sector)
//	PrintSector(sector Sector)
// PrintCategories(categories []Category)
// PrintCategory(category Category)
