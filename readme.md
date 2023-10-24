go-bindata data
env GOOS="windows" GOARCH="amd" go build -ldflags -H=windowsgui -o fakeserver.exe
