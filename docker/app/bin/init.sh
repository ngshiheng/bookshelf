#!/usr/bin/env bash
echo 'Runing migrations...'
/bookshelf/migrate up > /dev/null 2>&1 &

echo 'Deleting mysql-client...'
apk del mysql-client

echo 'Start application...'
/bookshelf/app