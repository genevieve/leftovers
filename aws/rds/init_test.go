package rds_test

import (
	. "github.com/onsi/gomega"

	"testing"
)

func TestRDS(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "aws/rds")
}
