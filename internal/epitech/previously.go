package epitech

import (
	"fmt"
	"net/url"
	"pld-maker/internal/airtable"
	"pld-maker/internal/tools"
)

func GetPreviousData(airtableCli airtable.Client) ( /*airtable.Sprints,*/ airtable.Sectors, map[string]airtable.Categories, map[string]airtable.Cards) {
	paramSprints := url.Values{"filterByFormula": {"FIND(\"Done\", {Status})"}}
	sprints := tools.Must(airtableCli.ListSprints(paramSprints))
	if len(sprints.Sprints) == 0 {
		return /*airtable.Sprints{},*/ airtable.Sectors{}, map[string]airtable.Categories{}, map[string]airtable.Cards{}
	}
	// Request Actual Sector
	paramSectors := url.Values{"filterByFormula": {""}, "sort[0][field]": {"Name"}, "sort[0][direction]": {"asc"}}
	sectors := tools.Must(airtableCli.ListSectors(paramSectors))
	// airtableCli.PrintSectors(sectors.Sectors, "")
	var categories = make(map[string]airtable.Categories)
	var cards = make(map[string]airtable.Cards)
	for _, sprint := range sprints.Sprints {
		fmt.Println(sprint.Fields.Title, ":")
		for _, sector := range sectors.Sectors {
			//Request Previous Categories and append them to other previous sprint's categories
			paramCategories := url.Values{"filterByFormula": {fmt.Sprintf("AND(FIND(\"%s\",CONCATENATE(\"\",{Sprint})),FIND(\"%s\",CONCATENATE(\"\",{Sector})))", sprints.Sprints[0].Fields.Title, sector.Fields.Name)}}
			tmpCategories := tools.Must(airtableCli.ListCategories(paramCategories))
			copyCategories := categories[sector.Fields.Name]
			copyCategories.Categories = append(copyCategories.Categories, tmpCategories.Categories...)
			categories[sector.Fields.Name] = copyCategories

			paramCards := url.Values{"filterByFormula": {fmt.Sprintf("AND(FIND(\"%s\",CONCATENATE(\"\",{Sprint})),FIND(\"%s\",CONCATENATE(\"\",{Sector})))", sprints.Sprints[0].Fields.Title, sector.Fields.Name)}}
			tmpCards := tools.Must(airtableCli.ListCards(paramCards))
			copyCards := cards[sector.Fields.Name]
			copyCards.Cards = append(copyCards.Cards, tmpCards.Cards...)
			cards[sector.Fields.Name] = copyCards
		}
	}
	return /*sprints,*/ sectors, categories, cards
}
