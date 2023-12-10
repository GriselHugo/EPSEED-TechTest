#!/bin/sh

container_name="ett-db"

if [ "$( docker container inspect -f '{{.State.Running}}' $container_name )" = "true" ]; then
    echo "Importing database..."
    docker compose exec database sh -c 'mariadb -uroot -pthis_is_a_secret_password ett-db < dump.sql'
    echo "Import done"
else
    echo "Container ett-db is not running, please run it first (docker compose up -d database)" >&2
fi
