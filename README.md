# Leftovers

Go library for cleaning up **orphaned IAAS resources**.

## Why you might be here?
- You `terraform apply`'d way back when and lost your `terraform.tfstate`
- You used the console or cli to create some infrastructure and want to clean up
- Your acceptance tests in CI failed, the container disappeared, and
infrastructure resources were tragically orphaned

## Command-line utility

[Repo](https://github.com/genevievelesperance/leftovers-cli)

```
brew tap genevievelesperance/tap
brew install leftovers
```

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
+ ec2 security groups (Note: for sgs that reference each other, the first pass will delete the references. Run through a second time.)
+ ec2 vpcs
+ ec2 subnets
+ ec2 route tables
+ ec2 internet gateways
+ ec2 network interfaces
+ elb load balancers
+ elbv2 load balancers
+ elbv2 target groups
+ s3 buckets
```

### What's up next?

```diff
- iam group policies
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
+ compute addresses
+ compute backend services
+ compute disks
+ compute firewalls
+ compute forwarding rules
+ compute global forwarding rules
+ compute global health checks
+ compute http health checks
+ compute https health checks
+ compute networks
+ compute target pools
+ compute target https proxies
+ compute target http proxies
+ compute url maps
+ compute vm instances
+ compute vm instance groups
```
### What's up next?

```diff
- compute routes
- compute vm instance templates
- compute snapshots
- compute images
```

## vSphere
### What's up next?

