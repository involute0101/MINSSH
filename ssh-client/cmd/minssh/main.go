package main

import (
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/urfave/cli/v2"
	"gopkg.in/ini.v1"
	"log"
	"os"
	ssh "sshlib"
	"strings"
)

const defaultConfigFilePath = "/usr/local/etc/minssh/sshconf.ini"
const defaultTCPConfigFilePath = "/usr/local/etc/mir/tcpstackconf.ini"

func main() {
	var configFilePath string
	tcpConfigFilePath := "/usr/local/etc/mir/tcpstackconf.ini"
	minsshClientApp := cli.NewApp()
	minsshClientApp.Name = "MINSSH-Client"
	minsshClientApp.Usage = " MINSSH-Client daemon program "
	minsshClientApp.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "f",
			Value:       defaultConfigFilePath,
			Usage:       "Config file path for MINSSH-Client",
			Destination: &configFilePath,
			Required:    false,
		},
		&cli.StringFlag{
			Name:        "tcpf",
			Value:       defaultTCPConfigFilePath,
			Usage:       "Config file path for MINTCP",
			Destination: &tcpConfigFilePath,
			Required:    false,
		},
	}
	minsshClientApp.Action = func(context *cli.Context) error {
		connInfo := context.Args().Get(0) //Get remote host information, similar to: user@8.8.8.8;
		if !strings.Contains(connInfo, "@") {
			log.Fatal("minssh connection failed, remote host configuration information error! \n")
			return nil
		}
		arr := strings.Split(connInfo, "@")
		userName := arr[0]
		remoteHost := arr[1]
		var passwd string
		fmt.Print("Password:")
		passwdBytes, scanErr := gopass.GetPasswd()
		if scanErr != nil {
			log.Fatal(scanErr)
		}
		passwd = string(passwdBytes)

		//Modify the tcp configuration file for remote connection
		cfg, err := ini.Load(tcpConfigFilePath)
		if err != nil {
			log.Fatal("Fail to read file:", err)
		}
		cfg.Section("MIR").Key("MIRTCPHost").SetValue(remoteHost)
		cfg.SaveTo(tcpConfigFilePath)

		c, err := ssh.NewSshTerminal(remoteHost, userName, passwd, "min-push-tcp", "/minssh/server", 2222)
		if err != nil {
			log.Fatal("err", err)
		}
		// 进入
		c.EnterTerminal()
		return nil
	}

	if err := minsshClientApp.Run(os.Args); err != nil {
		return
	}
}
