package nsxt

import (
	"errors"
	"fmt"
	"sync"

	"github.com/fatih/color"
	"github.com/genevieve/leftovers/nsxt/common"
	"github.com/genevieve/leftovers/nsxt/logicalrouting"
	nsxt "github.com/vmware/go-vmware-nsxt"
)

type resource interface {
	List(filter string) ([]common.Deletable, error)
	Type() string
}

type Leftovers struct {
	logger    logger
	resources []resource
}

// List will print all the resources that contain
// the provided filter in the resource's identifier.
func (l Leftovers) List(filter string) {
	l.logger.NoConfirm()
	var deletables []common.Deletable

	for _, r := range l.resources {
		list, err := r.List(filter)
		if err != nil {
			l.logger.Println(color.YellowString(err.Error()))
		}

		deletables = append(deletables, list...)
	}

	for _, d := range deletables {
		l.logger.Println(fmt.Sprintf("[%s: %s]", d.Type(), d.Name()))
	}
}

// Types will print all the resource types that can
// be deleted on this IaaS.
func (l Leftovers) Types() {
	for _, r := range l.resources {
		l.logger.Println(r.Type())
	}
}

// Delete will collect all resources that contain
// the provided filter in the resource's identifier, prompt
// you to confirm deletion (if enabled), and delete those
// that are selected.
func (l Leftovers) Delete(filter string) error {
	deletables := [][]common.Deletable{}

	for _, r := range l.resources {
		list, err := r.List(filter)
		if err != nil {
			l.logger.Println(color.YellowString(err.Error()))
		}

		deletables = append(deletables, list)
	}

	l.asyncDelete(deletables)

	return nil
}

// DeleteType will collect all resources of the provied type that contain
// the provided filter in the resource's identifier, prompt
// you to confirm deletion (if enabled), and delete those
// that are selected.
func (l Leftovers) DeleteType(filter, rType string) error {
	deletables := [][]common.Deletable{}

	for _, r := range l.resources {
		if r.Type() == rType {
			list, err := r.List(filter)
			if err != nil {
				l.logger.Println(color.YellowString(err.Error()))
			}

			deletables = append(deletables, list)
		}
	}

	l.asyncDelete(deletables)

	return nil
}

func (l Leftovers) asyncDelete(deletables [][]common.Deletable) {
	var wg sync.WaitGroup

	for _, list := range deletables {
		for _, d := range list {
			wg.Add(1)

			go func(d common.Deletable) {
				defer wg.Done()

				l.logger.Println(fmt.Sprintf("[%s: %s] Deleting...", d.Type(), d.Name()))

				err := d.Delete()
				if err != nil {
					l.logger.Println(fmt.Sprintf("[%s: %s] %s", d.Type(), d.Name(), color.YellowString(err.Error())))
				} else {
					l.logger.Println(fmt.Sprintf("[%s: %s] %s", d.Type(), d.Name(), color.GreenString("Deleted!")))
				}
			}(d)
		}

		wg.Wait()
	}
}

func NewLeftovers(logger logger, managerHost, user, password string) (Leftovers, error) {
	if managerHost == "" {
		return Leftovers{}, errors.New("Missing NSX-T manager host.")
	}

	if user == "" {
		return Leftovers{}, errors.New("Missing NSX-T username.")
	}

	if password == "" {
		return Leftovers{}, errors.New("Missing NSX-T password.")
	}

	nsxtClient, err := nsxt.NewAPIClient(&nsxt.Configuration{
		BasePath: fmt.Sprintf("https://%s/api/v1", managerHost),
		UserName: user,
		Password: password,
		Host:     managerHost,
		Insecure: true,
		RetriesConfiguration: nsxt.ClientRetriesConfiguration{
			MaxRetries:    1,
			RetryMinDelay: 100,
			RetryMaxDelay: 500,
		},
	})
	if err != nil {
		return Leftovers{}, fmt.Errorf("Error creating NSX-T API client: %s", err)
	}

	return Leftovers{
		logger: logger,
		resources: []resource{
			logicalrouting.NewTier1Routers(nsxtClient.LogicalRoutingAndServicesApi, nsxtClient.Context, logger),
			// TBD
		},
	}, nil
}
