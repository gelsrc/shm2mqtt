@echo off

setlocal

set GOPATH=W:\Gelicon\Source5\Привод-нефтесервис\Автомат\goshm

set GOOS=linux
set GOARCH=arm
set GOARM=7

go build
