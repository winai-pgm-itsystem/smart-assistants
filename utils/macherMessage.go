package utils

import (
	"regexp"
)

func MacherExpenseMode(message string) (bool, []string) {
	// regex := regexp.MustCompile(`^(\d+)([tfihmld])(.*)$`)
	regex := regexp.MustCompile(`^(\d+(\.\d+)?)([tfihmld])(.*)$`)
	match := regex.FindStringSubmatch(message)
	if match != nil {
		return true, match
	} else {
		return false, nil
	}
}
