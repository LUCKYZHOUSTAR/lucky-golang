package daemon

func (d *Daemon) SetQuitHandler(quit SignalHandlerFunc) {
	// this.quit = quit
}

func (d *Daemon) SetStopHandler(stop SignalHandlerFunc) {
	// this.stop = stop
}

func (d *Daemon) SetReloadHandler(reload SignalHandlerFunc) {
	// this.reload = reload
}

func (d *Daemon) Run(f func()) {
	f()
}
