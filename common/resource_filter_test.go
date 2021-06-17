package common

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RegexMatch", func() {
	It("doesn't match", func() {
		Expect(ResourceMatches("my-resource-name", "not-my-resource", false)).To(BeFalse())
	})

	It("does match", func() {
		Expect(ResourceMatches("my-resource-name", "my-res", false)).To(BeTrue())
	})

	It("matches for empty filter", func() {
		Expect(ResourceMatches("my-resource-name", "", false)).To(BeTrue())
	})

	Context("interpreting filter as regex", func() {
		It("doesn't match", func() {
			Expect(ResourceMatches("my-resource-name", "not-.*-resource", true)).To(BeFalse())
		})

		It("does match", func() {
			Expect(ResourceMatches("my-resource-name", "my-res.*-name", true)).To(BeTrue())
		})

		It("matches for empty filter", func() {
			Expect(ResourceMatches("my-resource-name", "", true)).To(BeTrue())
		})
	})

	Context("handling extended regex", func() {
		filter := `^(?=.*pull-1611)(?!.*iless-).*`
		Context("given a bosh deployed vm resource", func() {
			It("filters in the positive case", func() {
				name := "vm-ca9e5f2d-7ae9-40b8-66b3-22fc49318a75 (pull-1611-pcf-network, cf-8bf3da062289a6c57468, cf-8bf3da062289a6c57468-compute, compute, p-bosh, p-bosh-cf-8bf3da062289a6c57468, p-bosh-cf-8bf3da062289a6c57468-compute, pull-1611-vms)"
				Expect(ResourceMatches(name, filter, true)).To(BeTrue())
			})

			It("filters in the negative case", func() {
				name := "vm-b60ce906-54d3-46d2-771c-7a9c8aad53bb (iless-pull-1611-pcf-network, cf-13d3260375a79e249520, cf-13d3260375a79e249520-router, iless-pull-1611-vms, p-bosh, p-bosh-cf-13d3260375a79e249520, p-bosh-cf-13d3260375a79e249520-router, router)"
				Expect(ResourceMatches(name, filter, true)).To(BeFalse())
			})
		})

		Context("given a terraform deployed vm resource", func() {
			It("filters in the positive case", func() {
				name := "pull-1611-infrastructure-subnet"
				Expect(ResourceMatches(name, filter, true)).To(BeTrue())
			})

			It("filters in the negative case", func() {
				name := "iless-pull-1611-infrastructure-subnet"
				Expect(ResourceMatches(name, filter, true)).To(BeFalse())
			})
		})
	})
})
