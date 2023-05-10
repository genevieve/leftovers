package iam_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIam(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "gcp/iam")
}
