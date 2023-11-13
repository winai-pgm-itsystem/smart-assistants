package utils

import (
	"regexp"
	"strconv"
)

func MacherExpenseMode(message string) (bool, []string) {
	regex := regexp.MustCompile(`^(\d+(\.\d+)?)([tfihmld])(.*)$`)
	match := regex.FindStringSubmatch(message)
	if match != nil {
		return true, match
	} else {
		return false, nil
	}
}

func MacherPPQRMode(message string) (float64, bool) {

	regex := regexp.MustCompile(`^ppqr(\d+(?:\.\d*)?)([*/]?)(\d*(?:\.\d*)?)$`)
	match := regex.FindStringSubmatch(message)

	if match != nil {

		numStr := match[1]
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return 0, false
		}

		operator := match[2]
		secondNumStr := match[3]

		if operator == "" {
			return num, true
		}

		secondNum, err := strconv.ParseFloat(secondNumStr, 64)
		if err != nil {
			return 0, false
		}

		switch operator {
		case "":
			return 0, false
		case "*":
			return num * secondNum, true
		case "/":
			if secondNum == 0 {
				return 0, false
			}
			return num / secondNum, true
		default:
			return 0, false
		}
	}

	return 0, false
}
