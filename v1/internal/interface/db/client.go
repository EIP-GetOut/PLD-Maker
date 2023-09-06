package db

import (
	"net/url"
)

type Client interface {
	//Sprints
	ListSprints(params url.Values) []Sprint
	GetSprint(id string) Sprint
	PrintSprints(sprints []Sprint)
	PrintSprint(sprint Sprint)

	//Sectors
	ListSectors(params url.Values) []Sector
	GetSector(id string) Sector
	PrintSectors(sectors []Sector)
	PrintSector(sector Sector)

	//Categories
	GetCategories(params url.Values) []Category
	GetCategory(id string) []Category
	PrintCategories(categories []Category)
	PrintCategory(category Category)

	//Cards
	ListCards(params url.Values) []Card
	GetCard(id string) Card
	PrintCards(cards []Card)
	PrintCard(card Card)
}
