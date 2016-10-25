## Generate `Lorem ipsum` in your clipboard in no time

## Build

```sh
$> git clone git@github.com:dwi2/loremipsum.git # in your golang workspace
$> go get .
$> go install
```

## Run

```sh
$> /path/to/go/bin/loremipsum # copy full 'Lorem ipsum' in clipboard
$> /path/to/go/bin/loremipsum s # copy first sentence of 'Lorem ipsum' in clipboard
$> /path/to/go/bin/loremipsum t # copy only 'Lorem ipsum' in clipboard
```


## Synopsis

NAME:
   loremipsum - Generate `Lorem ipsum` in your clipboard in no time

USAGE:
   loremipsum [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     short, s  Generate first sentence of `Lorem ipsum` in clipboard
     tiny, t   Generate only `Lorem ipsum` in clipboard
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

