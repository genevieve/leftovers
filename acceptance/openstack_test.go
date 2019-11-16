package acceptance

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/openstack"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Openstack", func() {
	var (
		acc     OpenStackAcceptance
		stdout  *bytes.Buffer
		deleter openstack.Leftovers

		volumeName   string
		volumeID     string
		instanceName string
		instanceID   string
		imageName    string
		imageID      string
	)

	BeforeEach(func() {
		color.NoColor = true

		iaas := os.Getenv(LEFTOVERS_ACCEPTANCE)
		if strings.ToLower(iaas) != "openstack" {
			Skip("Skipping Openstack acceptance tests.")
		}
		acc = NewOpenStackAcceptance()
		err := acc.configureAuthClient()
		Expect(err).NotTo(HaveOccurred())

		noConfirm := true
		debug := true
		stdout = bytes.NewBuffer([]byte{})
		logger := app.NewLogger(stdout, os.Stdin, noConfirm, debug)

		deleter, err = openstack.NewLeftovers(logger, acc.AuthURL, acc.Username, acc.Password, acc.Domain, acc.TenantName, acc.Region)
		Expect(err).NotTo(HaveOccurred())
	})

	AfterEach(func() {
		Expect(acc.CleanUpTestResources()).NotTo(HaveOccurred())
	})

	Describe("NewLeftovers", func() {
		Context("when a user provides invalid credentials", func() {
			It("fails with a helpful error message", func() {
				_, err := openstack.NewLeftovers(nil, "", "", "", "", "", "")
				Expect(err).To(MatchError(ContainSubstring("failed to make authenticated client:")))
			})
		})
	})

	Describe("Types", func() {
		It("lists types of resources that can be deleted", func() {
			deleter.Types()

			Expect(stdout.String()).To(ContainSubstring("Volume"))
			Expect(stdout.String()).To(ContainSubstring("Compute Instance"))
			Expect(stdout.String()).To(ContainSubstring("Image"))
		})
	})

	Describe("List with Filter", func() {
		It("warns that the filter flag is not supported", func() {
			deleter.List("filter")

			Expect(stdout.String()).To(ContainSubstring("Warning: Filters are not supported for OpenStack."))
		})
	})

	Describe("List", func() {
		BeforeEach(func() {
			volumeName = "list-volume"
			volumeID = acc.CreateVolume(volumeName)
			Expect(acc.VolumeExists(volumeID)).To(BeTrue())
		})

		It("lists resources", func() {
			deleter.List("")
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s]", volumeName, volumeID)))
		})
	})

	Describe("Delete with Filter", func() {
		It("fails with a message that the filter flag is not supported", func() {
			err := deleter.Delete("filter")
			Expect(err).To(MatchError("cannot delete openstack resources using a filter"))
		})
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			imageName = "delete-image"
			imageID = acc.CreateImage(imageName)

			volumeName = "delete-volume"
			volumeID = acc.CreateVolume(volumeName)

			instanceName = "delete-instance"
			instanceID = acc.CreateComputeInstanceWithNetwork(instanceName, true)

			acc.AttachVolumeToComputeInstance(volumeID, instanceID)

			Expect(acc.VolumeExists(volumeID)).To(BeTrue())
			Expect(acc.ComputeInstanceExists(instanceID)).To(BeTrue())
			Expect(acc.ImageExists(imageID)).To(BeTrue())
		})

		It("deletes all resources", func() {
			err := deleter.Delete("")
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Compute Instance: %s %s] Deleted!", instanceName, instanceID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Compute Instance: %s %s] Deleting...", instanceName, instanceID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleting...", volumeName, volumeID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleted!", volumeName, volumeID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Image: %s %s] Deleted!", imageName, imageID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Image: %s %s] Deleting...", imageName, imageID)))

			Eventually(func() bool {
				return !(acc.VolumeExists(volumeID) || acc.ComputeInstanceExists(instanceID) || acc.ImageExists(imageID))
			}, "2s", "100ms").Should(BeTrue(), "Resources should have been deleted")
		})
	})

	Describe("DeleteByType with Filter", func() {
		It("fails with a message that the filter flag is not supported", func() {
			err := deleter.DeleteByType("filter", "Image")
			Expect(err).To(MatchError("cannot delete openstack resources using a filter"))

			Expect(stdout.String()).To(ContainSubstring("Error: Filters are not supported for OpenStack. Aborting deletion!"))
		})
	})

	Describe("DeleteByType", func() {
		BeforeEach(func() {
			imageName = "delete-type-image"
			imageID = acc.CreateImage(imageName)

			volumeName = "delete-type-volume"
			volumeID = acc.CreateVolume(volumeName)
		})

		It("deletes resources of a certain type", func() {
			err := deleter.DeleteByType("", "Image")
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Image: %s %s] Deleting...", imageName, imageID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Image: %s %s] Deleted!", imageName, imageID)))

			Consistently(func() bool {
				return acc.VolumeExists(volumeID)
			}, "2s", "100ms").Should(BeTrue(), "Volumes should not have been deleted")

			Eventually(func() bool {
				return acc.ImageExists(imageID)
			}).Should(BeFalse(), "Image should have been deleted")
		})
	})
})
