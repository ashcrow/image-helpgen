% $FGC/some-componenet(2) Container Image Pages
% Steve Milner
% May 2018

# NAME
$FGC/some-componenet - A key-value store for shared configuration and service discovery.

# DESCRIPTION
Long description documentation comes from the comments at the start of the Dockerfile. Note that spaces at the end of the line one each comment may be needed for formatting. 

If you want to have a second paragraph then add a comment line with nothing on it.


# USAGE
/usr/bin/docker run --cap-add NET_ADMIN -d --cap-add=SYS_ADMIN  -p 4001:4001 -p 7001:7001 -p 2379:2379 -p 2380:2380

# ENVIRONMENT VARIABLES

The image recognizes the following environment variables that you can set
during initialization be passing `-e VAR=VALUE` to the Docker run command.

|     Variable name        | Default |      Description                                           |
| :----------------------- | ------- | ---------------------------------------------------------- |
| `VERSION` | `0.1`   | TODO |
| `RELEASE` | `10`   | TODO |
| `ARCH` | `x86_64`   | TODO |


# SECURITY IMPLICATIONS

## Ports
|     Port Container | Port Host  |       Description             |
| :----------------- | -----------|-------------------------------|
| 4001 | 0 | TODO |
| 7001 | 0 | TODO |
| 2379 | 0 | TODO |
| 2380 | 0 | TODO |



## Volumes
|     Volume Container | Volume Host  |       Description             |
| :----------------- | -----------|-------------------------------|
| /test | TODO | TODO |
| /something | TODO | TODO |
| /else | TODO | TODO |
| /another | TODO | TODO |


## Daemon
This image is expected to be run as a daemon

## Expected Capabilities
- NET_ADMIN
- SYS_ADMIN


# SEE ALSO
http://example.org/asd/asd
