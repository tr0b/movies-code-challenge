#!/bin/sh

set -e

echo "ls -a /app"
ls -a /app
echo "run db migrations"
. /app/.env
/app/migrate -path /app/db/schema -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"
