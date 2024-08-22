package helpers

import (
	"regexp"
	"strings"
)

func PostUrlCreator(title string) (string, error) {
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return " ", err
	}

	url := reg.ReplaceAllString(title, "-")
	url = strings.ToLower(strings.Trim(url, "-"))

	return url, nil
}
