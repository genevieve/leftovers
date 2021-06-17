package common

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"strings"
)

func ResourceMatches(resourceName string, filter string, regex bool) bool {
	if len(resourceName) > 0 {
		if len(filter) > 0 {
			if regex {
				regexMatcher := regexp2.MustCompile(filter, 0)
				if isMatch, err := regexMatcher.MatchString(resourceName); isMatch {
					return true
				} else if err != nil {
					fmt.Printf("error in ResourceMatches: %s\n", err)
					return false
				} else {
					return false
				}
			} else {
				return strings.Contains(resourceName, filter)
			}
		} else if len(filter) == 0 {
			return true
		}
	}
	return false
}
