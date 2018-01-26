package aws

import (
	"errors"

	awslib "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	awsec2 "github.com/aws/aws-sdk-go/service/ec2"
	awselb "github.com/aws/aws-sdk-go/service/elb"
	awselbv2 "github.com/aws/aws-sdk-go/service/elbv2"
	awsiam "github.com/aws/aws-sdk-go/service/iam"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/genevievelesperance/leftovers/aws/ec2"
	"github.com/genevievelesperance/leftovers/aws/elb"
	"github.com/genevievelesperance/leftovers/aws/elbv2"
	"github.com/genevievelesperance/leftovers/aws/iam"
	"github.com/genevievelesperance/leftovers/aws/s3"
)

type resource interface {
	List(filter string) (map[string]string, error)
	Delete(items map[string]string) error
}

type Deleter struct {
	resources []resource
}

func (d Deleter) Delete(filter string) error {
	for _, r := range d.resources {
		items, err := r.List(filter)
		if err != nil {
			return err
		}

		r.Delete(items)
	}
	return nil
}

func NewDeleter(logger logger, accessKeyId, secretAccessKey, region string) (Deleter, error) {
	if accessKeyId == "" {
		return Deleter{}, errors.New("Missing aws access key id.")
	}

	if secretAccessKey == "" {
		return Deleter{}, errors.New("Missing secret access key.")
	}

	if region == "" {
		return Deleter{}, errors.New("Missing region.")
	}

	config := &awslib.Config{
		Credentials: credentials.NewStaticCredentials(accessKeyId, secretAccessKey, ""),
		Region:      awslib.String(region),
	}
	sess := session.New(config)

	iamClient := awsiam.New(sess)
	ec2Client := awsec2.New(sess)
	elbClient := awselb.New(sess)
	elbv2Client := awselbv2.New(sess)
	s3Client := awss3.New(sess)

	rolePolicies := iam.NewRolePolicies(iamClient, logger)
	userPolicies := iam.NewUserPolicies(iamClient, logger)
	accessKeys := iam.NewAccessKeys(iamClient, logger)
	internetGateways := ec2.NewInternetGateways(ec2Client, logger)
	routeTables := ec2.NewRouteTables(ec2Client, logger)
	subnets := ec2.NewSubnets(ec2Client, logger)
	bucketManager := s3.NewBucketManager(region)

	return Deleter{
		resources: []resource{
			iam.NewRoles(iamClient, logger, rolePolicies),
			iam.NewUsers(iamClient, logger, userPolicies, accessKeys),
			iam.NewPolicies(iamClient, logger),
			iam.NewInstanceProfiles(iamClient, logger),
			iam.NewServerCertificates(iamClient, logger),

			ec2.NewAddresses(ec2Client, logger),
			ec2.NewKeyPairs(ec2Client, logger),
			ec2.NewInstances(ec2Client, logger),
			ec2.NewSecurityGroups(ec2Client, logger),
			ec2.NewTags(ec2Client, logger),
			ec2.NewVolumes(ec2Client, logger),
			ec2.NewNetworkInterfaces(ec2Client, logger),
			ec2.NewVpcs(ec2Client, logger, routeTables, subnets, internetGateways),

			elb.NewLoadBalancers(elbClient, logger),
			elbv2.NewLoadBalancers(elbv2Client, logger),
			elbv2.NewTargetGroups(elbv2Client, logger),

			s3.NewBuckets(s3Client, logger, bucketManager),
		},
	}, nil
}
