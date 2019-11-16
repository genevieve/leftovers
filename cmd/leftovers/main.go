package main

import (
	"errors"
	"log"
	"os"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/aws"
	"github.com/genevieve/leftovers/azure"
	"github.com/genevieve/leftovers/commands"
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

type command interface {
	Execute(app.Options) error
}

var Version = "dev"

func main() {
	log.SetFlags(0)

	var o app.Options
	parser := flags.NewParser(&o, flags.HelpFlag|flags.PrintErrors)
	remaining, err := parser.ParseArgs(os.Args)
	if err != nil {
		return
	}

	if o.Version {
		log.Printf("%s\n", Version)
		return
	}

	cmd := "delete"
	if len(remaining) > 1 {
		cmd = "types"
	}
	if o.DryRun {
		cmd = "list"
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
		l, err = openstack.NewLeftovers(logger, o.OpenstackAuthUrl, o.OpenstackUsername, o.OpenstackPassword, o.OpenstackDomain, o.OpenstackTenant, o.OpenstackRegion)
	default:
		err = errors.New("Missing or unsupported BBL_IAAS.")
	}

	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}

	commandSet := map[string]command{}
	commandSet["delete"] = commands.NewDelete(l)
	commandSet["list"] = commands.NewList(l)
	commandSet["types"] = commands.NewTypes(l)

	err = commandSet[cmd].Execute(o)
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}
}
