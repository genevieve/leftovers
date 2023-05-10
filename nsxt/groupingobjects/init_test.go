package groupingobjects_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGroupingobjects(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "nsxt/groupingobjects")
}
