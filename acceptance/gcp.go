package acceptance

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/gcp/compute"
	"golang.org/x/oauth2/google"
	gcpcompute "google.golang.org/api/compute/v1"

	. "github.com/onsi/gomega"
)

const BBL_GCP_ZONE = "us-east1-b"

type GCPAcceptance struct {
	Key       []byte
	KeyPath   string
	ProjectId string
	Zone      string
	Logger    *app.Logger
}

func NewGCPAcceptance() *GCPAcceptance {
	return &GCPAcceptance{}
}

func (a *GCPAcceptance) ReadyToTest() bool {
	iaas := os.Getenv("LEFTOVERS_ACCEPTANCE")
	if iaas == "" {
		return false
	}

	if strings.ToLower(iaas) != "gcp" {
		return false
	}

	path := os.Getenv("BBL_GCP_SERVICE_ACCOUNT_KEY")
	if path == "" {
		return false
	}

	key, err := ioutil.ReadFile(path)
	Expect(err).NotTo(HaveOccurred())

	p := struct {
		ProjectId string `json:"project_id"`
	}{}
	err = json.Unmarshal(key, &p)
	Expect(err).NotTo(HaveOccurred())

	a.Key = key
	a.KeyPath = path
	a.ProjectId = p.ProjectId
	a.Zone = BBL_GCP_ZONE

	logger := app.NewLogger(os.Stdin, os.Stdout, true)
	a.Logger = logger

	return true
}

func (a *GCPAcceptance) InsertDisk(name string) {
	config, err := google.JWTConfigFromJSON([]byte(a.Key), gcpcompute.ComputeScope)
	Expect(err).NotTo(HaveOccurred())

	service, err := gcpcompute.New(config.Client(context.Background()))
	Expect(err).NotTo(HaveOccurred())

	list, err := service.Disks.List(a.ProjectId, a.Zone).Filter(fmt.Sprintf("name eq %s", name)).Do()
	if len(list.Items) > 0 {
		return
	}

	operation, err := service.Disks.Insert(a.ProjectId, a.Zone, &gcpcompute.Disk{Name: name}).Do()

	waiter := compute.NewOperationWaiter(operation, service, a.ProjectId, a.Logger)

	err = waiter.Wait()
	Expect(err).NotTo(HaveOccurred())
}
