compilerFlag=-gcflags=-trimpath=$(shell pwd) -asmflags=-trimpath=$(shell pwd)
goFiles=main.go
ver=$(shell date +"BuildDate:_%Y%m%d-%H:%M")
all: mac linux windows

windows:$(goFiles)
	GOOS=windows GOARCH=386  go build $(compilerFlag) -ldflags="-X main.version=\"$(ver)\""  -o ../json2struct-release/json2struct-w32.exe $(goFiles)
	cksum ../json2struct-release/json2struct-w32.exe > ../json2struct-release/json2struct-w32.exe.cksum
	GOOS=windows GOARCH=amd64  go build $(compilerFlag) -ldflags="-X main.version=\"$(ver)\""  -o ../json2struct-release/json2struct.exe $(goFiles)
	cksum ../json2struct-release/json2struct.exe > ../json2struct-release/json2struct.exe.cksum

linux: $(goFiles)
	GOOS=linux GOARCH=386  go build $(compilerFlag) -ldflags="-X main.version=\"$(ver)\""  -o ../json2struct-release/json2struct-l32 $(goFiles)
	GOOS=linux GOARCH=amd64  go build $(compilerFlag) -ldflags="-X main.version=\"$(ver)\""  -o ../json2struct-release/json2struct-l64 $(goFiles)
	cksum ../json2struct-release/json2struct-l32 > ../json2struct-release/json2struct-l32.cksum
	cksum ../json2struct-release/json2struct-l64 > ../json2struct-release/json2struct-l64.cksum
	upx ../json2struct-release/json2struct-l32
	upx ../json2struct-release/json2struct-l64

mac: $(goFiles)
	GOOS=darwin GOARCH=amd64  go build $(compilerFlag) -ldflags="-X main.version=\"$(ver)\"" -o ../json2struct-release/json2struct-mac $(goFiles)
	cksum ../json2struct-release/json2struct-mac > ../json2struct-release/json2struct.cksum
	upx ../json2struct-release/json2struct-mac

