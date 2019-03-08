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
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/imagedata"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"

	. "github.com/onsi/gomega"
)

type OpenStackAcceptance struct {
	Logger      *app.Logger
	AuthURL     string
	Domain      string
	Username    string
	Password    string
	NetworkName string
	Region      string
	TenantName  string

	testResources         []deletable
	volumesClient         *gophercloud.ServiceClient
	computeInstanceClient *gophercloud.ServiceClient
	imagesClient          *gophercloud.ServiceClient
}

type deletable interface {
	Delete() error
}

type testResource struct {
	deleteFunction func() error
}

func (t testResource) Delete() error {
	return t.deleteFunction()
}

func NewOpenStackAcceptance() *OpenStackAcceptance {
	return &OpenStackAcceptance{
		Logger:      app.NewLogger(os.Stdin, os.Stdout, true),
		AuthURL:     os.Getenv("BBL_OPENSTACK_AUTH_URL"),
		Domain:      os.Getenv("BBL_OPENSTACK_DOMAIN"),
		Username:    os.Getenv("BBL_OPENSTACK_USERNAME"),
		Password:    os.Getenv("BBL_OPENSTACK_PASSWORD"),
		NetworkName: os.Getenv("BBL_OPENSTACK_NETWORK_NAME"),
		Region:      os.Getenv("BBL_OPENSTACK_REGION"),
		TenantName:  os.Getenv("BBL_OPENSTACK_PROJECT"),
	}
}

func (o *OpenStackAcceptance) configureAuthClient() error {
	if o.volumesClient != nil {
		return nil
	}

	provider, err := openstack.AuthenticatedClient(gophercloud.AuthOptions{
		IdentityEndpoint: o.AuthURL,
		Username:         o.Username,
		Password:         o.Password,
		DomainName:       o.Domain,
		TenantName:       o.TenantName,
		AllowReauth:      true,
	})
	if err != nil {
		return fmt.Errorf("could not create authenticated client provider: %s", err)
	}

	endpointOpts := gophercloud.EndpointOpts{
		Region:       o.Region,
		Availability: gophercloud.AvailabilityPublic,
	}

	blockStorage, err := openstack.NewBlockStorageV3(provider, endpointOpts)

	if err != nil {
		return fmt.Errorf("some authentication error creating networking client: %s", err)
	}

	instanceClient, err := openstack.NewComputeV2(provider, endpointOpts)

	if err != nil {
		return fmt.Errorf("some authentication error creating compute instance client: %s", err)
	}

	imagesClient, err := openstack.NewImageServiceV2(provider, endpointOpts)

	if err != nil {
		return fmt.Errorf("some authentication error creating images client: %s", err)
	}

	o.volumesClient = blockStorage
	o.computeInstanceClient = instanceClient
	o.imagesClient = imagesClient
	return nil
}

func (o *OpenStackAcceptance) CreateVolume(name string) string {
	vol, err := volumes.Create(o.volumesClient, volumes.CreateOpts{
		Size: 1,
		Name: name,
	}).Extract()
	Expect(err).NotTo(HaveOccurred())

	o.addTestResource(func() error {
		return volumes.Delete(o.volumesClient, vol.ID, volumes.DeleteOpts{}).ExtractErr()
	})

	return vol.ID
}

func (o *OpenStackAcceptance) CreateComputeInstance(name string) string {
	imageID := o.CreateImage("my awesome empty image")
	instance, err := servers.Create(o.computeInstanceClient, servers.CreateOpts{
		Name:          name,
		FlavorName:    "m1.tiny",
		ImageRef:      imageID,
		ServiceClient: o.computeInstanceClient,
	}).Extract()
	Expect(err).NotTo(HaveOccurred())

	o.addTestResource(func() error {
		return servers.Delete(o.computeInstanceClient, instance.ID).ExtractErr()
	})

	return instance.ID
}

func (o *OpenStackAcceptance) addTestResource(deleteFunc func() error) {
	tr := testResource{
		deleteFunction: deleteFunc,
	}
	o.testResources = append(o.testResources, tr)
}

func (o *OpenStackAcceptance) CreateImage(name string) string {
	image, err := images.Create(o.imagesClient, images.CreateOpts{
		Name:            name,
		DiskFormat:      "iso",
		ContainerFormat: "ami",
	}).Extract()
	Expect(err).NotTo(HaveOccurred())

	o.addTestResource(func() error {
		return images.Delete(o.imagesClient, image.ID).ExtractErr()
	})

	// upload iso for awesome empty image
	res := imagedata.Upload(o.imagesClient, image.ID, strings.NewReader(""))
	Expect(res.Err).NotTo(HaveOccurred())

	return image.ID
}

func (o *OpenStackAcceptance) GetVolume(volumeID string) (volumes.Volume, error) {
	volume, err := volumes.Get(o.volumesClient, volumeID).Extract()

	if err != nil {
		return volumes.Volume{}, err
	}

	return *volume, nil
}

func (o *OpenStackAcceptance) GetImage(imageID string) (images.Image, error) {
	image, err := images.Get(o.imagesClient, imageID).Extract()

	if err != nil {
		return images.Image{}, err
	}

	return *image, nil
}

func (o *OpenStackAcceptance) DeleteVolume(volumeID string) error {
	return volumes.Delete(o.volumesClient, volumeID, volumes.DeleteOpts{}).ExtractErr()
}

func (o *OpenStackAcceptance) VolumeExists(volumeID string) bool {
	_, err := o.GetVolume(volumeID)

	if err != nil {
		_, ok := err.(gophercloud.ErrDefault404)
		Expect(ok).To(BeTrue(), fmt.Sprintf("Unexpected error: %s", err.Error()))
		return false
	}

	return true
}

func (o *OpenStackAcceptance) ImageExists(imageID string) bool {
	_, err := o.GetImage(imageID)

	if err != nil {
		_, ok := err.(gophercloud.ErrDefault404)
		Expect(ok).To(BeTrue(), fmt.Sprintf("Unexpected error: %s", err.Error()))
		return false
	}

	return true

}

func (o *OpenStackAcceptance) ComputeInstanceExists(instanceID string) bool {
	_, err := servers.Get(o.computeInstanceClient, instanceID).Extract()
	if err != nil {
		_, ok := err.(gophercloud.ErrDefault404)
		Expect(ok).To(BeTrue(), fmt.Sprintf("Unexpected error: %s", err.Error()))
		return false
	}
	return true
}

func (o OpenStackAcceptance) IsSafeToDeleteVolume(volumeID string) (bool, error) {
	volume, err := o.GetVolume(volumeID)
	if err != nil {
		return false, err
	}

	status := volume.Status
	if status == "available" || status == "error" || status == "error_restoring" || status == "error_extending" || status == "error_managing" {
		return true, nil
	}
	return false, nil
}

func (o *OpenStackAcceptance) SafeDeleteVolume(volumeID string) error {
	Eventually(func() bool {
		isSafeToDelete, err := o.IsSafeToDeleteVolume(volumeID)
		Expect(err).NotTo(HaveOccurred())
		return isSafeToDelete
	}, "1s").Should(BeTrue())

	return o.DeleteVolume(volumeID)
}

func (o *OpenStackAcceptance) CleanUpTestResources() error {
	for _, resource := range o.testResources {
		err := resource.Delete()
		if err != nil {
			_, ok := err.(gophercloud.ErrDefault404)
			if !ok {
				return err
			}
		}
	}

	return nil
}
