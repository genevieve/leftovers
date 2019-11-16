package main

import (
	"errors"
	"log"
	"os"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/aws"
	"github.com/genevieve/leftovers/azure"
	"github.com/genevieve/leftovers/gcp"
	"github.com/genevieve/leftovers/nsxt"
	"github.com/genevieve/leftovers/openstack"
	"github.com/genevieve/leftovers/vsphere"
	flags "github.com/jessevdk/go-flags"
)

type leftovers interface {
	Delete(filter string) error
	DeleteByType(filter, rType string) error
	List(filter string)
	ListByType(filter, rType string)
	Types()
}

var Version = "dev"

func main() {
	log.SetFlags(0)

	var o app.Options
	parser := flags.NewParser(&o, flags.HelpFlag|flags.PrintErrors)
	remaining, err := parser.ParseArgs(os.Args)
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}

	command := "destroy"
	if len(remaining) > 1 {
		command = remaining[1]
	}

	if o.Version {
		log.Printf("%s\n", Version)
		return
	}

	logger := app.NewLogger(os.Stdout, os.Stdin, o.NoConfirm, o.Debug)

	otherEnvVars := app.NewOtherEnvVars()
	otherEnvVars.LoadConfig(&o)

	var l leftovers

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
			log.Fatal("--no-confirm is not supported for vSphere.")
		}
		l, err = vsphere.NewLeftovers(logger, o.VSphereIP, o.VSphereUser, o.VSpherePassword, o.VSphereDC)
	case app.Openstack:
		l, err = openstack.NewLeftovers(logger, openstack.AuthArgs{
			AuthURL:    o.OpenstackAuthUrl,
			Username:   o.OpenstackUsername,
			Password:   o.OpenstackPassword,
			Domain:     o.OpenstackDomain,
			TenantName: o.OpenstackTenant,
			Region:     o.OpenstackRegion,
		})
	default:
		err = errors.New("Missing or unsupported BBL_IAAS.")
	}

	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}

	if command == "types" {
		l.Types()
		return
	}

	if o.DryRun {
		if o.Type == "" {
			l.List(o.Filter)
		} else {
			l.ListByType(o.Filter, o.Type)
		}
		return
	}

	if o.Type == "" {
		err = l.Delete(o.Filter)
	} else {
		err = l.DeleteByType(o.Filter, o.Type)
	}
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}
}
