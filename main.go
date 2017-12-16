package main

import (
	"log"
	"os"

	"github.com/genevievelesperance/leftovers/app"
	"github.com/genevievelesperance/leftovers/aws"
	"github.com/genevievelesperance/leftovers/azure"
	"github.com/genevievelesperance/leftovers/gcp"
	flags "github.com/jessevdk/go-flags"
)

type opts struct {
	IAAS      string `short:"i"  long:"iaas"        default:"aws"  env:"LEFTOVERS_IAAS"  description:"The IAAS for clean up."  `
	NoConfirm bool   `short:"n"  long:"no-confirm"                                       description:"Destroy resources without prompting. This is dangerous, make good choices!"`

	AWSAccessKeyID     string `long:"aws-access-key-id"     env:"AWS_ACCESS_KEY_ID"     description:"AWS access key id."`
	AWSSecretAccessKey string `long:"aws-secret-access-key" env:"AWS_SECRET_ACCESS_KEY" description:"AWS secret access key."`
	AWSRegion          string `long:"aws-region"            env:"AWS_REGION"            description:"AWS region."`

	AzureClientID       string `long:"azure-client-id"        env:"AZURE_CLIENT_ID"        description:"Azure client id."`
	AzureClientSecret   string `long:"azure-client-secret"    env:"AZURE_CLIENT_SECRET"    description:"Azure client secret."`
	AzureTenantID       string `long:"azure-tenant-id"        env:"AZURE_TENANT_ID"        description:"Azure tenant id."`
	AzureSubscriptionID string `long:"azure-subscription-id"  env:"AZURE_SUBSCRIPTION_ID"  description:"Azure subscription id."`

	GCPServiceAccountKey string `long:"gcp-service-account-key"  env:"GCP_SERVICE_ACCOUNT_KEY"  description:"GCP service account key path."`
}

type deleter interface {
	Delete() error
}

func main() {
	log.SetFlags(0)

	var c opts
	parser := flags.NewParser(&c, flags.HelpFlag|flags.PrintErrors)
	_, err := parser.ParseArgs(os.Args)
	if err != nil {
		os.Exit(0)
	}

	logger := app.NewLogger(os.Stdout, os.Stdin, c.NoConfirm)

	var d deleter
	switch c.IAAS {
	case "aws":
		d = aws.NewDeleter(logger, c.AWSAccessKeyID, c.AWSSecretAccessKey, c.AWSRegion)
	case "azure":
		d = azure.NewDeleter(logger, c.AzureClientID, c.AzureClientSecret, c.AzureSubscriptionID, c.AzureTenantID)
	case "gcp":
		d = gcp.NewDeleter(logger, c.GCPServiceAccountKey)
	}

	if err := d.Delete(); err != nil {
		log.Fatalf("\n\n%s\n", err)
	}
}
