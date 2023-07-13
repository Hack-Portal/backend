#!/bin/sh

set -e

echo "run db migration for"
/app/migrate -path /app/migration -database "$DB_SOURSE" -verbose up

echo "start the app"
exec "$@"