package openstack

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCompute(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "openstack")
}
