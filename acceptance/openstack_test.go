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
		iaas := os.Getenv(LEFTOVERS_ACCEPTANCE)
		if strings.ToLower(iaas) != "openstack" {
			Skip("Skipping Openstack acceptance tests.")
		}
	})

	Context("with incorrect openstack credentials", func() {
		Describe("NewLeftovers", func() {
			Context("given incorrect openstack credentials", func() {
				It("should return an empty Leftovers and write an error", func() {
					incorrectAuthArgs := openstack.AuthArgs{}

					leftovers, err := openstack.NewLeftovers(nil, incorrectAuthArgs)

					Expect(leftovers).To(Equal(openstack.Leftovers{}))
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(ContainSubstring("failed to make authenticated client"))
				})
			})
		})
	})

	Context("with correct openstack credentials and config", func() {
		BeforeEach(func() {

			acc = NewOpenStackAcceptance()
			err := acc.configureAuthClient()
			Expect(err).NotTo(HaveOccurred())

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

			color.NoColor = true
		})

		Describe("Dry run for Volumes", func() {
			var volumeID string

			BeforeEach(func() {
				volumeID = acc.CreateVolume("some name")
			})

			AfterEach(func() {
				err := acc.SafeDeleteVolume(volumeID)
				Expect(err).NotTo(HaveOccurred())
			})

			Context("a filter is supplied", func() {
				It("should log a warning", func() {
					leftovers.List("filter")

					Expect(stdout.String()).To(ContainSubstring("Warning: Filters are not supported for OpenStack."))
					Expect(acc.VolumeExists(volumeID)).To(BeTrue())
				})
			})

			It("lists resources without deleting", func() {
				leftovers.List("")

				Expect(acc.VolumeExists(volumeID)).To(BeTrue())
				Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s]", "some name", volumeID)))
			})
		})

		Describe("Types", func() {
			It("lists the resource types that can be deleted", func() {
				leftovers.Types()

				Expect(stdout.String()).To(ContainSubstring("Volume"))
			})
		})

		Describe("Delete", func() {
			var volumeID string

			BeforeEach(func() {
				volumeID = acc.CreateVolume("some name")
			})

			AfterEach(func() {
				if acc.VolumeExists(volumeID) {
					err := acc.SafeDeleteVolume(volumeID)
					Expect(err).ToNot(HaveOccurred())
				}
			})

			Context("a filter is supplied", func() {
				It("should panic with an error", func() {
					err := leftovers.Delete("filter")
					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("cannot delete openstack resources using a filter"))
					Expect(stdout.String()).To(ContainSubstring("Error: Filters are not supported for OpenStack. Aborting deletion!"))
					Expect(acc.VolumeExists(volumeID)).To(BeTrue())
				})
			})
			Context("given one volume associated with a project id", func() {
				BeforeEach(func() {
					Eventually(func() bool {
						isSafeToDelete, err := acc.IsSafeToDelete(volumeID)
						Expect(err).NotTo(HaveOccurred())
						return isSafeToDelete
					}, 100).Should(BeTrue())
				})

				It("deletes only that volume", func() {
					err := leftovers.Delete("")
					Expect(err).NotTo(HaveOccurred())
					Eventually(func() bool {
						return acc.VolumeExists(volumeID)
					}, 100).Should(BeFalse())
					Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleting...", "some name", volumeID)))
					Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleted!", "some name", volumeID)))
				})
			})
		})

		Describe("DeleteType for Volumes", func() {
			var volumeID string

			BeforeEach(func() {
				volumeID = acc.CreateVolume("some name")
				Eventually(func() bool {
					isSafeToDelete, err := acc.IsSafeToDelete(volumeID)
					Expect(err).NotTo(HaveOccurred())
					return isSafeToDelete
				}, 100).Should(BeTrue())
			})

			AfterEach(func() {
				if acc.VolumeExists(volumeID) {
					err := acc.SafeDeleteVolume(volumeID)
					Expect(err).ToNot(HaveOccurred())
				}
			})

			Context("when a filter is passed", func() {
				It("returns an error with a warning and doesnt delete anything", func() {
					err := leftovers.DeleteType("some filter", "Volume")
					Consistently(func() bool {
						return acc.VolumeExists(volumeID)
					}, "2s", "100ms").Should(BeTrue())

					Expect(err).To(HaveOccurred())
					Expect(err.Error()).To(Equal("cannot delete openstack resources using a filter"))
					Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("Error: Filters are not supported for OpenStack. Aborting deletion!")))
				})
			})

			Context("when a filter isn't passed", func() {
				// TODO: update with other resources
				It("deletes only volumes", func() {
					err := leftovers.DeleteType("", "Volume")
					Expect(err).NotTo(HaveOccurred())

					Eventually(func() bool {
						return acc.VolumeExists(volumeID)
					}, 100).Should(BeFalse())

					Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleting...", "some name", volumeID)))
					Expect(stdout.String()).To(ContainSubstring(fmt.Sprintf("[Volume: %s %s] Deleted!", "some name", volumeID)))
				})
			})
		})
	})
})
