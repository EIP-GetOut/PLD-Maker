package airtable

import (
	"fmt"
	"net/url"
	"pld-maker/v0/internal/tools"
)

func (cli *Client) GetPreviousData() (Sectors, map[string]Categories, map[string]Cards) {
	// Request Previous Sprints
	paramSprints := url.Values{"filterByFormula": {"FIND(\"Done\", {Status})"}}
	sprints := tools.Must(cli.ListSprints(paramSprints))
	if len(sprints.Sprints) == 0 {
		return /*airtable.Sprints{},*/ Sectors{}, map[string]Categories{}, map[string]Cards{}
	}
	// Request All Sector
	paramSectors := url.Values{"filterByFormula": {""}, "sort[0][field]": {"Name"}, "sort[0][direction]": {"asc"}}
	sectors := tools.Must(cli.ListSectors(paramSectors))
	//Sort (Categories & Cards) by Sectors
	var categories = make(map[string]Categories)
	var cards = make(map[string]Cards)
	for _, sprint := range sprints.Sprints {
		fmt.Println(sprint.Fields.Title, ":")
		for _, sector := range sectors.Sectors {
			//Request Previous Categories and append them to other previous sprint's categories
			paramCategories := url.Values{"filterByFormula": {fmt.Sprintf("AND(FIND(\"%s\",CONCATENATE(\"\",{Sprint})),FIND(\"%s\",CONCATENATE(\"\",{Sector})))", sprints.Sprints[0].Fields.Title, sector.Fields.Name)}}
			tmpCategories := tools.Must(cli.ListCategories(paramCategories))
			copyCategories := categories[sector.Fields.Name]
			copyCategories.Categories = append(copyCategories.Categories, tmpCategories.Categories...)
			categories[sector.Fields.Name] = copyCategories

			//Request Previous Cards and append them to other previous sprint's cards
			paramCards := url.Values{"filterByFormula": {fmt.Sprintf("AND(FIND(\"%s\",CONCATENATE(\"\",{Sprint})),FIND(\"%s\",CONCATENATE(\"\",{Sector})))", sprints.Sprints[0].Fields.Title, sector.Fields.Name)}}
			tmpCards := tools.Must(cli.ListCards(paramCards))
			copyCards := cards[sector.Fields.Name]
			copyCards.Cards = append(copyCards.Cards, tmpCards.Cards...)
			cards[sector.Fields.Name] = copyCards
		}
	}
	return sectors, categories, cards
}
