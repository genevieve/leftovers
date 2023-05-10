package route53_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRoute53(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "aws/route53")
}
