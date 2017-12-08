package awsec2_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAwsec2(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "awsec2")
}
