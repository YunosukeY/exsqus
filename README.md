# exsqus - Explain Slow Queries

[![build](https://github.com/YunosukeY/exsqus/actions/workflows/build.yaml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/exsqus/actions/workflows/build.yaml)
[![ci](https://github.com/YunosukeY/exsqus/actions/workflows/ci.yaml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/exsqus/actions/workflows/ci.yaml)
[![e2e](https://github.com/YunosukeY/exsqus/actions/workflows/e2e.yaml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/exsqus/actions/workflows/e2e.yaml)
[![lint](https://github.com/YunosukeY/exsqus/actions/workflows/lint.yml/badge.svg?branch=master&event=push)](https://github.com/YunosukeY/exsqus/actions/workflows/lint.yml)

Monitors a slow query log file and automatically shows their execution plans.

## QuickStart

```
docker pull kimitsu/exsqus
docker run --rm -e MYSQL_HOST={MYSQL_HOST} -e MYSQL_DATABASE={MYSQL_DATABASE} -e MYSQL_USER={MYSQL_USER} -e MYSQL_PASSWORD={MYSQL_PASSWORD} -v {PATH_TO_SLOW_QUERY_LOG_DIR}:/tmp kimitsu/exsqus
```

## Usage

Prerequisites (MySQL)

1. Enable the slow query log setting.
2. Mount the log directory path.

Configurations (exsqus)

1. Mount the log directory path.
2. Set environment variables about the path and MySQL connection configs.

### Environment Variables

#### `MYSQL_HOST` (required)

The host name of the MySQL server.

#### `MYSQL_PORT` (optional)

The port number of the MySQL server.
The default value is `3306`.

#### `MYSQL_PROTOCOL` (optional)

The protocol for connecting to the MySQL server
The default value is `tcp`.

#### `MYSQL_DATABASE` (required)

The database name.

#### `MYSQL_USER` (required)

The user name in the MySQL server.

#### `MYSQL_PASSWORD` (required)

The password in the MySQL server.

#### `LOG_FILE_PATH` (optional)

The path to the slow query log file in the exsqus container.
The default value is `/tmp/slow.log`.
If you want to specify multiple paths, separate them with a comma.

## Examples

### Docker

See the [docker-compose.yml](./docker-compose.yml).

### Kubernetes

An example of sidecar pattern: [manifest](./kind/deployment.yaml).
