package storage_test

import (
	. "github.com/onsi/gomega"

	"testing"
)

func TestStorage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "gcp/storage")
}
