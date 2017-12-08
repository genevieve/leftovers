package awsec2_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/genevievelesperance/leftovers/awsec2"
	"github.com/genevievelesperance/leftovers/awsec2/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Volumes", func() {
	var (
		ec2Client *fakes.EC2Client
		logger    *fakes.Logger

		volumes awsec2.Volumes
	)

	BeforeEach(func() {
		ec2Client = &fakes.EC2Client{}
		logger = &fakes.Logger{}

		volumes = awsec2.NewVolumes(ec2Client, logger)
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			logger.PromptCall.Returns.Proceed = true
			ec2Client.DescribeVolumesCall.Returns.Output = &ec2.DescribeVolumesOutput{
				Volumes: []*ec2.Volume{{
					VolumeId: aws.String("banana"),
					State:    aws.String("available"),
				}},
			}
		})

		It("deletes ec2 volumes", func() {
			err := volumes.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(ec2Client.DeleteVolumeCall.CallCount).To(Equal(1))
			Expect(ec2Client.DeleteVolumeCall.Receives.Input.VolumeId).To(Equal(aws.String("banana")))
			Expect(logger.PrintfCall.Messages).To(Equal([]string{"SUCCESS deleting volume banana\n"}))
		})

		Context("when the client fails to list volumes", func() {
			BeforeEach(func() {
				ec2Client.DescribeVolumesCall.Returns.Error = errors.New("some error")
			})

			It("does not try deleting them", func() {
				err := volumes.Delete()
				Expect(err.Error()).To(Equal("Describing volumes: some error"))

				Expect(ec2Client.DeleteVolumeCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete the volume", func() {
			BeforeEach(func() {
				ec2Client.DeleteVolumeCall.Returns.Error = errors.New("some error")
			})

			It("returns the error", func() {
				err := volumes.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PrintfCall.Messages).To(Equal([]string{"ERROR deleting volume banana: some error\n"}))
			})
		})

		Context("when the user responds no to the prompt", func() {
			BeforeEach(func() {
				logger.PromptCall.Returns.Proceed = false
			})

			It("returns the error", func() {
				err := volumes.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.Receives.Message).To(Equal("Are you sure you want to delete volume banana?"))
				Expect(ec2Client.DeleteVolumeCall.CallCount).To(Equal(0))
			})
		})

		Context("when the volume is not available", func() {
			BeforeEach(func() {
				ec2Client.DescribeVolumesCall.Returns.Output = &ec2.DescribeVolumesOutput{
					Volumes: []*ec2.Volume{{
						VolumeId: aws.String("banana"),
						State:    aws.String("nope"),
					}},
				}
			})
			It("does not prompt the user and it does not delete it", func() {
				err := volumes.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(logger.PromptCall.CallCount).To(Equal(0))
				Expect(ec2Client.DeleteVolumeCall.CallCount).To(Equal(0))
			})
		})
	})
})
