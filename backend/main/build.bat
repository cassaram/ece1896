@echo off

SET GOOS=linux
SET GOARCH=arm
go build -o "bin\linux_arm"

SET GOOS=linux
SET GOARCH=amd64
go build -o "bin\linux_amd64"

SET GOOS=windows
SET GOARCH=amd64
go build -o "bin\windows_amd64.exe"