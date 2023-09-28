package airtable

import (
	"encoding/json"
	"fmt"
	"net/url"
	"pld-maker/v0/internal/tools"
	"time"
)

type ReportFields struct {
	Name  string `json:"Name"`
	Notes string `json:"Notes"`
}

type Report struct {
	Id          string       `json:"id"`
	CreatedTime string       `json:"createdTime"`
	Fields      ReportFields `json:"fields"`
}

type Reports struct {
	Reports []Report `json:"records"`
	Offset  string   `json:"offset"`
}

func (cli *Client) ListReports(params url.Values) (Reports, error) {
	var reports Reports
	//Request
	header := url.Values{"Authorization": {"Bearer " + cli.Token}}
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Report?"+params.Encode(), header))
	//Json to Struct
	var tmp Reports
	if err := json.Unmarshal(data, &tmp); err != nil {
		return tmp, err
	}
	reports.Reports = append(reports.Reports, tmp.Reports...)
	for tmp.Offset != "" {
		params.Add("offset", tmp.Offset)
		time.Sleep(500 * time.Millisecond)
		data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Report?"+params.Encode(), header))

		if err := json.Unmarshal(data, &tmp); err != nil {
			return tmp, err
		}
		tmp.Offset = ""
		reports.Reports = append(reports.Reports, tmp.Reports...)
	}
	return reports, nil
}

func (cli *Client) GetReport(id string) (Report, error) {
	var report Report
	//Request
	header := url.Values{}
	header.Add("Authorization", "Bearer "+cli.Token)
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/Report/"+id, header))

	//Json to Struct
	if err := json.Unmarshal(data, &report); err != nil {
		return report, err
	}
	return report, nil
}

func (cli *Client) PrintReports(reports []Report, indent string) {
	fmt.Println("reports:")
	fmt.Println(indent + "{")
	fmt.Println(indent + "  records: [")
	for _, report := range reports {
		cli.PrintReport(report, indent+"    ")
	}
	fmt.Println(indent + "  ]")
	fmt.Println(indent + "}")
}

func (cli *Client) PrintReport(report Report, indent string) {
	fields := report.Fields

	fmt.Println(indent+"{", "")
	fmt.Println(indent+"  "+"id: ", report.Id)
	fmt.Println(indent+"  "+"createdTime: ", report.CreatedTime)
	fmt.Println(indent+"  "+"Fields: {", "")
	fmt.Println(indent+"    "+"name: ", fields.Name)
	fmt.Println(indent+"    "+"notes: ", fields.Notes)
	fmt.Println(indent+"  "+"}", "")
	fmt.Println(indent+"}", "")
}
