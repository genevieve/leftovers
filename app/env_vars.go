package app

import "os"

type Options struct {
	Version bool `short:"v"  long:"version"                     description:"Print version."`

	IAAS      string `short:"i"  long:"iaas"        env:"BBL_IAAS"  description:"The IaaS for clean up."  `
	NoConfirm bool   `short:"n"  long:"no-confirm"                  description:"Destroy resources without prompting. This is dangerous, make good choices!"`
	DryRun    bool   `short:"d"  long:"dry-run"                     description:"List all resources without deleting any."`
	Filter    string `short:"f"  long:"filter"                      description:"Filtering resources by an environment name."`
	Type      string `short:"t"  long:"type"                        description:"Type of resource to delete."`
	Debug     bool   `           long:"debug"                       description:"Print debug information."`

	AWSAccessKeyID       string `long:"aws-access-key-id"        env:"BBL_AWS_ACCESS_KEY_ID"        description:"AWS access key id."`
	AWSSecretAccessKey   string `long:"aws-secret-access-key"    env:"BBL_AWS_SECRET_ACCESS_KEY"    description:"AWS secret access key."`
	AWSAssumeRole        string `long:"aws-assume-role"          env:"BBL_AWS_ASSUME_ROLE"          description:"AWS assume role ARN."`
	AWSSessionToken      string `long:"aws-session-token"        env:"BBL_AWS_SESSION_TOKEN"        description:"AWS session token."`
	AWSRegion            string `long:"aws-region"               env:"BBL_AWS_REGION"               description:"AWS region."`
	AzureClientID        string `long:"azure-client-id"          env:"BBL_AZURE_CLIENT_ID"          description:"Azure client id."`
	AzureClientSecret    string `long:"azure-client-secret"      env:"BBL_AZURE_CLIENT_SECRET"      description:"Azure client secret."`
	AzureTenantID        string `long:"azure-tenant-id"          env:"BBL_AZURE_TENANT_ID"          description:"Azure tenant id."`
	AzureSubscriptionID  string `long:"azure-subscription-id"    env:"BBL_AZURE_SUBSCRIPTION_ID"    description:"Azure subscription id."`
	GCPServiceAccountKey string `long:"gcp-service-account-key"  env:"BBL_GCP_SERVICE_ACCOUNT_KEY"  description:"GCP service account key path."`
	VSphereIP            string `long:"vsphere-vcenter-ip"       env:"BBL_VSPHERE_VCENTER_IP"       description:"vSphere vCenter IP address."`
	VSpherePassword      string `long:"vsphere-vcenter-password" env:"BBL_VSPHERE_VCENTER_PASSWORD" description:"vSphere vCenter password."`
	VSphereUser          string `long:"vsphere-vcenter-user"     env:"BBL_VSPHERE_VCENTER_USER"     description:"vSphere vCenter username."`
	VSphereDC            string `long:"vsphere-vcenter-dc"       env:"BBL_VSPHERE_VCENTER_DC"       description:"vSphere vCenter datacenter."`
	NSXTManagerHost      string `long:"nsxt-manager-host"        env:"BBL_NSXT_MANAGER_HOST"        description:"NSX-T manager IP address or domain name."`
	NSXTUser             string `long:"nsxt-username"            env:"BBL_NSXT_USERNAME"            description:"NSX-T manager username."`
	NSXTPassword         string `long:"nsxt-password"            env:"BBL_NSXT_PASSWORD"            description:"NSX-T manager password."`
	OpenstackAuthUrl     string `long:"openstack-auth-url"       env:"BBL_OPENSTACK_AUTH_URL"       description:"Openstack auth URL."`
	OpenstackUsername    string `long:"openstack-username"       env:"BBL_OPENSTACK_USERNAME"       description:"Openstack username."`
	OpenstackPassword    string `long:"openstack-password"       env:"BBL_OPENSTACK_PASSWORD"       description:"Openstack password."`
	OpenstackDomain      string `long:"openstack-domain-name"    env:"BBL_OPENSTACK_DOMAIN"         description:"Openstack domain name."`
	OpenstackTenant      string `long:"openstack-project-name"   env:"BBL_OPENSTACK_PROJECT"        description:"Openstack project name."`
	OpenstackRegion      string `long:"openstack-region-name"    env:"BBL_OPENSTACK_REGION"         description:"Openstack region name."`
}

const (
	AWS       = "aws"
	GCP       = "gcp"
	Azure     = "azure"
	VSphere   = "vsphere"
	NSXT      = "nsxt"
	Openstack = "openstack"
)

type OtherEnvVars struct {
}

func NewOtherEnvVars() OtherEnvVars {
	return OtherEnvVars{}
}

func (e OtherEnvVars) LoadConfig(o *Options) {
	switch o.IAAS {
	case AWS:
		if o.AWSAccessKeyID == "" {
			o.AWSAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
		}
		if o.AWSSecretAccessKey == "" {
			o.AWSSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
		}
		if o.AWSAssumeRole == "" {
			o.AWSAssumeRole = os.Getenv(("AWS_ROLE_ARN"))
		}
		if o.AWSSessionToken == "" {
			o.AWSSessionToken = os.Getenv("AWS_SESSION_TOKEN")
		}
		if o.AWSRegion == "" {
			o.AWSRegion = os.Getenv("AWS_DEFAULT_REGION")
		}
	case Azure:
		if o.AzureClientID == "" {
			o.AzureClientID = os.Getenv("ARM_CLIENT_ID")
		}
		if o.AzureClientSecret == "" {
			o.AzureClientSecret = os.Getenv("ARM_CLIENT_SECRET")
		}
		if o.AzureSubscriptionID == "" {
			o.AzureSubscriptionID = os.Getenv("ARM_SUBSCRIPTION_ID")
		}
		if o.AzureTenantID == "" {
			o.AzureTenantID = os.Getenv("ARM_TENANT_ID")
		}
	case GCP:
		if o.GCPServiceAccountKey == "" {
			o.GCPServiceAccountKey = os.Getenv("GOOGLE_CREDENTIALS")
		}
	case NSXT:
		if o.NSXTManagerHost == "" {
			o.NSXTManagerHost = os.Getenv("NSXT_MANAGER_HOST")
		}
		if o.NSXTUser == "" {
			o.NSXTUser = os.Getenv("NSXT_USERNAME")
		}
		if o.NSXTPassword == "" {
			o.NSXTPassword = os.Getenv("NSXT_PASSWORD")
		}
	case VSphere:
		if o.VSphereUser == "" {
			o.VSphereUser = os.Getenv("VSPHERE_USER")
		}
		if o.VSpherePassword == "" {
			o.VSpherePassword = os.Getenv("VSPHERE_PASSWORD")
		}
		if o.VSphereDC == "" {
			o.VSphereDC = os.Getenv("VSPHERE_DATACENTER")
		}
		if o.VSphereIP == "" {
			o.VSphereIP = os.Getenv("VSPHERE_IP")
		}
	}
}
