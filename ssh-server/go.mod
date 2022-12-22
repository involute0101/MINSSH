module ssh-server

go 1.12

require (
	MINTCP v0.0.0
	github.com/SunnyQjm/daemon v1.0.2
	github.com/amoghe/go-crypt v0.0.0-20220222110647-20eada5f5964
	github.com/anmitsu/go-shlex v0.0.0-20200514113438-38f4b401e2be
	github.com/creack/pty v1.1.18
	github.com/gliderlabs/ssh v0.3.5
	golang.org/x/crypto v0.0.0-20220826181053-bd7e27e6170d
	minlib v0.0.0
)

replace minlib v0.0.0 => ../../minlib

replace MINTCP v0.0.0 => ../../MINTCP
