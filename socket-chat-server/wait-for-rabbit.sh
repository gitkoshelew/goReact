#!/bin/bash

set -e

host="rabbitmq"
port="5672"
cmd="$@"

echo cmd

>&2 echo "!!!!!!!! Check rabbit for available !!!!!!!!"

until curl http://"$host":"$port"; do
  >&2 echo "rabbit is unavailable - sleeping"
  sleep 1
done

>&2 echo "rabbit is up - executing command"

exec $cmd