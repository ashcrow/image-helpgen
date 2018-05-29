% rhscl/mariadb-102-rhel7(2) Container Image Pages
% SoftwareCollections.org
% May 2018

# NAME
rhscl/mariadb-102-rhel7 - MariaDB 10.2 SQL database server

# DESCRIPTION
 MariaDB is a multi-user, multi-threaded SQL database server. The container image provides a containerized packaging of the MariaDB mysqld daemon and client application. The mysqld server daemon accepts connections from clients and provides access to content from MariaDB databases on behalf of the clients.

 This image must forever use UID 27 for mysql user so our volumes are safe in the future. This should *never* change, the last test is there to make sure of that.

 Volumes: 

  /var/lib/mysql/data - Datastore for MySQL 

 Environment: 

  $MYSQL_USER - Database user name 

  $MYSQL_PASSWORD - User&#39;s password 

  $MYSQL_DATABASE - Name of the database to create 

  $MYSQL_ROOT_PASSWORD (Optional) - Password for the &#39;root&#39; MySQL account 


# USAGE
docker run -d -e MYSQL_USER=user -e MYSQL_PASSWORD=pass -e MYSQL_DATABASE=db -p 3306:3306 rhscl/mariadb-102-rhel7

# ENVIRONMENT VARIABLES

The image recognizes the following environment variables that you can set
during initialization by passing `-e VAR=VALUE` to the `docker run` command.

|     Variable name        | Default |      Description                                           |
| :----------------------- | ------- | ---------------------------------------------------------- |
| `MYSQL_VERSION` | `10.2`   | Current MySQL version number |
| `APP_DATA` | `/opt/app-root/src`   | Data directory used by database |
| `CONTAINER_SCRIPTS_PATH` | `/usr/share/container-scripts/mysql`   | Path to mysql script |
| `MYSQL_PREFIX` | `/opt/rh/rh-mariadb102/root/usr`   | Prefix for MySQL usr directory |
| `ENABLED_COLLECTIONS` | `rh-mariadb102`   | Software collections name |
| `BASH_ENV` | `${CONTAINER_SCRIPTS_PATH}/scl_enable`   | Path to enable software collections |
| `ENV` | `${CONTAINER_SCRIPTS_PATH}/scl_enable`   | Path to enable software collections |
| `PROMPT_COMMAND` | `&#34;. ${CONTAINER_SCRIPTS_PATH}/scl_enable&#34;`   | Shell prompt |


# SECURITY IMPLICATIONS
The following sections describe potential security issues related to how the container image was designed to run.

## Ports

Exposed TCP (default) or UDP ports that the container listens on at runtime include the following:

|     Port Container | Port Host  |       Description             |
| :----------------- | -----------|-------------------------------|
| 3306 | 3306 | TCP port used to access mariadb service |



## Volumes

Directories that are mounted from the host system to a mount point inside the container include the following:

|     Volume Container | Volume Host  |       Description             |
| :----------------- | -----------|-------------------------------|
| /var/lib/mysql/data | /var/lib/mysql/data | Database data directory mounted from host |


## Daemon
This image is expected to be run as a daemon


# SEE ALSO
https://access.redhat.com/documentation/en-us/red_hat_software_collections/3/html-single/using_red_hat_software_collections_container_images/#mariadb

