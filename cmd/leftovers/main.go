package main

import (
	"log"
	"os"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/aws"
	"github.com/genevieve/leftovers/azure"
	"github.com/genevieve/leftovers/gcp"
	"github.com/genevieve/leftovers/vsphere"
	flags "github.com/jessevdk/go-flags"
)

type opts struct {
	IAAS      string `short:"i"  long:"iaas"        env:"BBL_IAAS"  description:"The IAAS for clean up."  `
	NoConfirm bool   `short:"n"  long:"no-confirm"                  description:"Destroy resources without prompting. This is dangerous, make good choices!"`
	Filter    string `short:"f"  long:"filter"                      description:"Filtering resources by an environment name."`
	DryRun    bool   `short:"d"  long:"dry-run"                     description:"List all resources without deleting any."`

	AWSAccessKeyID         string `long:"aws-access-key-id"        env:"BBL_AWS_ACCESS_KEY_ID"        description:"AWS access key id."`
	AWSSecretAccessKey     string `long:"aws-secret-access-key"    env:"BBL_AWS_SECRET_ACCESS_KEY"    description:"AWS secret access key."`
	AWSRegion              string `long:"aws-region"               env:"BBL_AWS_REGION"               description:"AWS region."`
	AzureClientID          string `long:"azure-client-id"          env:"BBL_AZURE_CLIENT_ID"          description:"Azure client id."`
	AzureClientSecret      string `long:"azure-client-secret"      env:"BBL_AZURE_CLIENT_SECRET"      description:"Azure client secret."`
	AzureTenantID          string `long:"azure-tenant-id"          env:"BBL_AZURE_TENANT_ID"          description:"Azure tenant id."`
	AzureSubscriptionID    string `long:"azure-subscription-id"    env:"BBL_AZURE_SUBSCRIPTION_ID"    description:"Azure subscription id."`
	GCPServiceAccountKey   string `long:"gcp-service-account-key"  env:"BBL_GCP_SERVICE_ACCOUNT_KEY"  description:"GCP service account key path."`
	VSphereVCenterIP       string `long:"vsphere-vcenter-ip"       env:"BBL_VSPHERE_VCENTER_IP"       description:"vSphere vCenter IP address."`
	VSphereVCenterPassword string `long:"vsphere-vcenter-password" env:"BBL_VSPHERE_VCENTER_PASSWORD" description:"vSphere vCenter password."`
	VSphereVCenterUser     string `long:"vsphere-vcenter-user"     env:"BBL_VSPHERE_VCENTER_USER"     description:"vSphere vCenter username."`
	VSphereVCenterDC       string `long:"vsphere-vcenter-dc"       env:"BBL_VSPHERE_VCENTER_DC"       description:"vSphere vCenter datacenter."`
}

type leftovers interface {
	Delete(string) error
	List(string)
}

func main() {
	log.SetFlags(0)

	var c opts
	parser := flags.NewParser(&c, flags.HelpFlag|flags.PrintErrors)
	_, err := parser.ParseArgs(os.Args)
	if err != nil {
		os.Exit(0)
	}

	logger := app.NewLogger(os.Stdout, os.Stdin, c.NoConfirm || c.DryRun)

	var l leftovers

	switch c.IAAS {
	case "aws":
		l, err = aws.NewLeftovers(logger, c.AWSAccessKeyID, c.AWSSecretAccessKey, c.AWSRegion)
	case "azure":
		l, err = azure.NewLeftovers(logger, c.AzureClientID, c.AzureClientSecret, c.AzureSubscriptionID, c.AzureTenantID)
	case "gcp":
		l, err = gcp.NewLeftovers(logger, c.GCPServiceAccountKey)
	case "vsphere":
		if c.Filter == "" {
			log.Fatalf("--filter is required for vSphere.")
		}
		if c.NoConfirm {
			log.Fatalf("--no-confirm is not supported for vSphere.")
		}
		l, err = vsphere.NewLeftovers(logger, c.VSphereVCenterIP, c.VSphereVCenterUser, c.VSphereVCenterPassword, c.VSphereVCenterDC)
	default:
		log.Fatalf("Missing or unsupported BBL_IAAS.")
	}

	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}

	if c.DryRun {
		l.List(c.Filter)
		return
	}

	if err := l.Delete(c.Filter); err != nil {
		log.Fatalf("\n\n%s\n", err)
	}
}
