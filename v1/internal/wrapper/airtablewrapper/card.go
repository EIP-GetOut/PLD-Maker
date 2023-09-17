package airtablewrapper

import (
	"net/url"
	"pld-maker/v1/internal/interface/db"
)

type Assignee struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type CardFields struct {
	Title            string     `json:"Name"`
	Status           string     `json:"Status"`
	Progress         float64    `json:"progress"`
	AsWho            string     `json:"En tant que"`
	IWant            string     `json:"Je veux"`
	DefinitionOfDone string     `json:"Definition Of Done"`
	Description      string     `json:"Description"`
	Jh               float64    `json:"JH"`
	Assignees        []Assignee `json:"Assignees"`
	OrderedJH        string     `json:"Ordered JH"`

	Category []string `json:"Category"`
	Sprint   []string `json:"Sprint"`
	Secteur  []string `json:"Secteur"`
}

type Card struct {
	Id          string     `json:"id"`
	CreatedTime string     `json:"createdTime"`
	Fields      CardFields `json:"fields"`
}

type Cards struct {
	Cards  []Card `json:"records"`
	Offset string `json:"offset"`
}

// Interface & Generic Function
func (s Cards) GetOffset() string {
	return s.Offset
}

func (s Cards) SetOffset(str string) {
	s.Offset = str
}

func CardsAppender(tmp Cards) []db.Card {
	var result []db.Card

	for _, v := range tmp.Cards {
		result = append(result, CardTranslater(v))
	}
	return result
}

func CardTranslater(tmp Card) db.Card {
	var assignees []db.Assignee

	for _, v := range tmp.Fields.Assignees {
		assignees = append(assignees, db.Assignee{Name: v.Name, Email: v.Email})
	}
	return db.Card{
		Title:            tmp.Fields.Title,
		Status:           tmp.Fields.Status,
		Progress:         tmp.Fields.Progress,
		AsWho:            tmp.Fields.AsWho,
		IWant:            tmp.Fields.IWant,
		DefinitionOfDone: tmp.Fields.DefinitionOfDone,
		Description:      tmp.Fields.Description,
		Jh:               tmp.Fields.Jh,
		Assignees:        assignees,
		OrderedJH:        tmp.Fields.OrderedJH,
		Category:         tmp.Fields.Category,
		Sprint:           tmp.Fields.Sprint,
		Secteur:          tmp.Fields.Secteur,
	}
}

// Cards
func (cli *Client) ListCards(params url.Values) ([]db.Card, error) {
	return ListItems[Cards, db.Card](cli, "Card", params, CardsAppender)
}

func (cli *Client) GetCard(id string) (db.Card, error) {
	return GetItem[Card, db.Card](cli, "Card", id, CardTranslater)
}
