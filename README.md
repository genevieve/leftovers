# Leftovers

Clean up orphaned IAAS resources.

## Why you might be here?
- You `terraform apply`'d way back when and lost your `terraform.tfstate`
- You used the console or cli to create some infrastructure and want to clean up
- Your acceptance tests in CI failed, the container disappeared, and
infrastructure resources were tragically orphaned

## Currently deleting
- iam instance profiles
- iam roles
- iam server certificates
- ec2 volumes
- ec2 tags
- ec2 key pairs
- ec2 instances
- elb load balancers

## Upcoming
- elbv2 load balancers
- ec2 eips
- ec2 vpcs
- ec2 enis
- s3 buckets
