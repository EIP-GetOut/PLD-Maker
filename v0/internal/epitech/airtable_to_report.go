package epitech

import (
	"pld-maker/v0/internal/airtable"
	"pld-maker/v0/internal/pld"
)

func AirtableToPldReports(AirtableReports []airtable.Report) pld.ProgressReport {
	var result pld.ProgressReport

	for _, report := range AirtableReports {
		if report.Fields.Name == "Global" {
			result.Global = report.Fields.Notes
		} else if report.Fields.Name == "Problems" {
			result.Problems = report.Fields.Notes
		} else if report.Fields.Name == "Comments" {
			result.Comments = report.Fields.Notes
		} else {
			result.Reports = append(result.Reports, pld.Report{Name: report.Fields.Name, Notes: report.Fields.Notes})
		}
	}

	return result
}
