package eks_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestEKS(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "aws/eks")
}
