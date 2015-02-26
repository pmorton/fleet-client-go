package client

import (
	"github.com/coreos/fleet/schema"
	"github.com/juju/errgo"
	"errors"
	"fmt"
	execPkg "os/exec"
)

const (
	FLEETCTL        = "fleetctl"
	ENDPOINT_OPTION = "--endpoint"
	ENDPOINT_VALUE  = "http://127.0.0.1:4001"
)

type ClientCLI struct {
	etcdPeer string
}

func NewClientCLI() FleetClient {
	return NewClientCLIWithPeer(ENDPOINT_VALUE)
}

func NewClientCLIWithPeer(etcdPeer string) FleetClient {
	return &ClientCLI{
		etcdPeer: etcdPeer,
	}
}

func (this *ClientCLI) Submit(name, filePath string) error {
	cmd := execPkg.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "submit", filePath)
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) UnitStatus(name string) (UnitStatus, error) {
	units, err := this.StatusAll();
	if err != nil {
		return UnitStatus{}, err
	}
	for _, unit := range units {
			if unit.Unit == name {
					return unit, nil
			}
	}
	return UnitStatus{}, errors.New("Unit not found")
}

func (this *ClientCLI) Unit(name string) (*schema.Unit, error) {
	return nil, fmt.Errorf("Method not implemented: ClientCLI.Unit")
}

func (this *ClientCLI) Start(names []string) error {
	args := []string{ENDPOINT_OPTION, this.etcdPeer, "start", "--no-block=false"}
	for _, name := range names {
		args = append(args,name)
	}
	cmd := execPkg.Command(FLEETCTL, args...)
	buf, err := exec(cmd)
fmt.Printf("BUFFER: %s",buf)
	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Stop(names []string) error {
	args := []string{ENDPOINT_OPTION, this.etcdPeer, "stop", "--no-block=false"}
	for _, name := range names {
		args = append(args,name)
	}
	cmd := execPkg.Command(FLEETCTL, args...)
	_, err := exec(cmd)
	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Load(name string) error {
	cmd := execPkg.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "load", "--no-block=false", name)
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Unload(name string) error {
	cmd := execPkg.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "unload", "--no-block=false", name)
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Destroy(names []string) error {
	args := []string{ENDPOINT_OPTION, this.etcdPeer, "destroy"}
	for _, name := range names {
		args = append(args,name)
	}

	cmd := execPkg.Command(FLEETCTL, args... )
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}
