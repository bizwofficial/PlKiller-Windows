@echo off

color f0
set gopath=%cd%

dev\rsrc -manifest main.exe.manifest -ico plkiller.ico -o main.syso

go clean
go build -v
move plkiller_windows.exe plkiller.exe
echo Success!

pause
color
