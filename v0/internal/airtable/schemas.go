package airtable

import (
	"encoding/json"
	"fmt"
	"net/url"
	"pld-maker/v0/internal/tools"
	"time"
)

/*{
"records": [
	{
			"id": "recrKpE5gTq1lmaa6",
			"createdTime": "2023-07-16T16:04:48.000Z",
			"fields": {
					"File": [
							{
									"id": "atte7LbGUDYs9VKoQ",
									"width": 992,
									"height": 1028,
									"url": "https://v5.airtableusercontent.com/v1/20/20/1694880000000/QjBDs8Bx0ay6csYpQ6ebMQ/qbDEupysDgnoCx4LyhtPYFpkGQquq2NvfTyeVckAqB7peWluCvc4MM5Ohsv1PRrvfjHBuxTwQpNqcvIaB_5eZP_1E-hY_WKmeGSx1g_ZaaA_VtQZ9H1o1mjByZqe9FQX/eqmmZ3V-eCj_z9-neT-LaQykFcEqzqtS0NvBNlzSB4s",
									"filename": "Screenshot 2023-07-23 at 16.37.43.png",
									"size": 111269,
									"type": "image/png",
									"thumbnails": {
											"small": {
													"url": "https://v5.airtableusercontent.com/v1/20/20/1694880000000/Zdx40gg-jAv3yrIupZaNzQ/rB6SZaFmu-gM3KXmpxTlrQkAhiegCOiXyNc0B6LlxaPQqDZAjkoCf00APTEfj72ifiPKZkSWmK9tXlgXnD-wmg/lH5lNVMT00Y7IDVWvy_JCFT_JStec5to6mhV_rfZONw",
													"width": 35,
													"height": 36
											},
											"large": {
													"url": "https://v5.airtableusercontent.com/v1/20/20/1694880000000/fBe3eVyqdPNrPm3xditL6w/hq-EJYl5P_wdtlz-QJ6vMk72lzY0U4tsNNcyf1DfplCyeDCfJE7MAOwK7Gt4pT42qpo6rvHTCNvein3gaTUakQ/75XEFl5Oc-kWoE1FR2KVPKibhfAPeCcI823kvdOJHtc",
													"width": 512,
													"height": 531
											},
											"full": {
													"url": "https://v5.airtableusercontent.com/v1/20/20/1694880000000/7RZxLyqmdJsvB74n47Ac6g/-cdbqWgC1C3AFWxRJ8fZdzfFkkPx-7y3DNE0huxee2yj9WczZrsRY7QLq2fZUBZBSi9U-C_ksjrkzW0yrE3YGw/SYgVSmdU_zanytcd1tUEfDNQFdoSwXAFra9pUEqzD9c",
													"width": 3000,
													"height": 3000
											}
									}
							}
					],
					"Name": "Perry"
			}
	}
],
"offset": "itrp3aNxd4ZspzXWy/recrKpE5gTq1lmaa6"
}*/

type File struct {
	Id     string `json:"id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Url    string `json:"url"`
}

type SchemaFields struct {
	Files []File `json:"File"`
	Title string `json:"Name"`
}

type Schema struct {
	Id          string       `json:"id"`
	CreatedTime string       `json:"createdTime"`
	Fields      SchemaFields `json:"fields"`
}

type Schemas struct {
	Schemas []Schema `json:"records"`
	Offset  string   `json:"offset"`
}

func (cli *Client) ListSchemas(params url.Values) (Schemas, error) {
	var schemas Schemas
	//Request
	header := url.Values{"Authorization": {"Bearer " + cli.Token}}
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Schema?"+params.Encode(), header))
	//Json to Struct
	var tmp Schemas
	if err := json.Unmarshal(data, &tmp); err != nil {
		return tmp, err
	}
	schemas.Schemas = append(schemas.Schemas, tmp.Schemas...)
	for tmp.Offset != "" {
		params.Add("offset", tmp.Offset)
		time.Sleep(500 * time.Millisecond)
		data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Schema?"+params.Encode(), header))

		if err := json.Unmarshal(data, &tmp); err != nil {
			return tmp, err
		}
		tmp.Offset = ""
		schemas.Schemas = append(schemas.Schemas, tmp.Schemas...)
	}
	return schemas, nil
}

func (cli *Client) GetSchema(id string) (Schema, error) {
	var schema Schema
	//Request
	header := url.Values{}
	header.Add("Authorization", "Bearer "+cli.Token)
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Schema/"+id, header))

	//Json to Struct
	if err := json.Unmarshal(data, &schema); err != nil {
		return schema, err
	}
	return schema, nil
}

func (cli *Client) PrintSchemas(schemas []Schema, indent string) {
	fmt.Println("schemas:")
	fmt.Println(indent + "{")
	fmt.Println(indent + "  records: [")
	for _, schema := range schemas {
		cli.PrintSchema(schema, indent+"    ")
	}
	fmt.Println(indent + "  ]")
	fmt.Println(indent + "}")
}

func (cli *Client) PrintSchema(card Schema, indent string) {
	fields := card.Fields

	fmt.Println(indent+"{", "")
	fmt.Println(indent+"  "+"id: ", card.Id)
	fmt.Println(indent+"  "+"createdTime: ", card.CreatedTime)
	fmt.Println(indent+"  "+"Fields: {", "")
	fmt.Println(indent+"    "+"title: ", fields.Title)
	fmt.Println(indent+"    "+"progress: ", fields.Files)
	fmt.Println(indent+"    "+"AsWho: ", fields.Title)
	fmt.Println(indent+"  "+"}", "")
	fmt.Println(indent+"}", "")
}
