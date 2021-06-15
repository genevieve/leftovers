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
})
