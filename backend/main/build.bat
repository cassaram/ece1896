@echo off

SET GOOS=linux
SET GOARCH=arm
SET GOARM=7
go build -o "bin\sdesign-backend-linux_arm"

SET GOOS=linux
SET GOARCH=amd64
go build -o "bin\sdesign-backend-linux_amd64"

SET GOOS=windows
SET GOARCH=amd64
go build -o "bin\sdesign-backend-windows_amd64.exe"