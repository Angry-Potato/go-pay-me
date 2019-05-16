package schema

import "regexp"

func contains(haystack []string, needle string) bool {
	for _, hay := range haystack {
		if hay == needle {
			return true
		}
	}
	return false
}

func isUUID(uuid string) bool {
	if matches, err := regexp.MatchString("^[a-zA-Z0-9-]+$", uuid); !matches || err != nil {
		return false
	}
	return true
}

func isAmount(amount string) bool {
	if matches, err := regexp.MatchString("^[0-9]+(\\.[0-9][0-9])?$", amount); !matches || err != nil {
		return false
	}
	return true
}
