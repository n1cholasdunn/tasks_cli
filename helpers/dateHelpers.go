package helpers

import (
	"fmt"
	"time"
)

func ValidateDate(dateStr string) error {
	if dateStr == "" {
		return nil
	}
	_, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	return nil
}

func GetTodayAsRFC3339() string {
	return time.Now().Format(time.RFC3339)
}

func GetTomorrowAsRFC3339() string {
	tomorrow := time.Now().Add(24 * time.Hour)
	return tomorrow.Format(time.RFC3339)
}

func ConvertToRFC3339(dateStr string) (string, error) {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", fmt.Errorf("invalid date format: %v", err)
	}

	return t.Format(time.RFC3339), nil
}
