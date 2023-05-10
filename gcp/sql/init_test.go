package sql_test

import (
	. "github.com/onsi/gomega"

	"testing"
)

func TestSql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "gcp/sql")
}
