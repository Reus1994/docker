package daemon

import (
	"github.com/dotcloud/docker/daemon/graphdriver/devmapper"
	"github.com/dotcloud/docker/utils"
)

func (container *Container) Sweep() error {
	device := "eth0"
	if _, err := container.daemon.ExecIn(container.ID, "ip", []string{"link", "show", device}); err != nil {
		utils.Debugf("%s", err)
	} else {
		if _, err := container.daemon.ExecIn(container.ID, "ip", []string{"link", "delete", device}); err != nil {
			return err
		}
	}
	container.releaseNetwork()
	container.Config.Ip = ""
	container.Config.NetworkDisabled = true
	container.hostConfig.NetworkMode = "none"
	if err := container.ToDisk(); err != nil {
		return err
	}

	if err := container.Stop(10); err != nil {
		return err
	}

	return nil
}

func (container *Container) DeviceIsBusy() (bool, error) {
	if container.daemon.driver.String() == "devicemapper" {
		driver := container.daemon.driver.(*devmapper.Driver)
		utils.Debugf("%v", driver)
		devices := driver.DeviceSet
		if opencount, err := devices.OpenCount(container.ID); err != nil {
			utils.Debugf("%v", err)
			return false, err
		} else {
			utils.Debugf("%s: opencount=%d", container.ID, opencount)
			return opencount != 0, nil
		}
	}
	return false, nil
}
