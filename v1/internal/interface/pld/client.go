package pld

import (
	"pld-maker/v1/internal/interface/db"
	"pld-maker/v1/internal/interface/pdf"
)

type Client interface {
	//File
	NewFile(filename string)
	CloseFile()
	//HeaderFooter
	Header(left, center, right string)
	Footer(left, center, right string, footerParams *pdf.FooterParams)
	//FirstPage
	FirstPage(imageFilepath, title, lowTitle string)
	//Description
	Description(title, object, author, e_mail, promo, last_update, version string)
	//Version
	Versions(versions []db.Version)
	//Summary
	Summary(versions []db.Version, schemas []db.Schema, sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card)
	//Shema
	Schemas(schemas []db.Schema)

	//ListCards   []Sectors[[]Categories[Cards]]
	ListCards(currentSprint db.Sprint, sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card)
	//Display Cards   []Sectors[Cards]
	Cards(currentSprint db.Sprint, sprints []db.Sprint, sectors []db.Sector, categories []db.Category, cards []db.Card)
	//Report
	Report(reports []db.Report)
}
