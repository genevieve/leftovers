package gcp

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/common"
	"github.com/genevieve/leftovers/openstack/compute"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

type resource interface {
	List(filter string) ([]common.Deletable, error)
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
	resources    []resource
}

// NewLeftovers returns a new Leftovers for OpenStack that can be used to list resources,
// list types, or delete resources for the provided account. It returns an error
// if the credentials provided are invalid or if a client fails to be created.
func NewLeftovers(logger logger, authURL string) (Leftovers, error) {
	oc, err := openstack.NewClient(authURL)
	if err != nil {
		return Leftovers{}, err
	}

	service, err := openstack.NewComputeV2(oc, gophercloud.EndpointOpts{
		Region:       "",
		Availability: "",
	})

	client := compute.NewClient(service, logger)

	return Leftovers{
		logger:       logger,
		asyncDeleter: app.NewAsyncDeleter(logger),
		resources: []resource{
			compute.NewFloatingIPs(client, logger),
		},
	}, nil
}

// List will print all of the resources that match the provided filter.
func (l Leftovers) List(filter string) {
	l.logger.NoConfirm()

	var deletables []common.Deletable

	for _, r := range l.resources {
		list, err := r.List(filter)
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
// you to confirm deletion (if enabled), and delete those
// that are selected.
func (l Leftovers) Delete(filter string) error {
	deletables := [][]common.Deletable{}

	for _, r := range l.resources {
		list, err := r.List(filter)
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
	deletables := [][]common.Deletable{}

	for _, r := range l.resources {
		if r.Type() == rType {
			list, err := r.List(filter)
			if err != nil {
				l.logger.Println(color.YellowString(err.Error()))
			}

			deletables = append(deletables, list)
		}
	}

	return l.asyncDeleter.Run(deletables)
}
