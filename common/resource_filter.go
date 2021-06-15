package common

import (
	"regexp"
	"strings"
)

func ResourceMatches(resourceName string, filter string, regex bool) bool {
	if len(resourceName) > 0 {
		if len(filter) > 0 {
			if regex {
				regexMatcher := regexp.MustCompile(filter)
				return regexMatcher.MatchString(resourceName)
			} else {
				return strings.Contains(resourceName, filter)
			}

		} else if len(filter) == 0 {
			return true
		}
	}
	return false
}
