#Long description documentation comes from the comments 
#at the start of the Dockerfile. Note that spaces at the end of the line 
#one each comment may be needed for formatting. 
#
#If you want to have a second paragraph then add a comment line with nothing 
#on it.
FROM registry.fedoraproject.org/fedora:rawhide

# Denotes the version of this dockerfile
ENV VERSION=0.1
# Denotes the release of this dockerfile
ENV RELEASE=10
# Denotes the architecture
ENV ARCH=x86_64

LABEL com.redhat.component="some-compoenent" \
      name="$FGC/some-componenet" \
      version="$VERSION" \
      release="$RELEASE.$DISTTAG" \
      architecture="$ARCH" \
      summary="A key-value store for shared configuration and service discovery." \
      maintainer="Steve Milner <smilner@example.org>" \
      usage="/usr/bin/docker run --cap-add NET_ADMIN -d --cap-add=SYS_ADMIN  -p 4001:4001 -p 7001:7001 -p 2379:2379 -p 2380:2380" \
      url="http://example.org/asd/asd" \
      url="http://example.org/sdgfsdf"

RUN dnf -y --setopt=tsflags=nodocs install etcd hostname && \
    dnf clean all

LABEL INSTALL /usr/bin/docker run --rm \$OPT1 --privileged -v /:/host -e HOST=/host -e NAME=\$NAME -e IMAGE=\$IMAGE \$IMAGE \$OPT2 /usr/bin/install.sh  \$OPT3
LABEL UNINSTALL /usr/bin/docker run --rm \$OPT1 --privileged -v /:/host -e HOST=/host -e NAME=\$NAME -e IMAGE=\$IMAGE \$IMAGE \$OPT2 /usr/bin/uninstall.sh \$OPT3
LABEL RUN /usr/bin/docker run -d \$OPT1 -p 4001:4001 -p 7001:7001 -p 2379:2379 -p 2380:2380 --name \$NAME \$IMAGE \$OPT2 \$OPT3

COPY etcd-env.sh /usr/bin/etcd-env.sh
COPY install.sh  /usr/bin/install.sh
COPY uninstall.sh /usr/bin/uninstall.sh

COPY tmpfiles.template config.json.template service.template manifest.json /exports/

RUN mkdir -p /exports/hostfs/usr/local/bin/ && cp /usr/bin/etcdctl /exports/hostfs/usr/local/bin/etcdctl

RUN mkdir -p /exports/hostfs/etc/etcd && sed -e "/^ETCD_DATA_DIR/d" -e s"|^ETCD_NAME=|#ETCD_NAME=|" < /etc/etcd/etcd.conf > /exports/hostfs/etc/etcd/etcd.conf

#4001->4001 Use for something
EXPOSE 4001
#7001->7001 Used for something else
EXPOSE 7001
#2379->2379 Used for the another thing
EXPOSE 2379
#2380->2380 Used for the last thing
EXPOSE 2380

#/test->/tmp/test A test mount
VOLUME /test
#/something->/tmp/something Another test mount
VOLUME /something
#/else->/tmp/else The else test mount
VOLUME /else
#/another->/var/tmp/another The last test mount
VOLUME /another

CMD ["/usr/bin/etcd-env.sh", "/usr/bin/etcd"]
