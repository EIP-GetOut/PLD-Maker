package airtable

import (
	"encoding/json"
	"fmt"
	"net/url"
	"pld-maker/internal/tools"
)

type CategoryFields struct {
	Name  string   `json:"Name"`
	Cards []string `json:"Card"`
}

type Category struct {
	Id          string         `json:"id"`
	CreatedTime string         `json:"createdTime"`
	Fields      CategoryFields `json:"fields"`
}

type Categories struct {
	Categories []Category `json:"records"`
}

func (cli *Client) ListCategories(params *url.Values) (Categories, error) {
	var categories Categories
	var parameters string
	//Request
	if params != nil {
		parameters = "?" + (*params).Encode()
	}
	header := url.Values{}
	header.Add("Authorization", "Bearer "+cli.Token)
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Category"+parameters, header))
	//Json to Struct
	if err := json.Unmarshal(data, &categories); err != nil {
		return categories, err
	}
	return categories, nil
}

func (cli *Client) GetCategory(id string) (Category, error) {
	var category Category
	//Request
	header := url.Values{}
	header.Add("Authorization", "Bearer "+cli.Token)
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Cards/"+id, header))

	//Json to Struct
	if err := json.Unmarshal(data, &category); err != nil {
		return category, err
	}
	return category, nil
}

func (cli *Client) PrintCategories(categories []Category, indent string) {
	fmt.Println(indent + "{")

	fmt.Println(indent + "  records: [")
	for _, category := range categories {
		cli.PrintCategory(category, indent+"    ")
	}
	fmt.Println(indent + "  ]")
	fmt.Println(indent + "}")
}

func (cli *Client) PrintCategory(category Category, indent string) {
	fields := category.Fields

	fmt.Println(indent+"{", "")
	fmt.Println(indent+"  "+"id: ", category.Id)
	fmt.Println(indent+"  "+"createdTime: ", category.CreatedTime)
	fmt.Println(indent+"  "+"Fields: {", "")
	fmt.Println(indent+"    "+"Name: ", fields.Name)
	fmt.Println(indent+"    "+"Cards: ", fields.Cards)
	fmt.Println(indent+"  "+"}", "")
	fmt.Println(indent+"}", "")
}
