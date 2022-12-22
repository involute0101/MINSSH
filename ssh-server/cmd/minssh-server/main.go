package main

import (
	"fmt"
	"github.com/SunnyQjm/daemon"
	"log"
	"minlib/common"
	"os"
	"runtime"
	sshServer "ssh-server/cmd"
)

const (
	name                  = "minssh-server"                     // 服务的名字
	description           = "minssh-server programme"           // 服务描述
	defaultConfigFilePath = "/usr/local/etc/minssh/sshconf.ini" // MIR配置文件路径
)

var stdlog, errlog *log.Logger

type Service struct {
	daemon.Daemon
}

func (service *Service) Manage() (string, error) {
	usage := "Usage: minssh-server install | remove | start | stop | status"

	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	starter := sshServer.NewMINSSHStarter()
	starter.Start()
	return "", nil
}

func init() {
	stdlog = log.New(os.Stdout, "", 0)
	errlog = log.New(os.Stderr, "", 0)
}

func main() {
	sysType := runtime.GOOS

	daemonKind := daemon.SystemDaemon
	if sysType == "darwin" {
		// macos 系统
		daemonKind = daemon.GlobalDaemon
	} else if sysType == "linux" {
		// Linux 系统
		daemonKind = daemon.SystemDaemon
	} else if sysType == "windows" {
		// Windows 系统
		daemonKind = daemon.SystemDaemon
	} else {
		common.LogFatal("Not support system: ", sysType)
	}

	srv, err := daemon.New(name, description, daemonKind)
	if err != nil {
		errlog.Println("Error: ", err)
		os.Exit(1)
	}
	service := &Service{srv}
	status, err := service.Manage()
	if err != nil {
		errlog.Println(status, "\nError: ", err)
		os.Exit(1)
	}
	fmt.Println(status)
}
