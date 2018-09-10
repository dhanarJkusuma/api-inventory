package helper

import (
	"inventory_app/models"
	"time"
)

func GetCurrentDateWithFormat(format string) string {
	now := time.Now()
	return now.Format(format)
}

func IsCorrectDateFormat(format string, date string) error {
	_, err := time.Parse(format, date)
	if err != nil {
		return models.ERR_DATE_PARSING
	}
	return nil
}

func TimeToStringFormat(dt time.Time, pattern string) string {
	return dt.Format(pattern)
}
