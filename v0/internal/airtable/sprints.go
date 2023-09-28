package airtable

import (
	"encoding/json"
	"fmt"
	"net/url"
	"pld-maker/v0/internal/tools"
	"time"
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
	Offset  string   `json:"offset"`
}

func (cli *Client) ListSprints(params url.Values) (Sprints, error) {
	var sprints Sprints
	//Request
	header := url.Values{"Authorization": {"Bearer " + cli.Token}}
	fmt.Println(cli.APIpath + "/Sprint?" + params.Encode())
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Sprint?"+params.Encode(), header))
	//Json to Struct
	var tmp Sprints
	if err := json.Unmarshal(data, &tmp); err != nil {
		return tmp, err
	}
	sprints.Sprints = append(sprints.Sprints, tmp.Sprints...)
	for tmp.Offset != "" {
		params.Add("offset", tmp.Offset)
		time.Sleep(500 * time.Millisecond)
		data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Sprint?"+params.Encode(), header))
		if err := json.Unmarshal(data, &tmp); err != nil {
			return tmp, err
		}
		tmp.Offset = ""
		sprints.Sprints = append(sprints.Sprints, tmp.Sprints...)
	}
	return sprints, nil
}

func (cli *Client) GetSprint(id string) (Sprint, error) {
	var sprint Sprint
	//Request
	header := url.Values{}
	header.Add("Authorization", "Bearer "+cli.Token)
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Sprint/"+id, header))

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
