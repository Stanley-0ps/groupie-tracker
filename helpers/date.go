package helpers

import "time"

// FormatDate converts the API's day-month-year date strings into a readable
// format. Returning the original value keeps unexpected API data visible.
func FormatDate(date string) string {
	parsedDate, err := time.Parse("02-01-2006", date)
	if err != nil {
		return date
	}

	return parsedDate.Format("January 2, 2006")
}
