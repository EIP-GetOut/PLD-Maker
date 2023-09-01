package main

import (
	"pld-maker/test/pdf"
	"pld-maker/test/pld"
)

func main() {
	pdf := pdf.Pdf(&pld.Client{})
	//	pdf.PrintName()
	pdf.SetName("hey")
	pdf.PrintName()
}
