package main

import (
	"pld-maker/v1/test/wrapper/pdf"
	"pld-maker/v1/test/wrapper/pld"
)

func main() {
	pdf := pdf.Pdf(&pld.Client{})
	//	pdf.PrintName()
	pdf.SetName("hey")
	pdf.PrintName()
}
