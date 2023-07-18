package airtable

import (
	"encoding/json"
	"fmt"
	"net/url"
	"pld-maker/internal/tools"
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
}

func (cli *Client) ListSectors(params *url.Values) (Sectors, error) {
	var sectors Sectors
	var parameters string
	//Request
	if params != nil {
		parameters = "?" + (*params).Encode()
	}
	header := url.Values{}
	header.Add("Authorization", "Bearer "+cli.Token)
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Sector"+parameters, header))
	//Json to Struct
	if err := json.Unmarshal(data, &sectors); err != nil {
		return sectors, err
	}
	return sectors, nil
}

func (cli *Client) GetSector(id string) (Sector, error) {
	var sector Sector
	//Request
	header := url.Values{}
	header.Add("Authorization", "Bearer "+cli.Token)
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Cards/"+id, header))

	//Json to Struct
	if err := json.Unmarshal(data, &sector); err != nil {
		return sector, err
	}
	return sector, nil
}

func (cli *Client) PrintSectors(sectors []Sector, indent string) {
	fmt.Println("sectors:")
	fmt.Println(indent + "{")

	fmt.Println(indent + "  records: [")
	for _, sector := range sectors {
		cli.PrintSector(sector, indent+"    ")
	}
	fmt.Println(indent + "  ]")
	fmt.Println(indent + "}")
}

func (cli *Client) PrintSector(sector Sector, indent string) {
	fields := sector.Fields

	fmt.Println(indent+"{", "")
	fmt.Println(indent+"  "+"id: ", sector.Id)
	fmt.Println(indent+"  "+"createdTime: ", sector.CreatedTime)
	fmt.Println(indent+"  "+"Fields: {", "")
	fmt.Println(indent+"    "+"Name: ", fields.Name)
	fmt.Println(indent+"  "+"}", "")
	fmt.Println(indent+"}", "")
}
