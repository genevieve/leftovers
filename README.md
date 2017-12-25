# Leftovers

Clean up orphaned IAAS resources (mostly created by `bbl`.)

## Why you might be here?
- You `terraform apply`'d way back when and lost your `terraform.tfstate`
- You used the console or cli to create some infrastructure and want to clean up
- Your acceptance tests in CI failed, the container disappeared, and
infrastructure resources were tragically orphaned

## AWS
### What can you delete with this?

```diff
+ iam instance profiles (& detaching roles)
+ iam roles
+ iam role policies
+ iam user policies
+ iam server certificates
+ ec2 volumes
+ ec2 tags
+ ec2 key pairs
+ ec2 instances
+ ec2 security groups (Note: for groups that reference each other, the first pass will delete the references. Run through a second time.)
+ ec2 vpcs
+ ec2 subnets
+ ec2 route tables
+ ec2 internet gateways
+ ec2 network interfaces
+ elb load balancers
+ s3 buckets
```

### What's up next?

```diff
- iam group policies
- elbv2 load balancers
- ec2 eips
```

## Azure
### What can you delete with this?

```diff
+ resource groups
```

## GCP
### What can you delete with this?

```diff
+ compute disks
+ compute networks
+ compute vm instances
+ compute backend services
+ compute http health checks
+ compute target pools
+ compute forwarding rules
```
### What's up next?

```diff
- compute routes
- compute health checks
- compute https health checks
- compute firewall rules
- compute vm instance groups
- compute vm instance templates
- compute snapshots
- compute images
```

## vSphere
### What's up next?

## Installation

[Install go.](https://golang.org/doc/install)

```
$  go get github.com/genevievelesperance/leftovers
```

## Usage

```
Usage:
  leftovers [OPTIONS]

Application Options:
  -i, --iaas=                    The IAAS for clean up. (default: aws) [$BBL_IAAS]
  -n, --no-confirm               Destroy resources without prompting. This is dangerous, make good choices!
      --aws-access-key-id=       AWS access key id. [$BBL_AWS_ACCESS_KEY_ID]
      --aws-secret-access-key=   AWS secret access key. [$BBL_AWS_SECRET_ACCESS_KEY]
      --aws-region=              AWS region. [$BBL_AWS_REGION]
      --azure-client-id=         Azure client id. [$BBL_AZURE_CLIENT_ID]
      --azure-client-secret=     Azure client secret. [$BBL_AZURE_CLIENT_SECRET]
      --azure-tenant-id=         Azure tenant id. [$BBL_AZURE_TENANT_ID]
      --azure-subscription-id=   Azure subscription id. [$BBL_AZURE_SUBSCRIPTION_ID]
      --gcp-service-account-key= GCP service account key path. [$BBL_GCP_SERVICE_ACCOUNT_KEY]

Help Options:
  -h, --help                     Show this help message
```

