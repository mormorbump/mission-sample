#!/usr/bin/env sh

set -eu

echo "Waiting for MySQL to start..."
while ! mysqladmin ping -h"${DATABASE_HOST}" --silent; do
    sleep 1
done

echo "Creating database if it doesn't exist..."
export MYSQL_PWD=$DATABASE_PASSWORD
mysql -h "$DATABASE_HOST" -u "$DATABASE_USER" -e "CREATE DATABASE IF NOT EXISTS ${DATABASE_NAME};"
unset MYSQL_PWD
DB_SOURCE_NAME="mysql://${DATABASE_USER}:${DATABASE_PASSWORD}@tcp(${DATABASE_HOST}:3306)/${DATABASE_NAME}"

migrate --path /app/migrations --database "${DB_SOURCE_NAME}" -verbose up