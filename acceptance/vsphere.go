package acceptance

import (
	"context"
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
	Expect(vcenterIP).NotTo(Equal(""), "Missing $BBL_VSPHERE_VCENTER_IP.")

	vcenterUser := os.Getenv("BBL_VSPHERE_VCENTER_USER")
	Expect(vcenterUser).NotTo(Equal(""), "Missing $BBL_VSPHERE_VCENTER_USER.")

	vcenterPassword := os.Getenv("BBL_VSPHERE_VCENTER_PASSWORD")
	Expect(vcenterPassword).NotTo(Equal(""), "Missing $BBL_VSPHERE_VCENTER_PASSWORD.")

	datacenter := os.Getenv("BBL_VSPHERE_VCENTER_DC")
	Expect(datacenter).NotTo(Equal(""), "Missing $BBL_VSPHERE_VCENTER_DC.")

	return VSphereAcceptance{
		VCenterIP:       vcenterIP,
		VCenterUser:     vcenterUser,
		VCenterPassword: vcenterPassword,
		Datacenter:      datacenter,
		Logger:          app.NewLogger(os.Stdin, os.Stdout, true),
	}
}

func (v *VSphereAcceptance) CreateFolder(root, name string) {
	vCenterUrl, err := url.Parse("https://" + v.VCenterIP + "/sdk")
	Expect(err).NotTo(HaveOccurred())

	vCenterUrl.User = url.UserPassword(v.VCenterUser, v.VCenterPassword)

	vContext, cancel := context.WithTimeout(context.Background(), time.Minute*5)
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

	_, err = rootFolder.CreateFolder(ctx, name)
	Expect(err).NotTo(HaveOccurred())
}
