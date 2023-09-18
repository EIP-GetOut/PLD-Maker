package airtablewrapper

import (
	"encoding/json"
	"net/http"
	"net/url"
	"pld-maker/v1/internal/tools"
	"time"
)

type Client struct {
	APIpath string       `json:"airtable-api_path"`
	Token   string       `json:"airtable-api_key"`
	Client  *http.Client `json:"-"`
}

type Records interface {
	GetOffset() string
	SetOffset(str string)
}

const timeout = time.Second * 10

//type Client struct {
//    // Client fields here
//}

func NewClient(data []byte) (*Client, error) {
	var cli Client

	if err := json.Unmarshal(data, &cli); err != nil {
		return nil, err
	}
	cli.Client = &http.Client{Timeout: timeout}
	return &cli, nil
}

// Airtable call
func ListItems[T Records, G any](cli *Client, table string, params url.Values, appender func(T) []G) ([]G, error) {
	var (
		result []G
		tmp    T
		header url.Values = url.Values{"Authorization": {"Bearer " + cli.Token}}
	)

	if params == nil {
		params = url.Values{}
	}

	for {
		//Retrieve
		data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/"+table+"?"+params.Encode(), header))
		if err := json.Unmarshal(data, &tmp); err != nil {
			return nil, err
		}

		//Store
		result = append(result, appender(tmp)...)

		//Continue or Quit
		time.Sleep(250 * time.Millisecond)
		if tmp.GetOffset() != "" {
			params.Add("offset", tmp.GetOffset())
			tmp.SetOffset("")
		} else {
			break
		}
	}
	//	fmt.Println(tools.Yellow("List"+table+":"), tmp)
	//	fmt.Println(tools.Green("List"+table+":"), result)
	return result, nil
}

func GetItem[T any, G any](cli *Client, table string, id string, translater func(T) G) (G, error) {
	var (
		tmp T
	)
	//Request
	header := url.Values{}
	header.Add("Authorization", "Bearer "+cli.Token)
	data := tools.Must(tools.RequestGet(cli.Client, cli.APIpath+"/"+table+"/"+id, header))

	//Json to Struct
	if err := json.Unmarshal(data, &tmp); err != nil {
		return translater(tmp), err
	}
	//	fmt.Println(tools.Yellow("Get"+table+":"), tmp)
	//	fmt.Println(tools.Green("Get"+table+":"), translater(tmp))
	return translater(tmp), nil
}
