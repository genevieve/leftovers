package elb_test

import (
	. "github.com/onsi/gomega"

	"testing"
)

func TestELB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "aws/elb")
}
