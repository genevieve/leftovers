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
		regexFilter string
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
			regexFilter = "leftovers-acc-lis[t]{1}-a[l]{2}$"
			acc.CreateKeyPair(filter)
		})

		AfterEach(func() {
			err := deleter.Delete(filter, false)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists resources without deleting", func() {
			deleter.List(filter, false)

			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s]", filter))
			Expect(stdout.String()).NotTo(ContainSubstring("[EC2 Key Pair: %s] Deleting...", filter))
			Expect(stdout.String()).NotTo(ContainSubstring("[EC2 Key Pair: %s] Deleted!", filter))
		})
	})

	Describe("ListByType", func() {
		BeforeEach(func() {
			filter = "leftovers-acc-list-by-type"
			regexFilter = "leftovers-acc-lis[t]{1}-b[y]{1}-type$"
			acc.CreateKeyPair(filter)
		})

		AfterEach(func() {
			err := deleter.Delete(filter, false)
			Expect(err).NotTo(HaveOccurred())
		})

		It("lists resources of the specified type without deleting", func() {
			deleter.List(filter, false)

			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s]", filter))
			Expect(stdout.String()).NotTo(ContainSubstring("[EC2 Key Pair: %s] Deleting...", filter))
			Expect(stdout.String()).NotTo(ContainSubstring("[EC2 Key Pair: %s] Deleted!", filter))
		})

		It("lists resources of the specified type without deleting with the regex filter", func() {
			deleter.List(filter, true)

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
			regexFilter = "leftovers-acc-dele[t]{1}e-a[l]{2}$"
			acc.CreateKeyPair(filter)
		})

		It("deletes resources with the filter", func() {
			err := deleter.Delete(filter, false)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s] Deleting...", filter))
			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s] Deleted!", filter))
		})


		It("deletes resources with the regex filter", func() {
			err := deleter.Delete(regexFilter, true)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s] Deleting...", filter))
			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s] Deleted!", filter))
		})
	})

	Describe("DeleteByType", func() {
		BeforeEach(func() {
			filter = "leftovers-acc-delete-type"
			regexFilter = "leftovers-acc-dele[t]{1}e-type"
			acc.CreateKeyPair(filter)
		})

		It("deletes the key pair resources with the filter", func() {
			err := deleter.DeleteByType(filter, "ec2-key-pair", false)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s] Deleting...", filter))
			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s] Deleted!", filter))
		})

		It("deletes resources with the regex filter", func() {
			err := deleter.DeleteByType(regexFilter, "ec2-key-pair", true)
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s] Deleting...", filter))
			Expect(stdout.String()).To(ContainSubstring("[EC2 Key Pair: %s] Deleted!", filter))
		})
	})
})
