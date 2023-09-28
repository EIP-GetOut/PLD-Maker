package epitech

import (
	"pld-maker/v0/internal/airtable"
	"pld-maker/v0/internal/pld"
)

func AirtableToPldVersion(airtableVersions []airtable.Version) []pld.Version {
	var pldVersions []pld.Version
	for _, airVersion := range airtableVersions {
		pldVersions = append(pldVersions, pld.Version{
			Date:     airVersion.Fields.Date,
			Version:  airVersion.Fields.Version,
			Author:   airVersion.Fields.Author,
			Sections: airVersion.Fields.Sections,
			Comments: airVersion.Fields.Comments,
		})
	}
	return pldVersions
}
