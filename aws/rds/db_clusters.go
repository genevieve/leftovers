package rds

import (
	"fmt"
	awsrds "github.com/aws/aws-sdk-go/service/rds"

	"github.com/genevieve/leftovers/common"
)

//go:generate faux --interface dbClustersClient --output fakes/db_clusters_client.go
type dbClustersClient interface {
	DescribeDBClusters(*awsrds.DescribeDBClustersInput) (*awsrds.DescribeDBClustersOutput, error)
	DeleteDBCluster(*awsrds.DeleteDBClusterInput) (*awsrds.DeleteDBClusterOutput, error)
}

type DBClusters struct {
	client dbClustersClient
	logger logger
}

func NewDBClusters(client dbClustersClient, logger logger) DBClusters {
	return DBClusters{
		client: client,
		logger: logger,
	}
}

func (d DBClusters) List(filter string, regex bool) ([]common.Deletable, error) {
	dbClusters, err := d.client.DescribeDBClusters(&awsrds.DescribeDBClustersInput{})
	if err != nil {
		return nil, fmt.Errorf("Describing RDS DB Clusters: %s", err)
	}

	var resources []common.Deletable
	for _, db := range dbClusters.DBClusters {
		r := NewDBCluster(d.client, db.DBClusterIdentifier)

		if *db.Status == "deleting" {
			continue
		}

		if !common.ResourceMatches(r.Name(),  filter, regex) {
			continue
		}

		proceed := d.logger.PromptWithDetails(r.Type(), r.Name())
		if !proceed {
			continue
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (d DBClusters) Type() string {
	return "rds-db-cluster"
}
