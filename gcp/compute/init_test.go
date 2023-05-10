package compute_test

import (
	. "github.com/onsi/gomega"

	"testing"
)

func TestCompute(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "gcp/compute")
}
