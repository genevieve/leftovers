package acceptance

import (
	"fmt"
	"os"
	"strings"

	"github.com/genevieve/leftovers/app"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/volumeattach"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/imagedata"
	"github.com/gophercloud/gophercloud/openstack/imageservice/v2/images"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"

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

	testResources         []deletable
	volumesClient         *gophercloud.ServiceClient
	computeInstanceClient *gophercloud.ServiceClient
	imagesClient          *gophercloud.ServiceClient
	networkClient         *gophercloud.ServiceClient
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

	return &OpenStackAcceptance{
		Logger:     app.NewLogger(os.Stdin, os.Stdout, true, false),
		AuthURL:    authUrl,
		Domain:     domain,
		Username:   username,
		Password:   password,
		Region:     region,
		TenantName: tenant,
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

	networkClient, err := openstack.NewNetworkV2(provider, endpointOpts)

	if err != nil {
		return fmt.Errorf("some authentication error creating network client: %s", err)
	}

	o.volumesClient = blockStorage
	o.computeInstanceClient = instanceClient
	o.imagesClient = imagesClient
	o.networkClient = networkClient

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

func (o *OpenStackAcceptance) CreateNetworkWithCIDR(name, CIDR string) string {
	t := true

	network, err := networks.Create(o.networkClient, networks.CreateOpts{
		Name:         name,
		AdminStateUp: &t,
	}).Extract()
	Expect(err).NotTo(HaveOccurred())
	networkID := network.ID

	o.addTestResource(func() error {
		return networks.Delete(o.networkClient, networkID).ExtractErr()
	})

	gateway := "192.168.0.1"
	subnet, err := subnets.Create(o.networkClient, subnets.CreateOpts{
		NetworkID:   networkID,
		Name:        "leftovers-test-network",
		Description: "this is only for testing",
		CIDR:        CIDR,
		IPVersion:   gophercloud.IPv4,
		GatewayIP:   &gateway,
	}).Extract()
	Expect(err).NotTo(HaveOccurred())

	o.addTestResource(func() error {
		return subnets.Delete(o.networkClient, subnet.ID).ExtractErr()
	})

	return networkID
}

func (o *OpenStackAcceptance) CreateComputeInstanceWithNetwork(name string, withNetwork bool) string {
	imageID := o.CreateImage("my awesome empty image")
	serverCreateOpts := servers.CreateOpts{
		Name:          name,
		FlavorName:    "m1.tiny",
		ImageRef:      imageID,
		ServiceClient: o.computeInstanceClient,
	}

	if withNetwork {
		networkID := o.CreateNetworkWithCIDR("leftovers-test-network", "192.168.0.0/16")
		serverCreateOpts.Networks = []servers.Network{
			{UUID: networkID},
		}
	}

	instance, err := servers.Create(o.computeInstanceClient, serverCreateOpts).Extract()
	Expect(err).NotTo(HaveOccurred())

	o.addTestResource(func() error {
		return servers.Delete(o.computeInstanceClient, instance.ID).ExtractErr()
	})

	return instance.ID
}

func (o *OpenStackAcceptance) CreateComputeInstance(name string) string {
	return o.CreateComputeInstanceWithNetwork(name, false)
}

func (o *OpenStackAcceptance) AttachVolumeToComputeInstance(volumeID string, computeID string) {
	waitTimeInSeconds := 60
	err := servers.WaitForStatus(o.computeInstanceClient, computeID, "ACTIVE", waitTimeInSeconds)
	Expect(err).NotTo(HaveOccurred())

	// time.Sleep(180 * time.Second)

	_, err = volumeattach.Create(o.computeInstanceClient, computeID, volumeattach.CreateOpts{
		Device:   "/dev/ice",
		VolumeID: volumeID,
	}).Extract()
	Expect(err).NotTo(HaveOccurred())
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
	res := imagedata.Upload(o.imagesClient, image.ID, strings.NewReader("rubbish"))
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
