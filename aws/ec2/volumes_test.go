package ec2_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/genevievelesperance/leftovers/aws/ec2"
	"github.com/genevievelesperance/leftovers/aws/ec2/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Volumes", func() {
	var (
		client *fakes.VolumesClient
		logger *fakes.Logger

		volumes ec2.Volumes
	)

	BeforeEach(func() {
		client = &fakes.VolumesClient{}
		logger = &fakes.Logger{}

		volumes = ec2.NewVolumes(client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			client.DescribeVolumesCall.Returns.Output = &awsec2.DescribeVolumesOutput{
				Volumes: []*awsec2.Volume{{
					VolumeId: aws.String("banana"),
					State:    aws.String("available"),
				}},
			}
		})

		It("deletes ec2 volumes", func() {
			err := volumes.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(client.DeleteVolumeCall.CallCount).To(Equal(1))
			Expect(client.DeleteVolumeCall.Receives.Input.VolumeId).To(Equal(aws.String("banana")))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting volume banana\n"}))
		})

		Context("when the client fails to list volumes", func() {
			BeforeEach(func() {
				client.DescribeVolumesCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := volumes.Delete()
				Expect(err).To(MatchError("Describing volumes: some error"))

				Expect(client.DeleteVolumeCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the volume", func() {
			BeforeEach(func() {
				client.DeleteVolumeCall.Returns.Error = errors.New("some error")
			})

			It("logs the error", func() {
				err := volumes.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting volume banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("does not delete the volume", func() {
				err := volumes.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete volume banana?"))
				Expect(client.DeleteVolumeCall.CallCount).To(Equal(0))
			})
		})

		Context("when the volume is not available", func() {
			BeforeEach(func() {
				client.DescribeVolumesCall.Returns.Output = &awsec2.DescribeVolumesOutput{
					Volumes: []*awsec2.Volume{{
						VolumeId: aws.String("banana"),
						State:    aws.String("nope"),
					}},
				}
			})

			It("does not prompt the user and it does not delete it", func() {
				err := volumes.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(client.DeleteVolumeCall.CallCount).To(Equal(0))
			})
		})
	})
})
