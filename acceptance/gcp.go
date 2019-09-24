package acceptance

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/gcp/compute"
	"golang.org/x/oauth2/google"
	gcpcompute "google.golang.org/api/compute/v1"

	. "github.com/onsi/gomega"
)

type GCPAcceptance struct {
	Key       []byte
	KeyPath   string
	ProjectId string
	Zone      string
	Region    string
	Logger    *app.Logger
}

func NewGCPAcceptance() GCPAcceptance {
	path := os.Getenv("BBL_GCP_SERVICE_ACCOUNT_KEY")
	Expect(path).NotTo(Equal(""), "Missing $BBL_GCP_SERVICE_ACCOUNT_KEY.")

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
		Region:    "us-east1",
		Logger:    app.NewLogger(os.Stdin, os.Stdout, true, false),
	}
}

func (g GCPAcceptance) InsertDisk(name string) {
	config, err := google.JWTConfigFromJSON([]byte(g.Key), gcpcompute.ComputeScope)
	Expect(err).NotTo(HaveOccurred())

	service, err := gcpcompute.New(config.Client(context.Background()))
	Expect(err).NotTo(HaveOccurred())

	list, err := service.Disks.List(g.ProjectId, g.Zone).Filter(fmt.Sprintf("name eq %s", name)).Do()
	Expect(err).NotTo(HaveOccurred())
	if len(list.Items) > 0 {
		return
	}

	operation, err := service.Disks.Insert(g.ProjectId, g.Zone, &gcpcompute.Disk{Name: name}).Do()
	Expect(err).NotTo(HaveOccurred())
	waiter := compute.NewOperationWaiter(operation, service, g.ProjectId, g.Logger)

	err = waiter.Wait()
	Expect(err).NotTo(HaveOccurred())
}

func (g GCPAcceptance) InsertCloudRouter(name string) {
	config, err := google.JWTConfigFromJSON([]byte(g.Key), gcpcompute.ComputeScope)
	Expect(err).NotTo(HaveOccurred())

	service, err := gcpcompute.New(config.Client(context.Background()))
	Expect(err).NotTo(HaveOccurred())

	list, err := service.Routers.List(g.ProjectId, g.Region).Filter(fmt.Sprintf("name eq %s", name)).Do()
	Expect(err).NotTo(HaveOccurred())
	if len(list.Items) > 0 {
		return
	}

	network, err := service.Networks.Get(g.ProjectId, "default").Do()
	Expect(err).NotTo(HaveOccurred())

	router := &gcpcompute.Router{
		Name:    name,
		Network: network.SelfLink,
	}
	operation, err := service.Routers.Insert(g.ProjectId, g.Region, router).Do()
	Expect(err).NotTo(HaveOccurred())

	waiter := compute.NewOperationWaiter(operation, service, g.ProjectId, g.Logger)

	err = waiter.Wait()
	Expect(err).NotTo(HaveOccurred())
}
