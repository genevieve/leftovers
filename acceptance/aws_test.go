package acceptance

import (
	"bytes"
	"os"
	"strings"

	"github.com/fatih/color"
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

		color.NoColor = true
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

			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: leftovers-dry-run]"))
			Expect(stdout.String()).NotTo(ContainSubstring("[EC2 Key Pair: leftovers-acceptance] Deleting..."))
			Expect(stdout.String()).NotTo(ContainSubstring("[EC2 Key Pair: leftovers-acceptance] Deleted!"))
		})
	})

	Describe("Types", func() {
		It("lists the resource types that can be deleted", func() {
			deleter.Types()

			Expect(stdout.String()).To(ContainSubstring("vpc"))
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

			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: leftovers-acceptance] Deleting..."))
			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: leftovers-acceptance] Deleted!"))
		})
	})

	Describe("DeleteType", func() {
		BeforeEach(func() {
			filter = "lftvrs-acceptance-delete-type"
			acc.CreateKeyPair(filter)
		})

		It("deletes the key pair resources with the filter", func() {
			err := deleter.DeleteType(filter, "ec2-key-pair")
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: lftvrs-acceptance-delete-type] Deleting..."))
			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: lftvrs-acceptance-delete-type] Deleted!"))
		})
	})
})
