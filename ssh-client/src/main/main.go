package main

import (
	"sshlib"
	//"MINTCP/socket"
	"log"
	//"bufio"
	"fmt"
	//"minlib/common"
	//"minlib/security"
	"net"
	"os"
	//"ssh-client/src/main/process"
	//_struct "ssh-client/src/main/struct"
	//"strings"
	"time"
)

func main() {
	ce := func(err error, msg string) {
		if err != nil {
			log.Fatalf("%s error: %v", msg, err)
		}
	}
	sshConfig := &ssh.ClientConfig{
		User: "guozhan",
		Auth: []ssh.AuthMethod{
			ssh.Password("secret"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		ClientVersion:   "",
		Timeout:         10 * time.Second,
	}
	sshClient, err := ssh.Dial("min-push-tcp", "/localhost/operator", sshConfig)
	if err != nil {
		log.Fatal("创建ssh client 失败", err)
	}
	defer sshClient.Close()

	//创建ssh-session
	session, err := sshClient.NewSession()
	if err != nil {
		log.Fatal("创建ssh session 失败", err)
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	//下面的modes和RequestPty是为了防止出现交换机的时候错误
	modes := ssh.TerminalModes{
		ssh.ECHO:          0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	err = session.RequestPty("xterm", 25, 80, modes)

	//这里的session.Shell可以执行多条命令，但是session.Run只支持执行单条命令
	err = session.Shell()
	ce(err, "start shell")

	err = session.Wait()
	////执行远程命令
	//combo, err := session.CombinedOutput("whoami")
	//if err != nil {
	//	log.Fatal("远程执行cmd 失败", err.Error())
	//}
	//log.Println("命令输出:", string(combo))
}

func handleConnectionReader(c net.Conn) {
	//sshReaderWriter := _struct.NewSshReaderWriter(c)
	count := 0
	result := ""
	for {
		buf := make([]byte, 1024)
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			os.Exit(1)
		}
		if string(buf[:n]) == "登录成功！" {
			fmt.Printf("%s", buf[:n])
			fmt.Print("\n>")
			count = 0
		} else {
			if count == 0 {
				fmt.Printf("%s", buf[:n])
				count++
			} else if count == 1 {
				result = result + "[" + string(buf[:n-1])
				count++
			} else {
				result = result + " " + string(buf[:n-1])
				count = 0
				fmt.Print(result + "]>")
				result = ""
			}
		}
	}
}
