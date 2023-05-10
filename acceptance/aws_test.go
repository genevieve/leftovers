package acceptance

import (
	"bytes"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/aws"

	. "github.com/onsi/ginkgo/v2"
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
		debug := false
		stdout = bytes.NewBuffer([]byte{})
		logger := app.NewLogger(stdout, os.Stdin, noConfirm, debug)

		var err error
		deleter, err = aws.NewLeftovers(logger, acc.AccessKeyId, acc.SecretAccessKey, acc.SessionToken, acc.Region)
		Expect(err).NotTo(HaveOccurred())

		color.NoColor = true
	})

	Describe("List", func() {
		BeforeEach(func() {
			filter = "leftovers-acc-list-all"
			acc.CreateKeyPair(filter)
		})

		AfterEach(func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists resources without deleting", func() {
			deleter.List(filter)

			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s]", filter))
			Expect(stdout.String()).NotTo(ContainSubstring("[EC2 Key Pair: %s] Deleting...", filter))
			Expect(stdout.String()).NotTo(ContainSubstring("[EC2 Key Pair: %s] Deleted!", filter))
		})
	})

	Describe("ListByType", func() {
		BeforeEach(func() {
			filter = "leftovers-acc-list-by-type"
			acc.CreateKeyPair(filter)
		})

		AfterEach(func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists resources of the specified type without deleting", func() {
			deleter.List(filter)

			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s]", filter))
			Expect(stdout.String()).NotTo(ContainSubstring("[EC2 Key Pair: %s] Deleting...", filter))
			Expect(stdout.String()).NotTo(ContainSubstring("[EC2 Key Pair: %s] Deleted!", filter))
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
			filter = "leftovers-acc-delete-all"
			acc.CreateKeyPair(filter)
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s] Deleting...", filter))
			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s] Deleted!", filter))
		})
	})

	Describe("DeleteByType", func() {
		BeforeEach(func() {
			filter = "leftovers-acc-delete-type"
			acc.CreateKeyPair(filter)
		})

		It("deletes the key pair resources with the filter", func() {
			err := deleter.DeleteByType(filter, "ec2-key-pair")
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s] Deleting...", filter))
			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s] Deleted!", filter))
		})
	})
})
