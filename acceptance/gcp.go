package acceptance

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/gcp/compute"
	"github.com/genevieve/leftovers/gcp/container"
	"golang.org/x/oauth2/google"
	gcpcompute "google.golang.org/api/compute/v1"
	gcpcontainer "google.golang.org/api/container/v1"

	. "github.com/onsi/gomega"
)

type GCPAcceptance struct {
	Key       []byte
	KeyPath   string
	ProjectId string
	Zone      string
	Logger    *app.Logger
}

func NewGCPAcceptance() GCPAcceptance {
	path := os.Getenv("BBL_GCP_SERVICE_ACCOUNT_KEY")
	Expect(path).NotTo(Equal(""))

	key, err := ioutil.ReadFile(path)
	if err != nil {
		key = []byte(path)
	}

	p := struct {
		ProjectId string `json:"project_id"`
	}{}
	err = json.Unmarshal(key, &p)
	Expect(err).NotTo(HaveOccurred())

	return GCPAcceptance{
		Key:       key,
		KeyPath:   path,
		ProjectId: p.ProjectId,
		Zone:      "us-east1-b",
		Logger:    app.NewLogger(os.Stdin, os.Stdout, true),
	}
}

func (g GCPAcceptance) InsertDisk(name string) {
	config, err := google.JWTConfigFromJSON([]byte(g.Key), gcpcompute.ComputeScope)
	Expect(err).NotTo(HaveOccurred())

	service, err := gcpcompute.New(config.Client(context.Background()))
	Expect(err).NotTo(HaveOccurred())

	list, err := service.Disks.List(g.ProjectId, g.Zone).Filter(fmt.Sprintf("name eq %s", name)).Do()
	if len(list.Items) > 0 {
		return
	}

	operation, err := service.Disks.Insert(g.ProjectId, g.Zone, &gcpcompute.Disk{Name: name}).Do()

	waiter := compute.NewOperationWaiter(operation, service, g.ProjectId, g.Logger)

	err = waiter.Wait()
	Expect(err).NotTo(HaveOccurred())
}

func (g GCPAcceptance) InsertCluster(name string) {
	config, err := google.JWTConfigFromJSON([]byte(g.Key), gcpcompute.ComputeScope, gcpcontainer.CloudPlatformScope)
	Expect(err).NotTo(HaveOccurred())

	service, err := gcpcontainer.New(config.Client(context.Background()))
	Expect(err).NotTo(HaveOccurred())

	req := &gcpcontainer.CreateClusterRequest{
		Cluster: &gcpcontainer.Cluster{
			Name:             name,
			InitialNodeCount: 1,
		},
	}

	operation, err := service.Projects.Zones.Clusters.Create(g.ProjectId, g.Zone, req).Do()
	Expect(err).NotTo(HaveOccurred())

	waiter := container.NewOperationWaiter(operation, service, g.ProjectId, g.Logger)

	err = waiter.Wait()
	Expect(err).NotTo(HaveOccurred())
}
