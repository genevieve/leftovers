package openstack_test

import (
	"errors"

	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/openstack/fakes"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Volumes", func() {
	Describe("List", func() {
		var (
			subject           openstack.Volumes
			fakeLogger        *fakes.Logger
			fakeVolumesClient *fakes.VolumesClient
		)

		BeforeEach(func() {
			fakeLogger = &fakes.Logger{}
			fakeVolumesClient = &fakes.VolumesClient{}

			subject = openstack.NewVolumes(fakeVolumesClient, fakeLogger)
		})

		It("returns all the deletables", func() {
			fakeLogger.PromptWithDetailsCall.ReturnsForCall = append(
				fakeLogger.PromptWithDetailsCall.ReturnsForCall,
				fakes.LoggerPromptWithDetailsCallReturn{Bool: true},
				fakes.LoggerPromptWithDetailsCallReturn{Bool: true},
				fakes.LoggerPromptWithDetailsCallReturn{Bool: false},
			)

			volume := volumes.Volume{
				ID:   "some-ID",
				Name: "some-name",
			}
			otherVolume := volumes.Volume{
				ID:   "other-ID",
				Name: "other-name",
			}
			anotherVolume := volumes.Volume{
				ID:   "another-ID",
				Name: "another-name",
			}
			fakeVolumesClient.ListCall.Returns.Volumes = []volumes.Volume{
				volume,
				otherVolume,
				anotherVolume,
			}

			result, err := subject.List()
			Expect(err).NotTo(HaveOccurred())

			resultType := subject.Type()
			Expect(resultType).To(Equal("Volume"))

			Expect(result).To(HaveLen(2))
			Expect(result[0].Name()).To(Equal("some-name some-ID"))
			Expect(result[1].Name()).To(Equal("other-name other-ID"))
		})

		Context("when there are no volumes", func() {
			It("returns an empty list", func() {
				result, err := subject.List()
				Expect(err).NotTo(HaveOccurred())

				Expect(result).To(BeEmpty())
			})
		})

		Context("when an error occurs", func() {
			Context("when listing fails", func() {
				It("returns an error", func() {
					fakeVolumesClient.ListCall.Returns.Error = errors.New("error-description")

					result, err := subject.List()
					Expect(err).To(HaveOccurred())

					Expect(result).To(BeNil())
					Expect(err).To(MatchError("error-description"))
				})
			})
		})
	})
})
