package airtablewrapper

import (
	"net/url"
	"pld-maker/v1/internal/interface/db"
)

type CategoryFields struct {
	Name    string   `json:"Name"`
	Cards   []string `json:"Card"`
	Sprints []string `json:"Sprint"`
	Sectors []string `json:"Sector"`
}

type Category struct {
	Id          string         `json:"id"`
	CreatedTime string         `json:"createdTime"`
	Fields      CategoryFields `json:"fields"`
}

type Categories struct {
	Categories []Category `json:"records"`
	Offset     string     `json:"offset"`
}

// Interface & Generic Function
func (s Categories) GetOffset() string {
	return s.Offset
}

func (s Categories) SetOffset(str string) {
	s.Offset = str
}

func CategoriesAppender(tmp Categories) []db.Category {
	var result []db.Category

	for _, v := range tmp.Categories {
		result = append(result, CategoryTranslater(v))
	}
	return result
}

func CategoryTranslater(tmp Category) db.Category {
	return db.Category{
		Id:      tmp.Id,
		Name:    tmp.Fields.Name,
		Cards:   tmp.Fields.Cards,
		Sectors: tmp.Fields.Sectors,
		Sprints: tmp.Fields.Sprints,
	}
}

// Categories
func (cli *Client) ListCategories(params url.Values) ([]db.Category, error) {
	return ListItems[Categories, db.Category](cli, "Category", params, CategoriesAppender)
}

func (cli *Client) GetCategory(id string) (db.Category, error) {
	return GetItem[Category, db.Category](cli, "Category", id, CategoryTranslater)

}
