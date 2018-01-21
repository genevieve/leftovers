package acceptance

import (
	"os"
	"strings"
)

func ReadyToTest() bool {
	iaas := os.Getenv("LEFTOVERS_ACCEPTANCE")
	if iaas == "" {
		return false
	}

	if strings.ToLower(iaas) == "gcp" {
		return true
	}

	return false
}
