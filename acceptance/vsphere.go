package acceptance

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/genevieve/leftovers/app"
	"github.com/genevieve/leftovers/vsphere"
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

func NewVSphereAcceptance() *VSphereAcceptance {
	return &VSphereAcceptance{}
}

func (v *VSphereAcceptance) ReadyToTest() bool {
	iaas := os.Getenv("LEFTOVERS_ACCEPTANCE")
	if iaas == "" {
		return false
	}

	if strings.ToLower(iaas) != "vsphere" {
		return false
	}

	v.VCenterIP = os.Getenv("BBL_VSPHERE_VCENTER_IP")
	v.VCenterUser = os.Getenv("BBL_VSPHERE_VCENTER_USER")
	v.VCenterPassword = os.Getenv("BBL_VSPHERE_VCENTER_PASSWORD")
	v.Datacenter = os.Getenv("BBL_VSPHERE_VCENTER_DC")

	logger := app.NewLogger(os.Stdin, os.Stdout, true)
	v.Logger = logger

	return true
}

func (v *VSphereAcceptance) CreateFolder(root, name string) error {
	vCenterUrl, err := url.Parse("https://" + v.VCenterIP + "/sdk")
	if err != nil {
		return fmt.Errorf("Could not parse vCenter IP \"%s\" as IP address or URL.", v.VCenterIP)
	}

	vCenterUrl.User = url.UserPassword(v.VCenterUser, v.VCenterPassword)

	vContext, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()
	vimClient, err := govmomi.NewClient(vContext, vCenterUrl, true)
	if err != nil {
		return fmt.Errorf("Error setting up client: %s", err)
	}

	datacenter, err := vsphere.DatacenterFromID(vimClient, v.Datacenter)
	if err != nil {
		return fmt.Errorf("Failed to get datacenter: %s", err)
	}

	finder := find.NewFinder(vimClient.Client, true)
	finder.SetDatacenter(datacenter)
	ctx := context.Background()

	rootFolder, err := finder.Folder(ctx, root)
	if err != nil {
		return err
	}

	_, err = rootFolder.CreateFolder(ctx, name)
	if err != nil {
		return fmt.Errorf("Creating test folder: %s", err)
	}

	return nil
}
