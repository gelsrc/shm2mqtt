@echo off

setlocal

rem set GOPATH=...

set GOOS=linux
set GOARCH=arm
set GOARM=7

go build
