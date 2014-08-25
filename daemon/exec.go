package daemon

import (
	"github.com/docker/docker/engine"
	"github.com/docker/docker/pkg/log"
)

func (daemon *Daemon) ContainerExec(job *engine.Job) engine.Status {
	if len(job.Args) != 1 {
		return job.Errorf("Usage: %s CONTAINER\n", job.Name)
	}

	var (
		name    = job.Args[0]
		command = job.Getenv("command")
		args    = job.GetenvList("args")
	)

	log.Debugf("name %s, command %s, args %s", name, command, args)

	if container := daemon.Get(name); container != nil {
		if !container.State.IsRunning() {
			return job.Errorf("Container %s is not running", name)
		}

		output, err := daemon.ExecIn(container.ID, command, args)
		log.Infof("%s", output)
		job.Setenv("output", string(output))
		if err != nil {
			return job.Error(err)
		}

		return engine.StatusOK
	}
	return job.Errorf("No such container: %s", name)
}
