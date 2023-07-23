package airtable

import (
	"encoding/json"
	"fmt"
	"net/url"
	"pld-maker/internal/tools"
	"time"
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
	Assignee         []Assignee `json:"Assignees"`
	OrderedJH        string     `json:"Ordered JH"`
	Weight           int        `json:"pldWeight"`

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

func (cli *Client) ListCards(params url.Values) (Cards, error) {
	var cards Cards
	//Request
	header := url.Values{"Authorization": {"Bearer " + cli.Token}}
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Card?"+params.Encode(), header))
	//Json to Struct
	var tmp Cards
	if err := json.Unmarshal(data, &tmp); err != nil {
		return tmp, err
	}
	cards.Cards = append(cards.Cards, tmp.Cards...)
	for tmp.Offset != "" {
		params.Add("offset", tmp.Offset)
		time.Sleep(500 * time.Millisecond)
		data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Card?"+params.Encode(), header))

		if err := json.Unmarshal(data, &tmp); err != nil {
			return tmp, err
		}
		tmp.Offset = ""
		cards.Cards = append(cards.Cards, tmp.Cards...)
	}
	return cards, nil
}

func (cli *Client) GetCard(id string) (Card, error) {
	var card Card
	//Request
	header := url.Values{}
	header.Add("Authorization", "Bearer "+cli.Token)
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Card/"+id, header))

	//Json to Struct
	if err := json.Unmarshal(data, &card); err != nil {
		return card, err
	}
	return card, nil
}

func (cli *Client) PrintCards(cards []Card, indent string) {
	fmt.Println("cards:")
	fmt.Println(indent + "{")
	fmt.Println(indent + "  records: [")
	for _, card := range cards {
		cli.PrintCard(card, indent+"    ")
	}
	fmt.Println(indent + "  ]")
	fmt.Println(indent + "}")
}

func (cli *Client) PrintCard(card Card, indent string) {
	fields := card.Fields

	fmt.Println(indent+"{", "")
	fmt.Println(indent+"  "+"id: ", card.Id)
	fmt.Println(indent+"  "+"createdTime: ", card.CreatedTime)
	fmt.Println(indent+"  "+"Fields: {", "")
	fmt.Println(indent+"    "+"title: ", fields.Title)
	fmt.Println(indent+"    "+"progress: ", fields.Progress)
	fmt.Println(indent+"    "+"AsWho: ", fields.AsWho)
	fmt.Println(indent+"    "+"IWant: ", fields.IWant)
	fmt.Println(indent+"    "+"Description: ", fields.Description)
	fmt.Println(indent+"    "+"DefinitionOfDone: ", fields.DefinitionOfDone)
	fmt.Println(indent+"    "+"Jh: ", fields.Jh)
	fmt.Println(indent+"    "+"Assignee: ", fields.Assignee)
	fmt.Println(indent+"    "+"OrderedJH: ", fields.OrderedJH)
	fmt.Println(indent+"    "+"Category: ", fields.Category)
	fmt.Println(indent+"    "+"Secteur: ", fields.Secteur)
	fmt.Println(indent+"    "+"Sprint: ", fields.Sprint)
	fmt.Println(indent+"  "+"}", "")
	fmt.Println(indent+"}", "")
}
