package acceptance

import (
	"bytes"
	"os"
	"strings"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/aws"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AWS", func() {
	var (
		acc AWSAcceptance

		stdout  *bytes.Buffer
		filter  string
		deleter aws.Leftovers
	)

	BeforeEach(func() {
		iaas := os.Getenv(LEFTOVERS_ACCEPTANCE)
		if strings.ToLower(iaas) != "aws" {
			Skip("Skipping AWS acceptance tests.")
		}

		acc = NewAWSAcceptance()

		noConfirm := true
		stdout = bytes.NewBuffer([]byte{})
		logger := app.NewLogger(stdout, os.Stdin, noConfirm)

		var err error
		deleter, err = aws.NewLeftovers(logger, acc.AccessKeyId, acc.SecretAccessKey, acc.Region)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Dry run", func() {
		BeforeEach(func() {
			filter = "leftovers-dry-run"
			acc.CreateKeyPair(filter)
		})

		AfterEach(func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists resources without deleting", func() {
			deleter.List(filter)

			Expect(stdout.String()).To(ContainSubstring("EC2 Key Pair: leftovers-dry-run"))
			Expect(stdout.String()).NotTo(ContainSubstring("Are you sure you want to delete"))
			Expect(stdout.String()).NotTo(ContainSubstring("Deleting leftovers-dry-run."))
			Expect(stdout.String()).NotTo(ContainSubstring("SUCCESS deleting leftovers-dry-run!"))
			Expect(stdout.String()).NotTo(ContainSubstring("FAILED deleting key pair leftovers-dry-run"))
		})
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			filter = "leftovers-acceptance"
			acc.CreateKeyPair(filter)
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("Deleting leftovers-acceptance."))
			Expect(stdout.String()).To(ContainSubstring("SUCCESS deleting leftovers-acceptance!"))
			Expect(stdout.String()).NotTo(ContainSubstring("FAILED deleting key pair leftovers-acceptance"))
		})
	})
})
