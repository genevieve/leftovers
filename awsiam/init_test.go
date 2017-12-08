package awsiam_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAwsiam(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "awsiam")
}
