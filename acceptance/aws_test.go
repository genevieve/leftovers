package acceptance_test

import (
	"bytes"
	"os"

	"github.com/genevievelesperance/leftovers"
	"github.com/genevievelesperance/leftovers/acceptance"
	"github.com/genevievelesperance/leftovers/app"
	"github.com/genevievelesperance/leftovers/aws"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AWS", func() {
	var acc *acceptance.AWSAcceptance

	BeforeEach(func() {
		acc = acceptance.NewAWSAcceptance()

		if !acc.ReadyToTest() {
			Skip("Skipping acceptance tests.")
		}
	})

	Describe("Delete", func() {
		var (
			stdout  *bytes.Buffer
			logger  *app.Logger
			filter  string
			deleter leftovers.Deleter
		)

		BeforeEach(func() {
			noConfirm := true
			stdout = bytes.NewBuffer([]byte{})
			logger = app.NewLogger(stdout, os.Stdin, noConfirm)

			filter = "leftovers-acceptance"
			acc.CreateKeyPair(filter)

			var err error
			deleter, err = aws.NewDeleter(logger, acc.AccessKeyId, acc.SecretAccessKey, acc.Region)
			Expect(err).NotTo(HaveOccurred())
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout).To(ContainSubstring("SUCCESS deleting key pair"))
		})
	})
})
