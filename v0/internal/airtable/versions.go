package airtable

import (
	"encoding/json"
	"fmt"
	"net/url"
	"pld-maker/v0/internal/tools"
	"time"
)

type VersionFields struct {
	Date     string
	Version  string
	Author   string
	Sections string
	Comments string
}

type Version struct {
	Id          string        `json:"id"`
	CreatedTime string        `json:"createdTime"`
	Fields      VersionFields `json:"fields"`
}

type Versions struct {
	Versions []Version `json:"records"`
	Offset   string    `json:"offset"`
}

func (cli *Client) ListVersions(params url.Values) (Versions, error) {
	var versions Versions
	//Request
	header := url.Values{"Authorization": {"Bearer " + cli.Token}}
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Version?"+params.Encode(), header))
	//Json to Struct
	var tmp Versions
	if err := json.Unmarshal(data, &tmp); err != nil {
		return tmp, err
	}
	versions.Versions = append(versions.Versions, tmp.Versions...)
	for tmp.Offset != "" {
		params.Add("offset", tmp.Offset)
		time.Sleep(500 * time.Millisecond)
		data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Version?"+params.Encode(), header))
		//Json to Struct
		if err := json.Unmarshal(data, &tmp); err != nil {
			return tmp, err
		}
		tmp.Offset = ""
		versions.Versions = append(versions.Versions, tmp.Versions...)
	}
	return versions, nil
}

func (cli *Client) GetVersion(id string) (Version, error) {
	var version Version
	//Request
	header := url.Values{}
	header.Add("Authorization", "Bearer "+cli.Token)
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Version/"+id, header))

	//Json to Struct
	if err := json.Unmarshal(data, &version); err != nil {
		return version, err
	}
	return version, nil
}

func (cli *Client) PrintVersions(versions []Version, indent string) {
	fmt.Println("versions:")
	fmt.Println(indent + "{")

	fmt.Println(indent + "  records: [")
	for _, version := range versions {
		cli.PrintVersion(version, indent+"    ")
	}
	fmt.Println(indent + "  ]")
	fmt.Println(indent + "}")
}

func (cli *Client) PrintVersion(version Version, indent string) {
	fields := version.Fields

	fmt.Println(indent+"{", "")
	fmt.Println(indent+"  "+"id: ", version.Id)
	fmt.Println(indent+"  "+"createdTime: ", version.CreatedTime)
	fmt.Println(indent+"  "+"Fields: {", "")
	fmt.Println(indent+"    "+"Date: ", fields.Date)
	fmt.Println(indent+"    "+"Version: ", fields.Version)
	fmt.Println(indent+"    "+"Author: ", fields.Author)
	fmt.Println(indent+"    "+"Sections: ", fields.Sections)
	fmt.Println(indent+"    "+"Comments: ", fields.Comments)
	fmt.Println(indent+"  "+"}", "")
	fmt.Println(indent+"}", "")
}
