package app_test

import (
	"os"

	"github.com/genevieve/leftovers/app"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("OtherEnvVars", func() {
	var (
		options      *app.Options
		otherEnvVars app.OtherEnvVars
	)
	BeforeEach(func() {
		options = &app.Options{}
		otherEnvVars = app.NewOtherEnvVars()
	})

	Describe("LoadConfig", func() {
		Context("AWS", func() {
			BeforeEach(func() {
				options.IAAS = "aws"
				os.Setenv("AWS_ACCESS_KEY_ID", "banana")
				os.Setenv("AWS_SECRET_ACCESS_KEY", "kiwi")
				os.Setenv("AWS_SESSION_TOKEN", "pineapple")
				os.Setenv("AWS_DEFAULT_REGION", "plum")
			})

			AfterEach(func() {
				os.Unsetenv("AWS_ACCESS_KEY_ID")
				os.Unsetenv("AWS_SECRET_ACCESS_KEY")
				os.Unsetenv("AWS_SESSION_TOKEN")
				os.Unsetenv("AWS_DEFAULT_REGION")
			})

			It("checks for other possible env vars", func() {
				otherEnvVars.LoadConfig(options)

				Expect(options.AWSAccessKeyID).To(Equal("banana"))
				Expect(options.AWSSecretAccessKey).To(Equal("kiwi"))
				Expect(options.AWSSessionToken).To(Equal("pineapple"))
				Expect(options.AWSRegion).To(Equal("plum"))
			})
		})

		Context("Azure", func() {
			BeforeEach(func() {
				options.IAAS = "azure"
				os.Setenv("ARM_CLIENT_ID", "banana")
				os.Setenv("ARM_CLIENT_SECRET", "kiwi")
				os.Setenv("ARM_SUBSCRIPTION_ID", "pineapple")
				os.Setenv("ARM_TENANT_ID", "plum")
			})

			AfterEach(func() {
				os.Unsetenv("ARM_CLIENT_ID")
				os.Unsetenv("ARM_CLIENT_SECRET")
				os.Unsetenv("ARM_SUBSCRIPTION_ID")
				os.Unsetenv("ARM_TENANT_ID")
			})

			It("checks for other possible env vars", func() {
				otherEnvVars.LoadConfig(options)

				Expect(options.AzureClientID).To(Equal("banana"))
				Expect(options.AzureClientSecret).To(Equal("kiwi"))
				Expect(options.AzureSubscriptionID).To(Equal("pineapple"))
				Expect(options.AzureTenantID).To(Equal("plum"))
			})
		})

		Context("GCP", func() {
			BeforeEach(func() {
				options.IAAS = "gcp"
				os.Setenv("GOOGLE_CREDENTIALS", "banana")
			})

			AfterEach(func() {
				os.Unsetenv("GOOGLE_CREDENTIALS")
			})

			It("checks for other possible env vars", func() {
				otherEnvVars.LoadConfig(options)

				Expect(options.GCPServiceAccountKey).To(Equal("banana"))
			})
		})

		Context("NSXT", func() {
			BeforeEach(func() {
				options.IAAS = "nsxt"
				os.Setenv("NSXT_MANAGER_HOST", "banana")
				os.Setenv("NSXT_USERNAME", "kiwi")
				os.Setenv("NSXT_PASSWORD", "pineapple")
			})

			AfterEach(func() {
				os.Unsetenv("NSXT_MANAGER_HOST")
				os.Unsetenv("NSXT_USERNAME")
				os.Unsetenv("NSXT_PASSWORD")
			})

			It("checks for other possible env vars", func() {
				otherEnvVars.LoadConfig(options)

				Expect(options.NSXTManagerHost).To(Equal("banana"))
				Expect(options.NSXTUser).To(Equal("kiwi"))
				Expect(options.NSXTPassword).To(Equal("pineapple"))
			})
		})

		Context("vSphere", func() {
			BeforeEach(func() {
				options.IAAS = "vsphere"
				os.Setenv("VSPHERE_USER", "banana")
				os.Setenv("VSPHERE_PASSWORD", "kiwi")
				os.Setenv("VSPHERE_DATACENTER", "pineapple")
				os.Setenv("VSPHERE_IP", "pineapple")
			})

			AfterEach(func() {
				os.Unsetenv("VSPHERE_USER")
				os.Unsetenv("VSPHERE_PASSWORD")
				os.Unsetenv("VSPHERE_DATACENTER")
				os.Unsetenv("VSPHERE_IP")
			})

			It("checks for other possible env vars", func() {
				otherEnvVars.LoadConfig(options)

				Expect(options.VSphereUser).To(Equal("banana"))
				Expect(options.VSpherePassword).To(Equal("kiwi"))
				Expect(options.VSphereDC).To(Equal("pineapple"))
				Expect(options.VSphereIP).To(Equal("pineapple"))
			})
		})
	})
})
