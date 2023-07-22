package epitech

import "pld-maker/internal/airtable"

func MergeCategories(categoriesMaps ...map[string]airtable.Categories) map[string]airtable.Categories {
	result := make(map[string]airtable.Categories)
	for _, categoriesMap := range categoriesMaps {
		for key, value := range categoriesMap {
			airtableCategories := result[key]
			airtableCategories.Categories = append(result[key].Categories, value.Categories...)
			result[key] = airtableCategories
		}
	}
	return result
}
