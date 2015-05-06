package metricdriver

import (
	"github.com/docker/libcontainer/cgroups"
	"github.com/docker/libcontainer/cgroups/fs"
	"github.com/docker/libcontainer/configs"
)

func Get(id, parent string, pid int) (*cgroups.Stats, error) {

	c := &configs.Cgroup{
		Name:   id,
		Parent: parent,
	}

	stats, err := fs.GetAllStats(c, pid)

	if err != nil {
		return nil, err
	}

	return stats, nil
}
