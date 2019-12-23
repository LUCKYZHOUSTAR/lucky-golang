// +build !windows

package daemon

import (
	"os"
	"path/filepath"
	"runtime"
	"syscall"

	"mtime.com/framework/config"
	"mtime.com/framework/log"
	"mtime.com/framework/utility/app"
)

func (d *Daemon) SetQuitHandler(quit SignalHandlerFunc) {
	d.quit = quit
}

func (d *Daemon) SetStopHandler(stop SignalHandlerFunc) {
	d.stop = stop
}

func (d *Daemon) SetReloadHandler(reload SignalHandlerFunc) {
	d.reload = reload
}

func (d *Daemon) Run(f func()) {
	signal := app.GetSignalFlag()

	if *signal == "run" {
		f()
		return
	}

	// set handlers
	if d.quit == nil {
		d.quit = d.defaultQuit
	}
	if d.stop == nil {
		d.stop = d.defaultQuit
	}
	if d.reload == nil {
		d.reload = d.defaultReload
	}
	AddCommand(StringFlag(signal, "quit"), syscall.SIGQUIT, d.quit)
	AddCommand(StringFlag(signal, "stop"), syscall.SIGTERM, d.stop)
	AddCommand(StringFlag(signal, "reload"), syscall.SIGHUP, d.reload)

	logPath, err := d.getLogPath()
	if err != nil {
		d.exitWithError("get log path failed", err)
	}

	pidFilePath, err := d.getPidFilePath(logPath)
	if err != nil {
		d.exitWithError("get pid filepath failed", err)
	}

	logFileName := filepath.Join(logPath, "std.log")
	ctx := &Context{
		PidFileName: pidFilePath,
		PidFilePerm: 0644,
		LogFileName: logFileName, // 重定向stdout, stderr到文件
		LogFilePerm: 0640,
		WorkDir:     "/",
		// Umask:   027,
		// Args:        []string{"[sample]"},
	}

	if len(ActiveFlags()) > 0 {
		p, err := ctx.Search()
		if err != nil {
			d.exitWithError("send signal failed", err)
		} else {
			SendCommands(p)
		}
		return
	}

	p, err := ctx.Reborn()
	if err != nil {
		d.exitWithError("reborn failed", err)
	} else if p != nil {
		return
	}
	defer ctx.Release()

	runtime.GOMAXPROCS(runtime.NumCPU())
	go d.safeRun(f)

	log.Info("app", "started(version: %v)", config.GetAppConfig().App.Version)
	err = ServeSignals()
	if err != nil {
		log.Error("app", err.Error())
	}
	log.Info("app", "terminated")
}

func (d *Daemon) safeRun(f func()) {
	defer func() {
		if err := recover(); err != nil {
			d.exitWithError("app crashed", err)
		}
	}()

	f()
}

func (d *Daemon) exitWithError(msg string, err interface{}) {
	log.Fatal("app", "%v, error: %v", msg, err)
	os.Exit(1)
}

func (d *Daemon) getPidFilePath(logPath string) (string, error) {
	appName, err := app.GetAppName()
	if err != nil {
		return "", err
	}

	return filepath.Join(logPath, appName+".pid"), nil
}

func (d *Daemon) getLogPath() (string, error) {
	logPath := config.GetAppConfig().App.LogPath
	if logPath == "" {
		logPath = config.GetGlobalConfig().LogPath
		if logPath == "" {
			appPath, err := app.GetAppFolder()
			if err != nil {
				return "", err
			}

			logPath = filepath.Join(appPath, "logs")
		} else {
			logPath = filepath.Join(logPath, config.GetAppConfig().App.AppName)
		}
	}
	return logPath, nil
}

func (d *Daemon) defaultQuit(sig os.Signal) error {
	log.Info("app", "application is terminating...")
	// todo: process quit and stop
	// if sig == syscall.SIGQUIT {
	// 	<-done
	// }
	return ErrStop
}

func (d *Daemon) defaultReload(sig os.Signal) error {
	log.Info("app", "configuration reloaded")
	return nil
}
