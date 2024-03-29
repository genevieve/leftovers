package acceptance

import (
	"fmt"
	"os"
	"strings"

	"github.com/genevieve/leftovers/app"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/db/v1/flavors"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/imagedata"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"

	. "github.com/onsi/gomega"
)

type OpenStackAcceptance struct {
	Logger     *app.Logger
	AuthURL    string
	Domain     string
	Username   string
	Password   string
	Region     string
	TenantName string

	volumesClient         *gophercloud.ServiceClient
	computeInstanceClient *gophercloud.ServiceClient
	imagesClient          *gophercloud.ServiceClient
}

func NewOpenStackAcceptance() OpenStackAcceptance {
	authUrl := os.Getenv("BBL_OPENSTACK_AUTH_URL")
	Expect(authUrl).NotTo(BeEmpty(), "Missing $BBL_OPENSTACK_AUTH_URL.")

	domain := os.Getenv("BBL_OPENSTACK_DOMAIN")
	Expect(domain).NotTo(BeEmpty(), "Missing $BBL_OPENSTACK_DOMAIN.")

	username := os.Getenv("BBL_OPENSTACK_USERNAME")
	Expect(username).NotTo(BeEmpty(), "Missing $BBL_OPENSTACK_USERNAME.")

	password := os.Getenv("BBL_OPENSTACK_PASSWORD")
	Expect(password).NotTo(BeEmpty(), "Missing $BBL_OPENSTACK_PASSWORD.")

	region := os.Getenv("BBL_OPENSTACK_REGION")
	Expect(region).NotTo(BeEmpty(), "Missing $BBL_OPENSTACK_REGION.")

	tenant := os.Getenv("BBL_OPENSTACK_PROJECT")
	Expect(tenant).NotTo(BeEmpty(), "Missing $BBL_OPENSTACK_PROJECT.")

	provider, err := openstack.AuthenticatedClient(gophercloud.AuthOptions{
		IdentityEndpoint: authUrl,
		Username:         username,
		Password:         password,
		DomainName:       domain,
		TenantName:       tenant,
		AllowReauth:      true,
	})
	Expect(err).NotTo(HaveOccurred())

	endpointOpts := gophercloud.EndpointOpts{
		Region:       region,
		Availability: gophercloud.AvailabilityPublic,
	}

	blockStorage, err := openstack.NewBlockStorageV3(provider, endpointOpts)
	Expect(err).NotTo(HaveOccurred())

	instanceClient, err := openstack.NewComputeV2(provider, endpointOpts)
	Expect(err).NotTo(HaveOccurred())

	imagesClient, err := openstack.NewImageServiceV2(provider, endpointOpts)
	Expect(err).NotTo(HaveOccurred())

	return OpenStackAcceptance{
		Logger:     app.NewLogger(os.Stdin, os.Stdout, true, false),
		AuthURL:    authUrl,
		Domain:     domain,
		Username:   username,
		Password:   password,
		Region:     region,
		TenantName: tenant,

		volumesClient:         blockStorage,
		computeInstanceClient: instanceClient,
		imagesClient:          imagesClient,
	}
}

func (o OpenStackAcceptance) CreateVolume(name string) string {
	vol, err := volumes.Create(o.volumesClient, volumes.CreateOpts{
		Size: 1,
		Name: name,
	}).Extract()
	Expect(err).NotTo(HaveOccurred())

	return vol.ID
}

func (o OpenStackAcceptance) GetVolume(volumeID string) (volumes.Volume, error) {
	volume, err := volumes.Get(o.volumesClient, volumeID).Extract()
	if err != nil {
		return volumes.Volume{}, err
	}

	return *volume, nil
}

func (o OpenStackAcceptance) VolumeExists(volumeID string) bool {
	_, err := o.GetVolume(volumeID)
	if err != nil {
		_, ok := err.(gophercloud.ErrDefault404)
		Expect(ok).To(BeTrue(), fmt.Sprintf("Unexpected error: %s", err))
		return false
	}

	return true
}

func (o OpenStackAcceptance) DeleteVolume(volumeID string) {
	volumes.Delete(o.volumesClient, volumeID, volumes.DeleteOpts{})
}

func (o *OpenStackAcceptance) CreateComputeInstance(name string) string {
	imageID := o.CreateImage(fmt.Sprintf("empty-image-%s", name))

	all, err := flavors.List(o.computeInstanceClient).AllPages()
	Expect(err).NotTo(HaveOccurred())

	f, err := flavors.ExtractFlavors(all)
	Expect(err).NotTo(HaveOccurred())

	var flavorID string
	for _, flavor := range f {
		if flavor.Name == "m1.tiny" {
			flavorID = flavor.StrID
			break
		}
	}
	serverCreateOpts := servers.CreateOpts{
		Name:      name,
		FlavorRef: flavorID,
		ImageRef:  imageID,
	}

	instance, err := servers.Create(o.computeInstanceClient, serverCreateOpts).Extract()
	Expect(err).NotTo(HaveOccurred())

	return instance.ID
}

func (o OpenStackAcceptance) DeleteInstance(instanceID string) {
	servers.Delete(o.computeInstanceClient, instanceID)
}

func (o OpenStackAcceptance) ComputeInstanceExists(instanceID string) bool {
	_, err := servers.Get(o.computeInstanceClient, instanceID).Extract()
	if err != nil {
		_, ok := err.(gophercloud.ErrDefault404)
		Expect(ok).To(BeTrue(), fmt.Sprintf("Unexpected error: %s", err))
		return false
	}
	return true
}

func (o *OpenStackAcceptance) CreateImage(name string) string {
	image, err := images.Create(o.imagesClient, images.CreateOpts{
		Name:            name,
		DiskFormat:      "iso",
		ContainerFormat: "ami",
	}).Extract()
	Expect(err).NotTo(HaveOccurred())

	res := imagedata.Upload(o.imagesClient, image.ID, strings.NewReader("rubbish"))
	Expect(res.Err).NotTo(HaveOccurred())

	return image.ID
}

func (o OpenStackAcceptance) GetImage(imageID string) (images.Image, error) {
	image, err := images.Get(o.imagesClient, imageID).Extract()
	if err != nil {
		return images.Image{}, err
	}

	return *image, nil
}

func (o OpenStackAcceptance) ImageExists(imageID string) bool {
	_, err := o.GetImage(imageID)
	if err != nil {
		_, ok := err.(gophercloud.ErrDefault404)
		Expect(ok).To(BeTrue(), fmt.Sprintf("Unexpected error: %s", err))
		return false
	}

	return true
}

func (o OpenStackAcceptance) DeleteImage(imageID string) {
	images.Delete(o.imagesClient, imageID)
}
