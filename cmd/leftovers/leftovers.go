package main

import (
	"errors"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/aws"
	"github.com/genevieve/leftovers/azure"
	"github.com/genevieve/leftovers/gcp"
	"github.com/genevieve/leftovers/nsxt"
	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/vsphere"
)

type leftovers interface {
	Delete(filter string) error
	DeleteByType(filter, rType string) error
	List(filter string)
	ListByType(filter, rType string)
	Types()
}

func GetLeftovers(logger *app.Logger, o app.Options) (leftovers, error) {
	var (
		l   leftovers
		err error
	)

	switch o.IAAS {
	case app.AWS:
		l, err = aws.NewLeftovers(logger, o.AWSAccessKeyID, o.AWSSecretAccessKey, o.AWSSessionToken, o.AWSRegion)
	case app.Azure:
		l, err = azure.NewLeftovers(logger, o.AzureClientID, o.AzureClientSecret, o.AzureSubscriptionID, o.AzureTenantID)
	case app.GCP:
		l, err = gcp.NewLeftovers(logger, o.GCPServiceAccountKey)
	case app.NSXT:
		l, err = nsxt.NewLeftovers(logger, o.NSXTManagerHost, o.NSXTUser, o.NSXTPassword)
	case app.VSphere:
		if o.NoConfirm {
			err = errors.New("--no-confirm is not supported for vSphere.")
		} else {
			l, err = vsphere.NewLeftovers(logger, o.VSphereIP, o.VSphereUser, o.VSpherePassword, o.VSphereDC)
		}
	case app.Openstack:
		l, err = openstack.NewLeftovers(logger, o.OpenstackAuthUrl, o.OpenstackUsername, o.OpenstackPassword, o.OpenstackDomain, o.OpenstackTenant, o.OpenstackRegion)
	default:
		err = errors.New("Missing or unsupported BBL_IAAS.")
	}

	if err != nil {
		return nil, err
	}
	return l, nil
}
