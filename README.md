go-bindata data
$env:GOARCH="386"
go build -ldflags -H=windowsgui -o fakeserver.exe
