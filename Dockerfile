# Generates image help and man files based on Dockerfile contents.
FROM alpine:latest

ENV VERSION=1.0 RELEASE=1 ARCH=x86_64
LABEL  name="image-helpgen" \
      version="$VERSION" \
      release="$RELEASE.$DISTTAG" \
      architecture="$ARCH" \
      summary="Generates image help and man files based on Dockerfile contents." \
      maintainer="Steve Milner <smilner@redhat.com>" \
      usage="/usr/bin/docker run --rm -v `pwd`:/data:z image-helpgen:latest $OPTIONS " \
      url="http://github.com/ashcrow/image-helpgen"

VOLUME /data

COPY image-helpgen /usr/bin/image-helpgen
COPY image-files/run-in-container.sh /usr/bin/run-in-container.sh
COPY image-files/image-helpgen.1 /help.1

ENTRYPOINT [ "/usr/bin/run-in-container.sh" ]