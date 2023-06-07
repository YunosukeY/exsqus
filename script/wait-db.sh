#!/bin/bash

source .env
until mysqladmin ping -h 127.0.0.1 -P 3306 -u ${MYSQL_USER} -p${MYSQL_PASSWORD}; do
  echo 'waiting mysql'
  sleep 2
done

echo "mysql started"
