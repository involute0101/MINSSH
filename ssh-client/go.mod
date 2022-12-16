module ssh-client

go 1.18

require (
	golang.org/x/text v0.3.3
	sshlib v0.0.0
)

require (
	MINTCP v0.0.0 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/liyue201/gostl v1.0.1 // indirect
	github.com/mutecomm/go-sqlcipher/v4 v4.4.2 // indirect
	github.com/tylertreat/BoomFilters v0.0.0-20210315201527-1a82519a3e43 // indirect
	github.com/zeebo/xxh3 v1.0.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.21.0 // indirect
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550 // indirect
	golang.org/x/sys v0.0.0-20210510120138-977fb7262007 // indirect
	minlib v0.0.0 // indirect
)

replace minlib v0.0.0 => ../../minlib

replace MINTCP v0.0.0 => ../../MINTCP

replace sshlib v0.0.0 => ../sshlib
