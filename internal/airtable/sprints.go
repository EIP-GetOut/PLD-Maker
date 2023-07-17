package airtable

import (
	"encoding/json"
	"fmt"
	"net/url"
	"pld-maker/internal/tools"
)

type SprintFields struct {
	Title  string   `json:"Name"`
	Cards  []string `json:"Card"`
	Status string   `json:"Status"`
	Number int      `json:"Number"`
}

type Sprint struct {
	Id          string       `json:"id"`
	CreatedTime string       `json:"createdTime"`
	Fields      SprintFields `json:"fields"`
}

type Sprints struct {
	Sprints []Sprint `json:"records"`
}

func (cli *Client) ListSprints(params *url.Values) (Sprints, error) {
	var sprints Sprints
	var parameters string

	if params != nil {
		parameters = "?" + (*params).Encode()
	}
	header := url.Values{}
	header.Add("Authorization", "Bearer "+cli.Token)
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Sprint"+parameters, header))
	//Json to Struct
	if err := json.Unmarshal(data, &sprints); err != nil {
		return sprints, err
	}
	return sprints, nil
}

func (cli *Client) GetSprint(id string) (Sprint, error) {
	var sprint Sprint
	//Request
	header := url.Values{}
	header.Add("Authorization", "Bearer "+cli.Token)
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Cards/"+id, header))

	//Json to Struct
	if err := json.Unmarshal(data, &sprint); err != nil {
		return sprint, err
	}
	return sprint, nil
}

func (cli *Client) PrintSprints(sprints []Sprint, indent string) {
	fmt.Println("sprints:")
	fmt.Println(indent + "{")
	fmt.Println(indent + "  records: [")
	for _, sprint := range sprints {
		cli.PrintSprint(sprint, indent+"    ")
	}
	fmt.Println(indent + "  ]")
	fmt.Println(indent + "}")
}

func (cli *Client) PrintSprint(sprint Sprint, indent string) {
	fields := sprint.Fields

	fmt.Println("sprint:")
	fmt.Println(indent+"{", "")
	fmt.Println(indent+"  "+"id: ", sprint.Id)
	fmt.Println(indent+"  "+"createdTime: ", sprint.CreatedTime)
	fmt.Println(indent+"  "+"Fields: {", "")
	fmt.Println(indent+"    "+"title: ", fields.Title)
	fmt.Println(indent+"    "+"status: ", fields.Status)
	fmt.Println(indent+"    "+"cards: ", fields.Cards)
	fmt.Println(indent+"  "+"}", "")
	fmt.Println(indent+"}", "")
}
