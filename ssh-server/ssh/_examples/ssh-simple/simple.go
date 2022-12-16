package main

import (
	"bufio"
	"fmt"
	"github.com/amoghe/go-crypt"
	"github.com/creack/pty"
	"io"
	"log"
	"os"
	"os/exec"
	ssh "ssh-server"
	"strings"
	"syscall"
	"unsafe"
)

func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}

func getUserPasswd(userName string) (string, error) {
	file, err := os.Open("/etc/shadow")
	if err != nil {
		return "", err
	}
	defer file.Close()
	buf := bufio.NewScanner(file)
	userInfo := ""
	for {
		if !buf.Scan() {
			return "", nil
		}
		userInfo = buf.Text()
		if strings.HasPrefix(userInfo, userName) { //找到linux用户信息
			passwdInfo := userInfo[strings.Index(userInfo, userName)+len(userName)+1 : len(userInfo)]
			return passwdInfo[0:strings.Index(passwdInfo, ":")], nil
		}
	}
	return "", nil
}

func main() {
	ssh.Handle(func(s ssh.Session) {

		//cmd := exec.Command("/bin/sh", "/home/guozhan/login.sh")
		cmd := exec.Command("/bin/bash")
		ptyReq, winCh, isPty := s.Pty()
		if isPty {
			cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
			f, err := pty.Start(cmd)
			if err != nil {
				panic(err)
			}
			go func() {
				for win := range winCh {
					setWinsize(f, win.Width, win.Height)
				}
			}()
			go func() {
				io.Copy(f, s) // stdin
			}()
			io.Copy(s, f) // stdout
			cmd.Wait()
		} else {
			io.WriteString(s, "No PTY requested.\n")
			s.Exit(1)
		}

		//io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
	})

	//以下内容应为用户输入或配置文件设定
	identityName := "/minssh/server"
	unlockPasswd := "123456"
	network := "min-push-tcp"
	stackAddr := "/tmp/mir-tcp-message-channel-stack.sock"
	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil, identityName, unlockPasswd, network, stackAddr,
		ssh.PasswordAuth(func(ctx ssh.Context, userPwd string) bool {
			log.Println("The currently logged in user is : ", ctx.User())
			systemPasswd, err := getUserPasswd(ctx.User())
			if err != nil || systemPasswd == "" {
				log.Fatal("something wrong in finding user password , maybe user not exist : ", err)
				return false
			}
			if strings.HasPrefix(systemPasswd, "$") {
				salt := systemPasswd[0:strings.LastIndex(systemPasswd, "$")] //其实这里是加密算法和盐值
				encryptedPwd, err := crypt.Crypt(userPwd, salt)
				if err != nil {
					fmt.Errorf("Encryption process error : ", err)
					return false
				}
				//log.Println("encryptedPwd : ", encryptedPwd)
				if systemPasswd == encryptedPwd {
					log.Println("login successfully!")
					return true
				} else {
					log.Fatal("wrong password!")
					return false
				}
			}
			log.Fatal("Login is not supported for this user type at the moment...Please contact the administrator or change the user login")
			return false
		})))
}
