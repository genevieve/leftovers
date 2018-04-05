package acceptance

import (
	"context"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/vsphere"
	. "github.com/onsi/gomega"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
)

type VSphereAcceptance struct {
	VCenterIP       string
	VCenterUser     string
	VCenterPassword string
	Datacenter      string
	Logger          *app.Logger
}

func NewVSphereAcceptance() VSphereAcceptance {
	vcenterIP := os.Getenv("BBL_VSPHERE_VCENTER_IP")
	Expect(vcenterIP).NotTo(Equal(""))

	vcenterUser := os.Getenv("BBL_VSPHERE_VCENTER_USER")
	Expect(vcenterUser).NotTo(Equal(""))

	vcenterPassword := os.Getenv("BBL_VSPHERE_VCENTER_PASSWORD")
	Expect(vcenterPassword).NotTo(Equal(""))

	datacenter := os.Getenv("BBL_VSPHERE_VCENTER_DC")
	Expect(datacenter).NotTo(Equal(""))

	return VSphereAcceptance{
		VCenterIP:       vcenterIP,
		VCenterUser:     vcenterUser,
		VCenterPassword: vcenterPassword,
		Datacenter:      datacenter,
		Logger:          app.NewLogger(os.Stdin, os.Stdout, true),
	}
}

func (v *VSphereAcceptance) CreateFolder(root, name string) {
	log.SetFlags(log.Ldate | log.Ltime)
	vCenterUrl, err := url.Parse("https://" + v.VCenterIP + "/sdk")
	Expect(err).NotTo(HaveOccurred())

	vCenterUrl.User = url.UserPassword(v.VCenterUser, v.VCenterPassword)

	vContext, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	vimClient, err := govmomi.NewClient(vContext, vCenterUrl, true)
	Expect(err).NotTo(HaveOccurred())

	datacenter, err := vsphere.DatacenterFromID(vimClient, v.Datacenter)
	Expect(err).NotTo(HaveOccurred())

	finder := find.NewFinder(vimClient.Client, true)
	finder.SetDatacenter(datacenter)
	ctx := context.Background()

	rootFolder, err := finder.Folder(ctx, root)
	Expect(err).NotTo(HaveOccurred())

	log.Println("creating folder")
	_, err = rootFolder.CreateFolder(ctx, name)
	Expect(err).NotTo(HaveOccurred())

	log.Println("verifying folder creation")
	created := false
	for !created {
		createdFolder, err := finder.Folder(ctx, name)
		if err != nil {
			log.Printf("got %+v, %s\n", createdFolder, err)
		} else {
			log.Printf("got %+v\n", createdFolder)
		}
		created = (err == nil) && (createdFolder != nil)
		if !created {
			log.Println("waiting for folder to be created")
			time.Sleep(10 * time.Second)
		}
	}
	log.Println("created folder!")
}
