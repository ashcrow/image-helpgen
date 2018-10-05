# image-helpgen

Generates image help and man files based on input. Initial idea from [atomic-wg](https://pagure.io/atomic-wg/issue/354) on [pagure](https://pagure.io/).

[![Build Status](https://travis-ci.org/ashcrow/image-helpgen.svg)](https://travis-ci.org/ashcrow/image-helpgen/)

See [/example](/example) for an example run generating a``Dockerfile`` to ``help.md`` and ``help.1``.

## Build

### Binary
**Note**: ``image-helpgen`` uses [dep](https://github.com/golang/dep/)

```
$ make deps  # Installs dep and pulls deps in
$ make build # Builds binary
$ ls image*
image-helpgen
```

### Container Image

**Note**: You will need the `podman` binary or `docker` binary and service to build images.

```
$ make image # build the container image. Override IMAGE_BUILDER= if you want to use a different build binary than what was found

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

### Running in a container
```
$ sudo podman run --rm -v `pwd`:/data:z image-helpgen:1.0 dockerfile --dockerfile Dockerfile 
```

## How to document

### Long Description
To create a long description for your Dockerfile start the Dockerfile with comments.

*Example*

```
#This is my long description. This is a single line.
#
#Now this is a different line with a newline between.
#This is the last line and by not providing another comment hash it will be the end of the long description.

[...]
```

### Environment Variables (`ENV`)
To document an environment variable:

1. The `ENV` must be on it's own line with only one variable
2. Place a comment above the `ENV` directive with your description

*Example*
```
# Defines the version of this Dockerfile
ENV DOCKERVERSION="1.2.3"
```

### Ports (`EXPOSE`)
To document a port exposures:

1. The `EXPOSE` must be on it's own line with only one port listed
2. Place a comment above the `EXPOSE` line with $CONTAINERPORT->$HOSTPORT $DESCRIPTION

*Example*
```
#8080->80 Provides access to the web server
EXPOSE 8080
```


### Volumes (`VOLUME`)
To document a volume:

1. The `VOLUME` must be on it's own line with only one volume listed
2. Place a commend above the `VOLUME` line with $CONTAINERVOLUME->$HOSTVOLUME $DESCRIPTION

*Example*
```
#/rootfs->/ The hosts root filesystem for the container to modify
VOLUME /rootfs
```