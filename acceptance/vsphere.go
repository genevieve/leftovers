package acceptance

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/vsphere"
	. "github.com/onsi/gomega"
	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/find"
	"github.com/vmware/govmomi/object"
	"github.com/vmware/govmomi/vim25/types"
)

type VSphereAcceptance struct {
	Datacenter      string
	Datastore       string
	ResourcePool    string
	VCenterIP       string
	VCenterUser     string
	VCenterPassword string
	VCenterClient   *govmomi.Client
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

	datastore := os.Getenv("BBL_VSPHERE_VCENTER_DS")
	Expect(datastore).NotTo(Equal(""), "Missing $BBL_VSPHERE_VCENTER_DS.")

	resourcePool := os.Getenv("BBL_VSPHERE_VCENTER_RP")
	Expect(datastore).NotTo(Equal(""), "Missing $BBL_VSPHERE_VCENTER_RP.")

	vCenterUrl, err := url.Parse("https://" + vcenterIP + "/sdk")
	Expect(err).NotTo(HaveOccurred())

	vCenterUrl.User = url.UserPassword(vcenterUser, vcenterPassword)

	vContext, _ := context.WithTimeout(context.Background(), time.Minute*5)

	vimClient, err := govmomi.NewClient(vContext, vCenterUrl, true)
	Expect(err).NotTo(HaveOccurred())

	return VSphereAcceptance{
		VCenterIP:       vcenterIP,
		VCenterUser:     vcenterUser,
		VCenterPassword: vcenterPassword,
		Datacenter:      datacenter,
		Datastore:       datastore,
		ResourcePool:    resourcePool,
		VCenterClient:   vimClient,
		Logger:          app.NewLogger(os.Stdin, os.Stdout, true, false),
	}
}

func (v *VSphereAcceptance) CreateFolder(root, name string) *object.Folder {
	rootFolder := v.FindFolder(root)

	ctx := context.Background()
	folder, err := rootFolder.CreateFolder(ctx, name)
	Expect(err).NotTo(HaveOccurred())

	return folder
}

func (v *VSphereAcceptance) FindFolder(folder string) *object.Folder {
	searcher := object.NewSearchIndex(v.VCenterClient.Client)
	ctx := context.Background()
	result, err := searcher.FindByInventoryPath(ctx, fmt.Sprintf("/%s/vm/%s", v.Datacenter, folder))
	Expect(err).NotTo(HaveOccurred())
	rootFolder, ok := result.(*object.Folder)
	Expect(ok).To(BeTrue(), "object was not of type 'folder'")

	return rootFolder
}

func (v *VSphereAcceptance) CreateVM(folder *object.Folder, name string) {
	spec := &types.VirtualMachineConfigSpec{
		Name: name,
	}

	spec.Files = &types.VirtualMachineFileInfo{
		VmPathName: fmt.Sprintf("[%s]", v.Datastore),
	}

	datacenter, err := vsphere.DatacenterFromID(v.VCenterClient, v.Datacenter)
	Expect(err).NotTo(HaveOccurred())

	finder := find.NewFinder(v.VCenterClient.Client, true)
	finder.SetDatacenter(datacenter)

	ctx := context.Background()
	rootPool, err := finder.ResourcePoolOrDefault(ctx, v.ResourcePool)
	Expect(err).NotTo(HaveOccurred())

	task, err := folder.CreateVM(ctx, *spec, rootPool, nil)
	Expect(err).NotTo(HaveOccurred())

	err = task.Wait(ctx)
	Expect(err).NotTo(HaveOccurred())
}
