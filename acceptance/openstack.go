package acceptance

import (
	"fmt"
	"os"

	"github.com/genevieve/leftovers/app"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"

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

	client *gophercloud.ServiceClient
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
	if o.client != nil {
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

	blockStorage, err := openstack.NewBlockStorageV3(provider, gophercloud.EndpointOpts{
		Region:       o.Region,
		Availability: gophercloud.AvailabilityPublic,
	})
	if err != nil {
		return fmt.Errorf("some authentication error creating networking client: %s", err)
	}

	o.client = blockStorage
	return nil
}

func (o *OpenStackAcceptance) CreateVolume(name string) string {
	vol, err := volumes.Create(o.client, volumes.CreateOpts{
		Size: 1,
		Name: name,
	}).Extract()
	Expect(err).NotTo(HaveOccurred())

	return vol.ID
}

func (o *OpenStackAcceptance) GetVolume(volumeID string) (volumes.Volume, error) {
	volume, err := volumes.Get(o.client, volumeID).Extract()

	if err != nil {
		return volumes.Volume{}, err
	}

	return *volume, nil
}

func (o *OpenStackAcceptance) DeleteVolume(volumeID string) error {
	return volumes.Delete(o.client, volumeID, volumes.DeleteOpts{}).ExtractErr()
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

func (o OpenStackAcceptance) IsSafeToDelete(volumeID string) (bool, error) {
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
		isSafeToDelete, err := o.IsSafeToDelete(volumeID)
		Expect(err).NotTo(HaveOccurred())
		return isSafeToDelete
	}, "1s").Should(BeTrue())

	return o.DeleteVolume(volumeID)
}
