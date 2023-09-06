package fpdfwrapper

import (
	"fmt"
	"pld-maker/v1/internal/interface/pdf"
)

func (cli *Client) Title(str string, params *pdf.TextParams) {
	fmt.Println("Title")
}

func (cli *Client) SubTitle(str string, params *pdf.TextParams) {
	fmt.Println("SubTitle")
}

func (cli *Client) Heading1(str string, params *pdf.TextParams) {
	fmt.Println("Heading1")
}

func (cli *Client) Heading2(str string, params *pdf.TextParams) {
	fmt.Println("Heading2")
}

func (cli *Client) Text(str string, params *pdf.TextParams) {
	fmt.Println("Text")
}
