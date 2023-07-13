#!/bin/sh

set -e

echo "run db migration for"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURSE" -verbose up

echo "start the app"
exec "$@"