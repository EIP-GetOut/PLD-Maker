package epitech

import (
	"fmt"
	"net/url"
	"pld-maker/internal/airtable"
	"pld-maker/internal/tools"
)

func GetCurrentData(airtableCli airtable.Client) (airtable.Sprints, airtable.Sectors, map[string]airtable.Categories, map[string]airtable.Cards) {
	paramSprints := url.Values{"filterByFormula": {"FIND(\"In progress\", {Status})"}}
	sprints := tools.Must(airtableCli.ListSprints(paramSprints))
	if len(sprints.Sprints) < 1 {
		panic("no sprint in progress")
	} else if len(sprints.Sprints) != 1 {
		panic(fmt.Sprintf("Error: %s", "multiple sprint In progress"))
	}
	// Request Actual Sector
	paramSectors := url.Values{"filterByFormula": {""}, "sort[0][field]": {"Name"}, "sort[0][direction]": {"asc"}}
	sectors := tools.Must(airtableCli.ListSectors(paramSectors))
	// airtableCli.PrintSectors(sectors.Sectors, "")
	var categories = make(map[string]airtable.Categories)
	var cards = make(map[string]airtable.Cards)
	for _, sector := range sectors.Sectors {
		//Request Actual Categories
		paramCategories := url.Values{"filterByFormula": {fmt.Sprintf("AND(FIND(\"%s\",CONCATENATE(\"\",{Sprint})),FIND(\"%s\",CONCATENATE(\"\",{Sector})))", sprints.Sprints[0].Fields.Title, sector.Fields.Name)}}
		categories[sector.Fields.Name] = tools.Must(airtableCli.ListCategories(paramCategories))
		//		airtableCli.PrintCategories(categories[sector.Fields.Name].Categories, "%")

		//Request Actual Cards
		paramCards := url.Values{"filterByFormula": {fmt.Sprintf("AND(FIND(\"%s\",CONCATENATE(\"\",{Sprint})),FIND(\"%s\",CONCATENATE(\"\",{Sector})))", sprints.Sprints[0].Fields.Title, sector.Fields.Name)}}
		cards[sector.Fields.Name] = tools.Must(airtableCli.ListCards(paramCards))
		//		airtableCli.PrintCards(cards[sector.Fields.Name].Cards, "*")
	}
	return sprints, sectors, categories, cards
}
