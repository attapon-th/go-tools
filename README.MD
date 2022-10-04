# GO Tools

## go script 
```shell
go install github.com/attapon-th/go-tools/gosc@latest
```

```raw
$ gosc -h

## output
script with go same shell-script

Usage:
  gosc [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  replace     Replace string from STDIN
  sed         Replace string from STDIN with MatchRegexp
  upver       Up version in file 

Flags:
  -h, --help     help for gosc
  -t, --toggle   Help message for toggle

Use "gosc [command] --help" for more information about a command.
```