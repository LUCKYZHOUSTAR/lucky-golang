package daemon

import (
	"strings"

	"mtime.com/framework/config"
)

type Daemon struct {
	quit   SignalHandlerFunc
	stop   SignalHandlerFunc
	reload SignalHandlerFunc
}

func (d *Daemon) SetVersion(version string) {
	version = strings.Replace(version, " ", ".", -1)
	version = strings.Replace(version, "-", "", -1)
	version = strings.Replace(version, ":", "", -1)
	if version != "" {
		config.GetAppConfig().App.Version = version
	}
}
