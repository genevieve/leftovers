# Leftovers

Clean up orphaned IAAS resources.

## Why you might be here?
- You `terraform apply`'d way back when and lost your `terraform.tfstate`
- You used the console or cli to create some infrastructure and want to clean up
- Your acceptance tests in CI failed, the container disappeared, and
infrastructure resources were tragically orphaned

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
  -i, --iaas=                  The IAAS for clean up. (default: aws) [$LEFTOVERS_IAAS]
  -n, --no-confirm             Destroy resources without prompting. This is dangerous, make good choices!
      --aws-access-key-id=     AWS access key id. [$AWS_ACCESS_KEY_ID]
      --aws-secret-access-key= AWS secret access key. [$AWS_SECRET_ACCESS_KEY]
      --aws-region=            AWS region. [$AWS_REGION]
      --azure-client-id=       Azure client id. [$AZURE_CLIENT_ID]
      --azure-client-secret=   Azure client secret. [$AZURE_CLIENT_SECRET]
      --azure-tenant-id=       Azure tenant id. [$AZURE_TENANT_ID]
      --azure-subscription-id= Azure subscription id. [$AZURE_SUBSCRIPTION_ID]

Help Options:
  -h, --help                   Show this help message
```

## AWS
### What can you delete with this?

```diff
+ iam instance profiles (& detaching roles)
+ iam roles
+ iam role policies
+ iam server certificates
+ ec2 volumes
+ ec2 tags
+ ec2 key pairs
+ ec2 instances
+ ec2 security groups
+ ec2 vpcs
+ elb load balancers
```

### What's up next?

```diff
- iam group policies
- iam user policies
- elbv2 load balancers
- ec2 eips
- ec2 enis
- s3 buckets
```

## Azure
### What can you delete with this?

```diff
+ resource groups
```

## GCP
### What's up next?

```diff
- compute disks
- compute health checks
- compute vm instances
- compute vm instance groups
- compute vm instance templates
- compute snapshots
- compute images
```

## vSphere
