package acceptance

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

const LEFTOVERS_ACCEPTANCE = "LEFTOVERS_ACCEPTANCE"

func TestAcceptance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "acceptance")
}
