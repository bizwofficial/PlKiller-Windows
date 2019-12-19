# PlKiller Windows
Welcome to PlKiller! PlKiller is a remote controlling software which can let you shut other devices down.

## Build
To build PlKiller, you can simply double-click `BuildAll.bat`, or type this in shell:
```
go build
```
But before this, you should change `$GOPATH$` to current directory by double-click run `Set_GOPATH.bat`.

## Run
Before you use PlKiller, you need to create a new configuration file which includes all names of devices you want to kill.Format is like this:
```
Computer 1
Computer 2
...
```
And then run this in shell:
```
plkiller <configuration file name> [-im <process to kill>] [-v]
```

## Specific Process Mode
You can kill a specific process on remote device by adding argument `-im` in calling PlKiller.

## Verbosed Mode
PlKiller is able to receive a optional argument `-v`, additional information will be showed on the screen.
