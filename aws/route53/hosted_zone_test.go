package route53_test

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/genevieve/leftovers/aws/route53"
	"github.com/genevieve/leftovers/aws/route53/fakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HostedZone", func() {
	var (
		client     *fakes.HostedZonesClient
		recordSets *fakes.RecordSets
		id         *string
		name       *string
		filter     string

		hostedZone route53.HostedZone
	)

	BeforeEach(func() {
		client = &fakes.HostedZonesClient{}
		recordSets = &fakes.RecordSets{}
		id = aws.String("the-zone-id")
		name = aws.String("the-zone-name")
		filter = "zone"

		hostedZone = route53.NewHostedZone(client, id, name, recordSets, filter)
	})

	Describe("Delete", func() {
		It("deletes all record sets and deletes the hosted zone", func() {
			err := hostedZone.Delete()
			Expect(err).NotTo(HaveOccurred())

			Expect(recordSets.GetCall.CallCount).To(Equal(1))
			Expect(recordSets.GetCall.Receives.HostedZoneId).To(Equal(id))

			Expect(recordSets.DeleteAllCall.CallCount).To(Equal(1))
			Expect(recordSets.DeleteAllCall.Receives.HostedZoneId).To(Equal(id))

			Expect(recordSets.DeleteWithFilterCall.CallCount).To(Equal(0))

			Expect(client.DeleteHostedZoneCall.CallCount).To(Equal(1))
			Expect(client.DeleteHostedZoneCall.Receives.DeleteHostedZoneInput.Id).To(Equal(id))
		})

		Context("when the zone does not contain the filter", func() {
			BeforeEach(func() {
				filter = "banana"
				hostedZone = route53.NewHostedZone(client, id, name, recordSets, filter)
			})

			It("deletes only record sets in the zone that contain the filter", func() {
				err := hostedZone.Delete()
				Expect(err).NotTo(HaveOccurred())

				Expect(recordSets.GetCall.CallCount).To(Equal(1))
				Expect(recordSets.GetCall.Receives.HostedZoneId).To(Equal(id))

				Expect(recordSets.DeleteAllCall.CallCount).To(Equal(0))

				Expect(recordSets.DeleteWithFilterCall.CallCount).To(Equal(1))
				Expect(recordSets.DeleteWithFilterCall.Receives.HostedZoneId).To(Equal(id))
				Expect(recordSets.DeleteWithFilterCall.Receives.Filter).To(Equal("banana"))

				Expect(client.DeleteHostedZoneCall.CallCount).To(Equal(0))
			})
		})

		Context("when record sets fails to get", func() {
			BeforeEach(func() {
				recordSets.GetCall.Returns.Error = errors.New("ruhroh")
			})

			It("returns the error", func() {
				err := hostedZone.Delete()
				Expect(err).To(MatchError("Get Record Sets: ruhroh"))
			})
		})

		Context("when deleting all record sets fails", func() {
			BeforeEach(func() {
				recordSets.DeleteAllCall.Returns.Error = errors.New("ruhroh")
			})

			It("returns the error", func() {
				err := hostedZone.Delete()
				Expect(err).To(MatchError("Delete All Record Sets: ruhroh"))
			})
		})

		Context("when deleting record sets with filter fails", func() {
			BeforeEach(func() {
				filter = "banana"
				hostedZone = route53.NewHostedZone(client, id, name, recordSets, filter)
				recordSets.DeleteWithFilterCall.Returns.Error = errors.New("ruhroh")
			})

			It("returns the error", func() {
				err := hostedZone.Delete()
				Expect(err).To(MatchError("Delete Record Sets With Filter: ruhroh"))
			})
		})

		Context("when the client fails to delete the zone", func() {
			BeforeEach(func() {
				client.DeleteHostedZoneCall.Returns.Error = errors.New("ruhroh")
			})

			It("returns the error", func() {
				err := hostedZone.Delete()
				Expect(err).To(MatchError("Delete: ruhroh"))
			})
		})
	})

	Describe("Name", func() {
		It("returns the identifier", func() {
			Expect(hostedZone.Name()).To(Equal("the-zone-name"))
		})
	})

	Describe("Type", func() {
		It("returns the type", func() {
			Expect(hostedZone.Type()).To(Equal("Route53 Hosted Zone"))
		})
	})
})
