module ssh-client

go 1.18

require (
	github.com/howeyc/gopass v0.0.0-20210920133722-c8aef6fb66ef
	github.com/urfave/cli/v2 v2.3.0
	golang.org/x/text v0.3.3
	gopkg.in/ini.v1 v1.62.0
	sshlib v0.0.0
)

require (
	MINTCP v0.0.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/liyue201/gostl v1.0.1 // indirect
	github.com/mutecomm/go-sqlcipher/v4 v4.4.2 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
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
