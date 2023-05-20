编译

set GOARCH=arm64
set GOOS=linux
go build -trimpath -o bin/check.out main.go
