# xccache-sweeper

## A Work In Progress

xccache-sweeper is still in active development.

## Instllation

### go get

```shell
$ go get github.com/yutailang0119/xccache-sweeper
```

### From Source

```shell
$ git clone https://github.com/yutailang0119/xccache-sweeper
$ cd xccache-sweeper
$ make install
```

## Usege

```shell
NAME:
   xccache-sweeper - Sweep Xcode caches

USAGE:
   xccache-sweeper [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
     archives       Sweep Archives. Defaults is /Users/user/Library/Developer/Xcode/Archives
     deriveddata    Sweep DerivedData. Defaults is /Users/user/Library/Developer/Xcode/DerivedData
     caches         Sweep Archives and DerivedData.
     devicesupport  Sweep Device Support. ~/Library/Developer/Xcode/*DeviceSupport
     help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

