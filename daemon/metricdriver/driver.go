package metricdriver

import (
	"github.com/dotcloud/docker/pkg/libcontainer/cgroups"
	"github.com/dotcloud/docker/pkg/libcontainer/cgroups/fs"
)

func Get(id, parent string) (*cgroups.Stats, error) {

	c := &cgroups.Cgroup{
		Name:   id,
		Parent: parent,
	}

	stats, err := fs.GetStats(c)

	if err != nil {
		return nil, err
	}

	return stats, nil
}
