# image-helpgen

Generates image help and man files based on input. Initial idea from [atomic-wg](https://pagure.io/atomic-wg/issue/354) on [pagure](https://pagure.io/).

[![Build Status](https://travis-ci.org/ashcrow/image-helpgen.svg)](https://travis-ci.org/ashcrow/image-helpgen/)

See [/example](/example) for an example run generating a``Dockerfile`` to ``help.md`` and ``help.1``.

## Build
**Note**: ``image-helpgen`` uses [govendor](https://github.com/kardianos/govendor)

```
$ make deps  # Installs govendor and pulls deps in
$ make build # Builds binary
$ ls image*
image-helpgen
```

## Install
**Note**: To see a list of variables which can be overridden run ```make help```

```
$ PREFIX=/install/root make install  # clean, build, and installs with a PREFIX
rm -f image-helpgen
go build -ldflags '-X main.version=0.0.0 -X main.commitHash=a2b5dd271de018879fbd75bd321e85e65d4d722b -X main.buildTime=1512415719 -X main.defaultTemplate=/install/root/etc/image-helpgen/template.tpl' -o image-helpgen main.go
strip image-helpgen
install -d /install/root/etc/image-helpgen/
install --mode 644 template.tpl /install/root/etc/image-helpgen/template.tpl
install -d /install/root/usr/bin
install --mode 755 image-helpgen /install/root/usr/bin/image-helpgen
```

## Example

### Commands
```
Usage: ./image-helpgen <command> [args]
Commands:
  guide: Asks for input and builds markdown and man output
  dockerfile: Parses a Dockerfile and generates a markdown template
  man: Generate man page off of a previously filled out markdown template
```

### From a Dockerfile
```
$ ./image-helpgen dockerfile -dockerfile /path/to/my/Dockerfile
$ ls help*
help.md
$ <edit help.md>
$ ./image-helpgen man
$ ls help*
help.1 help.md
```

### Guide with a specific template and output files
```
$ ./image-helpgen guide -template template.tpl -basename myFile
Image Name: MyImage
...
$ ls myFile*
myFile.1 myFile.md
```
