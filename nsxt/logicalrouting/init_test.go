package logicalrouting

import (
	. "github.com/onsi/gomega"

	"testing"
)

func TestDns(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "nsxt/logicalrouting")
}
