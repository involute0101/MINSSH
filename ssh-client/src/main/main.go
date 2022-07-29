package main

import (
	"fmt"
	ssh "sshlib"
)

func main() {
	//以下内容应该为用户输入或者配置文件设定
	//这里做example直接使用

	ip := "127.0.0.1"
	identity := "/localhost/operator"
	passwd := "123456"
	network := "min-push-tcp"
	addr := "/localhost/operator"
	port := 2222

	//建立ssh连接，启动终端
	c, err := ssh.NewSshTerminal(ip, identity, passwd, network, addr, port)
	if err != nil {
		fmt.Println("err", err)
	}
	// 进入
	c.EnterTerminal()

	/*ce := func(err error, msg string) {
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

	err = session.Wait()*/
}
