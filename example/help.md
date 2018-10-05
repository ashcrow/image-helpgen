% $FGC/some-componenet(2) Container Image Pages
% Steve Milner
% October 2018

# NAME
$FGC/some-componenet - A key-value store for shared configuration and service discovery.

# DESCRIPTION
Long description documentation comes from the comments at the start of the Dockerfile. Note that spaces at the end of the line one each comment may be needed for formatting. 

If you want to have a second paragraph then add a comment line with nothing on it.


# USAGE
/usr/bin/docker run --cap-add NET_ADMIN -d --cap-add=SYS_ADMIN  -p 4001:4001 -p 7001:7001 -p 2379:2379 -p 2380:2380

# Default Command
/usr/bin/etcd-env.sh /usr/bin/etcd

# ENVIRONMENT VARIABLES

The image recognizes the following environment variables that you can set
during initialization by passing `-e VAR=VALUE` to the `docker run` command.

|     Variable name        | Default |      Description                                           |
| :----------------------- | ------- | ---------------------------------------------------------- |
| ` VERSION` | `0.1`   |  Denotes the version of this dockerfile |
| ` RELEASE` | `10`   |  Denotes the release of this dockerfile |
| ` ARCH` | `x86_64`   |  Denotes the architecture |


# SECURITY IMPLICATIONS
The following sections describe potential security issues related to how the container image was designed to run.

## Ports

Exposed TCP (default) or UDP ports that the container listens on at runtime include the following:

|     Port Container | Port Host  |       Description             |
| :----------------- | -----------|-------------------------------|
| 4001 | 4001 | Use for something |
| 7001 | 7001 | Used for something else |
| 2379 | 2379 | Used for the another thing |
| 2380 | 2380 | Used for the last thing |


## Volumes

Directories that are mounted from the host system to a mount point inside the container include the following:

|     Volume Container | Volume Host  |       Description             |
| :----------------- | -----------|-------------------------------|
| /test | /tmp/test | A test mount |
| /something | /tmp/something | Another test mount |
| /else | /tmp/else | The else test mount |
| /another | /var/tmp/another | The last test mount |


## Daemon
This image is expected to be run as a daemon


## Expected Capabilities

This container needs to open one or more Linux capabilities (see `man capabilities 7`) to the host computer. The following capababilities (added with the \-\-cap\-add option) are expected:


- NET_ADMIN
- SYS_ADMIN


# SEE ALSO

http://example.org/asd/asd
http://example.org/sdgfsdf
