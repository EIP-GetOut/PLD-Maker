package airtable

import (
	"fmt"
	"net/url"
	"pld-maker/internal/tools"
)

func (cli *Client) GetCurrentData() (Sprints, Sectors, map[string]Categories, map[string]Cards) {
	// Request Current Sprint
	paramSprints := url.Values{"filterByFormula": {"FIND(\"In progress\", {Status})"}}
	sprints := tools.Must(cli.ListSprints(paramSprints))
	if len(sprints.Sprints) < 1 {
		panic("no sprint in progress")
	} else if len(sprints.Sprints) != 1 {
		panic(fmt.Sprintf("Error: %s", "multiple sprint In progress"))
	}
	// Request Current Sector
	paramSectors := url.Values{"filterByFormula": {""}, "sort[0][field]": {"Name"}, "sort[0][direction]": {"asc"}}
	sectors := tools.Must(cli.ListSectors(paramSectors))

	//Sort (Categories & Cards) by Sectors
	var categories = make(map[string]Categories)
	var cards = make(map[string]Cards)
	for _, sector := range sectors.Sectors {
		//Request Current Categories
		paramCategories := url.Values{"filterByFormula": {fmt.Sprintf("AND(FIND(\"%s\",CONCATENATE(\"\",{Sprint})),FIND(\"%s\",CONCATENATE(\"\",{Sector})))", sprints.Sprints[0].Fields.Title, sector.Fields.Name)}}
		categories[sector.Fields.Name] = tools.Must(cli.ListCategories(paramCategories))

		//Request Current Cards
		paramCards := url.Values{"filterByFormula": {fmt.Sprintf("AND(FIND(\"%s\",CONCATENATE(\"\",{Sprint})),FIND(\"%s\",CONCATENATE(\"\",{Sector})))", sprints.Sprints[0].Fields.Title, sector.Fields.Name)}}
		cards[sector.Fields.Name] = tools.Must(cli.ListCards(paramCards))
	}
	return sprints, sectors, categories, cards
}
