module ssh-server

go 1.12

require (
	MINTCP v0.0.0
	github.com/anmitsu/go-shlex v0.0.0-20200514113438-38f4b401e2be
	github.com/creack/pty v1.1.18
	golang.org/x/crypto v0.0.0-20210616213533-5ff15b29337e
	minlib v0.0.0
)

replace minlib v0.0.0 => ./../../../minlib

replace MINTCP v0.0.0 => ./../../../MINTCP
