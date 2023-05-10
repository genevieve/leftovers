package route53_test

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsroute53 "github.com/aws/aws-sdk-go/service/route53"
	"github.com/genevieve/leftovers/aws/route53"
	"github.com/genevieve/leftovers/aws/route53/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("RecordSets", func() {
	var (
		client         *fakes.RecordSetsClient
		hostedZoneId   *string
		hostedZoneName string

		recordSets route53.RecordSets
	)

	BeforeEach(func() {
		client = &fakes.RecordSetsClient{}
		hostedZoneId = aws.String("zone-id")
		hostedZoneName = "zone-name"

		recordSets = route53.NewRecordSets(client)
	})

	Describe("Get", func() {
		BeforeEach(func() {
			client.ListResourceRecordSetsCall.Returns.ListResourceRecordSetsOutput = &awsroute53.ListResourceRecordSetsOutput{
				ResourceRecordSets: []*awsroute53.ResourceRecordSet{{
					Name: aws.String("the-name"),
					Type: aws.String("something-else"),
				}},
				IsTruncated: aws.Bool(false),
			}
		})

		It("gets the record sets", func() {
			records, err := recordSets.Get(hostedZoneId)
			Expect(err).NotTo(HaveOccurred())

			Expect(records).To(HaveLen(1))

			Expect(client.ListResourceRecordSetsCall.CallCount).To(Equal(1))
			Expect(client.ListResourceRecordSetsCall.Receives.ListResourceRecordSetsInput.HostedZoneId).To(Equal(hostedZoneId))
		})

		Context("when there are pages of record sets", func() {
			var inputs []*awsroute53.ListResourceRecordSetsInput

			BeforeEach(func() {
				client.ListResourceRecordSetsCall.Stub = func(input *awsroute53.ListResourceRecordSetsInput) (*awsroute53.ListResourceRecordSetsOutput, error) {
					inputs = append(inputs, input)

					switch client.ListResourceRecordSetsCall.CallCount {
					case 1:
						return &awsroute53.ListResourceRecordSetsOutput{
							ResourceRecordSets: []*awsroute53.ResourceRecordSet{{
								Type: aws.String("something-else"),
							}},
							NextRecordName: aws.String("one-more-thing"),
							IsTruncated:    aws.Bool(true),
						}, nil
					case 2:
						return &awsroute53.ListResourceRecordSetsOutput{
							ResourceRecordSets: []*awsroute53.ResourceRecordSet{{
								Type: aws.String("one-more-thing"),
							}},
							IsTruncated: aws.Bool(false),
						}, nil
					default:
						return nil, nil
					}

				}
			})

			It("loops over the list request", func() {
				records, err := recordSets.Get(hostedZoneId)
				Expect(err).NotTo(HaveOccurred())

				Expect(records).To(HaveLen(2))

				Expect(client.ListResourceRecordSetsCall.CallCount).To(Equal(2))
				Expect(inputs[0].HostedZoneId).To(Equal(hostedZoneId))
				Expect(inputs[0].StartRecordName).To(BeNil())
				Expect(inputs[1].StartRecordName).To(Equal(aws.String("one-more-thing")))
			})
		})

		Context("when the client fails to list resource record sets", func() {
			BeforeEach(func() {
				client.ListResourceRecordSetsCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				_, err := recordSets.Get(hostedZoneId)
				Expect(err).To(MatchError("List Resource Record Sets: banana"))
			})
		})
	})

	Describe("Delete", func() {
		var records []*awsroute53.ResourceRecordSet

		BeforeEach(func() {
			records = []*awsroute53.ResourceRecordSet{{
				Name: aws.String(hostedZoneName),
				Type: aws.String("something-else"),
			}}
		})

		It("deletes the record sets", func() {
			err := recordSets.DeleteAll(hostedZoneId, hostedZoneName, records)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ChangeResourceRecordSetsCall.CallCount).To(Equal(1))
			Expect(client.ChangeResourceRecordSetsCall.Receives.ChangeResourceRecordSetsInput.HostedZoneId).To(Equal(hostedZoneId))
			Expect(client.ChangeResourceRecordSetsCall.Receives.ChangeResourceRecordSetsInput.ChangeBatch.Changes[0].Action).To(Equal(aws.String("DELETE")))
			Expect(client.ChangeResourceRecordSetsCall.Receives.ChangeResourceRecordSetsInput.ChangeBatch.Changes[0].ResourceRecordSet.Type).To(Equal(aws.String("something-else")))
		})

		Context("when the resource record set is of type NS", func() {
			BeforeEach(func() {
				records = []*awsroute53.ResourceRecordSet{{
					Name: aws.String(hostedZoneName),
					Type: aws.String("NS"),
				}}
			})

			It("does not try to delete it", func() {
				err := recordSets.DeleteAll(hostedZoneId, hostedZoneName, records)
				Expect(err).NotTo(HaveOccurred())

				Expect(client.ChangeResourceRecordSetsCall.CallCount).To(Equal(0))
			})
		})

		Context("when the resource record set is of type SOA", func() {
			BeforeEach(func() {
				records = []*awsroute53.ResourceRecordSet{{
					Name: aws.String(hostedZoneName),
					Type: aws.String("SOA"),
				}}
			})

			It("does not try to delete it", func() {
				err := recordSets.DeleteAll(hostedZoneId, hostedZoneName, records)
				Expect(err).NotTo(HaveOccurred())

				Expect(client.ChangeResourceRecordSetsCall.CallCount).To(Equal(0))
			})
		})

		Context("when the client fails to delete resource record sets", func() {
			BeforeEach(func() {
				client.ChangeResourceRecordSetsCall.Returns.Error = errors.New("banana")
			})

			It("returns the error", func() {
				err := recordSets.DeleteAll(hostedZoneId, hostedZoneName, records)
				Expect(err).To(MatchError("Delete Resource Record Sets: banana"))
			})
		})
	})

	Describe("DeleteWithFilter", func() {
		var (
			records []*awsroute53.ResourceRecordSet
			filter  string
		)

		BeforeEach(func() {
			records = []*awsroute53.ResourceRecordSet{
				{Name: aws.String(fmt.Sprintf("kiwi-%s", hostedZoneName)), Type: aws.String("A")},
				{Name: aws.String(fmt.Sprintf("banana-%s", hostedZoneName)), Type: aws.String("A")},
			}
			filter = "banana"
		})

		It("deletes the record sets that contain the filter", func() {
			err := recordSets.DeleteWithFilter(hostedZoneId, hostedZoneName, records, filter)
			Expect(err).NotTo(HaveOccurred())

			Expect(client.ChangeResourceRecordSetsCall.CallCount).To(Equal(1))
			Expect(client.ChangeResourceRecordSetsCall.Receives.ChangeResourceRecordSetsInput.HostedZoneId).To(Equal(hostedZoneId))
			Expect(client.ChangeResourceRecordSetsCall.Receives.ChangeResourceRecordSetsInput.ChangeBatch.Changes[0].Action).To(Equal(aws.String("DELETE")))
			Expect(client.ChangeResourceRecordSetsCall.Receives.ChangeResourceRecordSetsInput.ChangeBatch.Changes[0].ResourceRecordSet.Type).To(Equal(aws.String("A")))
		})

		Context("if the record set is NS type", func() {
			BeforeEach(func() {
				records = []*awsroute53.ResourceRecordSet{
					{Name: aws.String(fmt.Sprintf("kiwi-%s", hostedZoneName)), Type: aws.String("A")},
					{Name: aws.String(fmt.Sprintf("banana-%s", hostedZoneName)), Type: aws.String("NS")},
				}
				filter = "banana"
			})

			It("deletes the record sets that contain the filter", func() {
				err := recordSets.DeleteWithFilter(hostedZoneId, hostedZoneName, records, filter)
				Expect(err).NotTo(HaveOccurred())

				Expect(client.ChangeResourceRecordSetsCall.CallCount).To(Equal(1))
				Expect(client.ChangeResourceRecordSetsCall.Receives.ChangeResourceRecordSetsInput.HostedZoneId).To(Equal(hostedZoneId))
				Expect(client.ChangeResourceRecordSetsCall.Receives.ChangeResourceRecordSetsInput.ChangeBatch.Changes[0].Action).To(Equal(aws.String("DELETE")))
				Expect(client.ChangeResourceRecordSetsCall.Receives.ChangeResourceRecordSetsInput.ChangeBatch.Changes[0].ResourceRecordSet.Type).To(Equal(aws.String("NS")))
			})
		})

		Context("when the client fails to delete resource record sets", func() {
			BeforeEach(func() {
				client.ChangeResourceRecordSetsCall.Returns.Error = errors.New("ruhroh")
			})

			It("returns the error", func() {
				err := recordSets.DeleteWithFilter(hostedZoneId, hostedZoneName, records, filter)
				Expect(err).To(MatchError(fmt.Sprintf("Delete Resource Record Sets in Hosted Zone %s: ruhroh", hostedZoneName)))
			})
		})
	})
})
