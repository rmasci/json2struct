compilerFlag=-gcflags=-trimpath=$(shell pwd) -asmflags=-trimpath=$(shell pwd)
goFiles=main.go
ver=$(shell date +"BuildDate:_%Y%m%d-%H:%M")
path=binaries
all: mac linux windows

windows:$(goFiles)
	GOOS=windows GOARCH=386  go build $(compilerFlag) -ldflags="-X main.version=\"$(ver)\""  -o binaries/json2struct-w32.exe $(goFiles)
	cksum binaries/json2struct-w32.exe > binaries/json2struct-w32.exe.cksum
	GOOS=windows GOARCH=amd64  go build $(compilerFlag) -ldflags="-X main.version=\"$(ver)\""  -o binaries/json2struct.exe $(goFiles)
	cksum binaries/json2struct.exe > binaries/json2struct.exe.cksum

linux: $(goFiles)
	GOOS=linux GOARCH=386  go build $(compilerFlag) -ldflags="-X main.version=\"$(ver)\""  -o binaries/json2struct-l32 $(goFiles)
	GOOS=linux GOARCH=amd64  go build $(compilerFlag) -ldflags="-X main.version=\"$(ver)\""  -o binaries/json2struct-l64 $(goFiles)
	cksum binaries/json2struct-l32 > binaries/json2struct-l32.cksum
	cksum binaries/json2struct-l64 > binaries/json2struct-l64.cksum
	#upx binaries/json2struct-l32
	#upx binaries/json2struct-l64

mac: $(goFiles)
	GOOS=darwin GOARCH=amd64  go build $(compilerFlag) -ldflags="-X main.version=\"$(ver)\"" -o binaries/json2struct-mac $(goFiles)
	cksum binaries/json2struct-mac > binaries/json2struct.cksum
	#upx binaries/json2struct-mac

