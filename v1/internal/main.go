package main

import (
	"pld-maker/v1/internal/fpdfwrapper"
	"pld-maker/v1/internal/interface/pdf"
	"pld-maker/v1/internal/tools"
)

func main() {
	//credential := tools.Must(os.ReadFile("./conf/credential.json"))
	//dbCli := db.Client(tools.Must(airtable.NewClient(credential)))
	pdfCli := pdf.Client(tools.Must(fpdfwrapper.NewClient()))
	// pldCli := pld.Client(tools.Must(epipld.NewClient()))
	pdfCli.NewFile("hello")
	pdfCli.Header("string", "", "")
	pdfCli.CloseFile()
}
