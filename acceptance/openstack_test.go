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

	Describe("a deletable volume journey", func() {
		It("deletes volumes", func() {
			By("failing to create a new Leftovers when openstack can't authenticate")
			incorrectAuthArgs := openstack.AuthArgs{}
			var err error
			leftovers, err = openstack.NewLeftovers(nil, incorrectAuthArgs)

			Expect(leftovers).To(Equal(openstack.Leftovers{}))
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("failed to make authenticated client"))

			By("listing all resources when calling Types")
			noConfirm := true
			stdout = bytes.NewBuffer([]byte{})
			logger := app.NewLogger(stdout, os.Stdin, noConfirm)

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

			By("warning the user when a filter is passed to List")
			volumeID := acc.CreateVolume("some name")
			leftovers.List("filter")

			Expect(stdout.String()).To(ContainSubstring("Warning: Filters are not supported for OpenStack."))
			Expect(acc.VolumeExists(volumeID)).To(BeTrue())

			By("listing all resources when a filter isn't passed to List")
			leftovers.List("")

			Expect(acc.VolumeExists(volumeID)).To(BeTrue())
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s]", "some name", volumeID)))

			By("warning the user and aborting when a filter is passed to Delete")
			err = leftovers.Delete("filter")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("cannot delete openstack resources using a filter"))
			Expect(stdout.String()).To(ContainSubstring("Error: Filters are not supported for OpenStack. Aborting deletion!"))
			Expect(acc.VolumeExists(volumeID)).To(BeTrue())

			By("deleting all resources when a filter isn't passed to Delete")
			Eventually(func() bool {
				isSafeToDelete, err := acc.IsSafeToDelete(volumeID)
				Expect(err).NotTo(HaveOccurred())
				return isSafeToDelete
			}, "2s").Should(BeTrue(), "Volume status should have transitioned to a deletable status")

			err = leftovers.Delete("")
			Expect(err).NotTo(HaveOccurred())
			Eventually(func() bool {
				return acc.VolumeExists(volumeID)
			}, "2s").Should(BeFalse(), "Volume should have been deleted")
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleting...", "some name", volumeID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleted!", "some name", volumeID)))

			By("deleting based on the type")
			volumeID = acc.CreateVolume("some name")
			Eventually(func() (bool, error) {
				return acc.IsSafeToDelete(volumeID)
			}, "2s").Should(BeTrue(), "Volume status should have transitioned to a deletable status")

			By("returning an error with a warning when a filter is passed to DeleteType")
			err = leftovers.DeleteType("some filter", "Volume")
			Consistently(func() bool {
				return acc.VolumeExists(volumeID)
			}, "2s", "100ms").Should(BeTrue(), "Volume should not have been deleted")

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("cannot delete openstack resources using a filter"))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("Error: Filters are not supported for OpenStack. Aborting deletion!")))

			By("deleting the correct type when DeleteType is passed no filter and volume")
			err = leftovers.DeleteType("", "Volume")
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() bool {
				return acc.VolumeExists(volumeID)
			}, "2s").Should(BeFalse(), "Volume should have been deleted")

			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleting...", "some name", volumeID)))
			Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleted!", "some name", volumeID)))
		})
	})
})
