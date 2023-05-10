package azure_test

import (
	. "github.com/onsi/gomega"

	"testing"
)

func TestAzure(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "azure")
}
