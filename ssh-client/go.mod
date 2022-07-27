module ssh-client

go 1.18

require (
	MINTCP v0.0.0
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	minlib v0.0.0
	sshlib v0.0.0
)

replace minlib v0.0.0 => ../../minlib

replace MINTCP v0.0.0 => ../../MINTCP

replace sshlib v0.0.0 => ../sshlib
