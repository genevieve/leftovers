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
		acc       *OpenStackAcceptance
		stdout    *bytes.Buffer
		leftovers openstack.Leftovers
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
	})

	AfterEach(func() {
		err := acc.CleanUpTestResources()
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Deleting OpenStack Resources Journey", func() {
		It("deletes the appropriate OpenStack resources", func() {
			By("failing to create a new Leftovers when openstack can't authenticate")
			incorrectAuthArgs := openstack.AuthArgs{}
			var err error
			leftovers, err = openstack.NewLeftovers(nil, incorrectAuthArgs)

			Expect(leftovers).To(Equal(openstack.Leftovers{}))
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to make authenticated client"))

			By("listing all resources when calling Types")
			noConfirm := true
			debug := true
			stdout = bytes.NewBuffer([]byte{})
			logger := app.NewLogger(stdout, os.Stdin, noConfirm, debug)
			leftovers, err = openstack.NewLeftovers(logger, openstack.AuthArgs{
				AuthURL:    acc.AuthURL,
				Username:   acc.Username,
				Password:   acc.Password,
				Domain:     acc.Domain,
				Region:     acc.Region,
				TenantName: acc.TenantName,
			})
			Expect(err).NotTo(HaveOccurred())
			leftovers.Types()

			Expect(stdout.String()).To(ContainSubstring("Volume"))
			Expect(stdout.String()).To(ContainSubstring("Compute Instance"))
			Expect(stdout.String()).To(ContainSubstring("Image"))

			By("warning the user when a filter is passed to List")
			volumeID := acc.CreateVolume("some volume")
			instanceID := acc.CreateComputeInstance("some instance")
			imageID := acc.CreateImage("some image")
			leftovers.List("filter")

			Expect(stdout.String()).To(ContainSubstring("Warning: Filters are not supported for OpenStack."))
			Expect(acc.VolumeExists(volumeID)).To(BeTrue())
			Expect(acc.ComputeInstanceExists(instanceID)).To(BeTrue())
			Expect(acc.ImageExists(imageID)).To(BeTrue())

			By("listing all resources when a filter isn't passed to List")
			leftovers.List("")

			Expect(stdout.String()).To(ContainSubstring("Listing Volumes..."))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s]", "some volume", volumeID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Compute Instance: %s %s]", "some instance", instanceID)))
			Expect(acc.VolumeExists(volumeID)).To(BeTrue())
			Expect(acc.ComputeInstanceExists(instanceID)).To(BeTrue())
			Expect(acc.ImageExists(imageID)).To(BeTrue())

			By("passing a filter to DeleteByType")
			err = leftovers.DeleteByType("some filter", "Volume")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("cannot delete openstack resources using a filter"))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("Error: Filters are not supported for OpenStack. Aborting deletion!")))
			Consistently(func() bool {
				return acc.VolumeExists(volumeID) && acc.ComputeInstanceExists(instanceID) && acc.ImageExists(imageID)
			}, "2s", "100ms").Should(BeTrue(), "Resources should not have been deleted")

			By("deleting by type 'Volume'")
			Eventually(func() (bool, error) {
				return acc.IsSafeToDeleteVolume(volumeID)
			}, "10s").Should(BeTrue(), "Volume status should have transitioned to a deletable status")
			err = leftovers.DeleteByType("", "Volume")

			Expect(err).NotTo(HaveOccurred())
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleting...", "some volume", volumeID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleted!", "some volume", volumeID)))
			Eventually(func() bool {
				return acc.VolumeExists(volumeID)
			}, "2s", "100ms").Should(BeFalse(), "Volume should have been deleted")
			Consistently(func() bool {
				return acc.ComputeInstanceExists(instanceID) && acc.ImageExists(imageID)
			}, "2s", "100ms").Should(BeTrue(), "Compute Instance and image should not have been deleted")

			By("deleting by type 'Compute Instance'")
			volumeID = acc.CreateVolume("some other volume")
			err = leftovers.DeleteByType("", "Compute Instance")

			Expect(err).NotTo(HaveOccurred())
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Compute Instance: %s %s] Deleting...", "some instance", instanceID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Compute Instance: %s %s] Deleted!", "some instance", instanceID)))
			Consistently(func() bool {
				return acc.VolumeExists(volumeID) && acc.ImageExists(imageID)
			}, "2s", "100ms").Should(BeTrue(), "Volume and image should not have been deleted")
			Eventually(func() bool {
				return acc.ComputeInstanceExists(instanceID)
			}, "2s", "100ms").Should(BeFalse(), "Compute Instance should have been deleted")

			By("deleting by type 'Image'", func() {
				volumeID = acc.CreateVolume("yet another volume")
				instanceID = acc.CreateComputeInstance("yet another compute instance")
				err = leftovers.DeleteByType("", "Image")
				Expect(err).NotTo(HaveOccurred())

				Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Image: %s %s] Deleting...", "some image", imageID)))
				Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Image: %s %s] Deleted!", "some image", imageID)))

				Consistently(func() bool {
					return acc.VolumeExists(volumeID) && acc.ComputeInstanceExists(instanceID)
				}, "2s", "100ms").Should(BeTrue(), "Volume and compute instance should not have been deleted")

				Eventually(func() bool {
					return acc.ImageExists(imageID)
				}).Should(BeFalse(), "Image should have been deleted")
			})

			By("passing a filter to Delete")
			instanceID = acc.CreateComputeInstanceWithNetwork("some other instance", true)
			imageID = acc.CreateImage("some other image")
			acc.AttachVolumeToComputeInstance(volumeID, instanceID)
			err = leftovers.Delete("filter")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("cannot delete openstack resources using a filter"))
			Expect(stdout.String()).To(ContainSubstring("Error: Filters are not supported for OpenStack. Aborting deletion!"))
			Expect(acc.VolumeExists(volumeID)).To(BeTrue())
			Expect(acc.ComputeInstanceExists(instanceID)).To(BeTrue())
			Expect(acc.ImageExists(imageID)).To(BeTrue())

			By("deleting all resources when a filter isn't passed to Delete")
			// Eventually(func() (bool, error) {
			// 	return acc.IsSafeToDeleteVolume(volumeID)
			// }, "10s").Should(BeTrue(), "Volume status should have transitioned to a deletable status")
			err = leftovers.Delete("")
			Expect(err).NotTo(HaveOccurred())

			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Compute Instance: %s %s] Deleted!", "some other instance", instanceID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Compute Instance: %s %s] Deleting...", "some other instance", instanceID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleting...", "yet another volume", volumeID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleted!", "yet another volume", volumeID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Image: %s %s] Deleted!", "some other image", imageID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Image: %s %s] Deleting...", "some other image", imageID)))
			Eventually(func() bool {
				return !(acc.VolumeExists(volumeID) || acc.ComputeInstanceExists(instanceID) || acc.ImageExists(imageID))
			}, "2s", "100ms").Should(BeTrue(), "Resources should have been deleted")
		})
	})
})
