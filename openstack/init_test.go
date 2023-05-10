package openstack

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCompute(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "openstack")
}

//go:generate faux --package github.com/gophercloud/gophercloud/pagination --interface Page --output fakes/page.go
