package openstack

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/common"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

type listTyper interface {
	List() ([]common.Deletable, error)
	Type() string
}

type logger interface {
	Printf(message string, a ...interface{})
	Println(message string)
	PromptWithDetails(resourceType, resourceName string) bool
	NoConfirm()
}

type Leftovers struct {
	logger       logger
	asyncDeleter app.AsyncDeleter
	resources    []listTyper
}

type AuthArgs struct {
	AuthURL    string
	Username   string
	Password   string
	Domain     string
	Region     string
	TenantName string
}

// NewLeftovers returns a new Leftovers for OpenStack that can be used to list resources,
// list types, or delete resources for the provided account. It returns an error
// if the credentials provided are invalid or if a client fails to be created.
func NewLeftovers(logger logger, authArgs AuthArgs) (Leftovers, error) {
	provider, err := openstack.AuthenticatedClient(gophercloud.AuthOptions{
		IdentityEndpoint: authArgs.AuthURL,
		Username:         authArgs.Username,
		Password:         authArgs.Password,
		DomainName:       authArgs.Domain,
		TenantName:       authArgs.TenantName,
		AllowReauth:      true,
	})
	if err != nil {
		return Leftovers{}, fmt.Errorf("failed to make authenticated client: %s", err)
	}

	openstackOptions := gophercloud.EndpointOpts{
		Region:       authArgs.Region,
		Availability: gophercloud.AvailabilityPublic,
	}

	serviceBS, err := openstack.NewBlockStorageV3(provider, openstackOptions)
	if err != nil {
		return Leftovers{}, fmt.Errorf("failed to create volume block storage client: %s", err)
	}

	serviceComputeInstance, err := openstack.NewComputeV2(provider, openstackOptions)
	if err != nil {
		return Leftovers{}, fmt.Errorf("failed to create compute instance client: %s", err)
	}

	serviceImages, err := openstack.NewImageServiceV2(provider, openstackOptions)
	if err != nil {
		return Leftovers{}, fmt.Errorf("failed to create images client: %s", err)
	}

	return Leftovers{
		logger:       logger,
		asyncDeleter: app.NewAsyncDeleter(logger),
		resources: []listTyper{
			NewVolumes(NewVolumesBlockStorageClient(VolumesAPI{serviceClient: serviceBS}), logger),
			NewComputeInstances(NewComputeInstanceClient(ComputeAPI{serviceClient: serviceComputeInstance}), logger),
			NewImages(NewImagesClient(ImageAPI{serviceClient: serviceImages}), logger),
		}}, nil
}

// List will print all of the resources that match the provided filter.
func (l Leftovers) List(filter string) {
	l.logger.NoConfirm()

	if filter != "" {
		l.logger.Println(color.YellowString("Warning: Filters are not supported for OpenStack."))
		return
	}

	var deletables []common.Deletable

	for _, r := range l.resources {
		list, err := r.List()
		if err != nil {
			l.logger.Println(color.YellowString(err.Error()))
		}

		deletables = append(deletables, list...)
	}

	for _, d := range deletables {
		l.logger.Println(fmt.Sprintf("[%s: %s]", d.Type(), d.Name()))
	}
}

// Types will print all the resource types that can
// be deleted on this IaaS.
func (l Leftovers) Types() {
	l.logger.NoConfirm()

	for _, r := range l.resources {
		l.logger.Println(r.Type())
	}
}

// Delete will collect all resources that contain
// the provided filter in the resource's identifier, prompt
// you to confirm deletion (if enabled), and delete thoseu
// that are selected.
func (l Leftovers) Delete(filter string) error {
	if filter != "" {
		l.logger.Println(color.RedString("Error: Filters are not supported for OpenStack. Aborting deletion!"))
		return errors.New("cannot delete openstack resources using a filter")
	}

	deletables := [][]common.Deletable{}

	for _, r := range l.resources {
		list, err := r.List()
		if err != nil {
			l.logger.Println(color.YellowString(err.Error()))
		}

		deletables = append(deletables, list)
	}

	return l.asyncDeleter.Run(deletables)
}

// DeleteType will collect all resources of the provied type that contain
// the provided filter in the resource's identifier, prompt
// you to confirm deletion (if enabled), and delete those
// that are selected.
func (l Leftovers) DeleteType(filter, rType string) error {
	if filter != "" {
		l.logger.Println(color.RedString("Error: Filters are not supported for OpenStack. Aborting deletion!"))
		return errors.New("cannot delete openstack resources using a filter")
	}

	deletables := [][]common.Deletable{}

	for _, r := range l.resources {
		if r.Type() == rType {
			list, err := r.List()
			if err != nil {
				l.logger.Println(color.YellowString(err.Error()))
			}

			deletables = append(deletables, list)
		}
	}

	return l.asyncDeleter.Run(deletables)
}
