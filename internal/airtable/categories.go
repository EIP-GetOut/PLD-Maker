package airtable

import (
	"encoding/json"
	"fmt"
	"net/url"
	"pld-maker/internal/tools"
	"time"
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
	Offset     string     `json:"offset"`
}

func (cli *Client) ListCategories(params url.Values) (Categories, error) {
	var categories Categories
	//Request
	header := url.Values{"Authorization": {"Bearer " + cli.Token}}
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Category?"+params.Encode(), header))
	//Json to Struct
	var tmp Categories
	if err := json.Unmarshal(data, &tmp); err != nil {
		return tmp, err
	}
	categories.Categories = append(categories.Categories, tmp.Categories...)
	for tmp.Offset != "" {
		params.Add("offset", tmp.Offset)
		time.Sleep(500 * time.Millisecond)
		data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Category?"+params.Encode(), header))
		//Json to Struct
		if err := json.Unmarshal(data, &tmp); err != nil {
			return tmp, err
		}
		tmp.Offset = ""
		categories.Categories = append(categories.Categories, tmp.Categories...)
	}
	return categories, nil
}

func (cli *Client) GetCategory(id string) (Category, error) {
	var category Category
	//Request
	header := url.Values{}
	header.Add("Authorization", "Bearer "+cli.Token)
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Category/"+id, header))

	//Json to Struct
	if err := json.Unmarshal(data, &category); err != nil {
		return category, err
	}
	return category, nil
}

func (cli *Client) PrintCategories(categories []Category, indent string) {
	fmt.Println("categories:")
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
