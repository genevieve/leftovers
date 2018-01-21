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

type logger interface {
	Printf(m string, a ...interface{})
	Prompt(m string) bool
}

type resource interface {
	Delete() error
}

type Deleter struct {
	resources []resource
}

func (d Deleter) Delete(filter string) error {
	for _, r := range d.resources {
		if err := r.Delete(); err != nil {
			return err
		}
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

	ro := iam.NewRoles(iamClient, logger, rolePolicies)
	us := iam.NewUsers(iamClient, logger, userPolicies, accessKeys)
	po := iam.NewPolicies(iamClient, logger)
	ip := iam.NewInstanceProfiles(iamClient, logger)
	sc := iam.NewServerCertificates(iamClient, logger)

	ad := ec2.NewAddresses(ec2Client, logger)
	ke := ec2.NewKeyPairs(ec2Client, logger)
	in := ec2.NewInstances(ec2Client, logger)
	se := ec2.NewSecurityGroups(ec2Client, logger)
	ta := ec2.NewTags(ec2Client, logger)
	vo := ec2.NewVolumes(ec2Client, logger)
	ni := ec2.NewNetworkInterfaces(ec2Client, logger)
	vp := ec2.NewVpcs(ec2Client, logger, routeTables, subnets, internetGateways)

	l1 := elb.NewLoadBalancers(elbClient, logger)
	l2 := elbv2.NewLoadBalancers(elbv2Client, logger)

	bu := s3.NewBuckets(s3Client, logger, bucketManager)

	resources := []resource{ip, ro, us, us, po, l1, l2, sc, vo, ta, ad, ke, in, se, bu, ni, vp}

	return Deleter{
		resources: resources,
	}, nil
}
