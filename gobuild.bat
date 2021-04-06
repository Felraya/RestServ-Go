@echo off
rem : RUN "go clean" and "go build" 
rem : the resulting executables are in the current folder
rem : use "go help build" command to see all available options
@echo on
go clean -cache ./...
go build -o . -v -x ./...
@echo off
if %ERRORLEVEL% GEQ 1 echo !!!!! ERROR !!!!!
