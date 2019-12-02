#!/bin/sh

nohup /docker-authz/docker-authz > /dev/null &

if [ "$1" = 'dockerd' ]; then
        uid="$(id -u)"
        if [ "$uid" != '0' ]; then
                echo "Invalid uid"
                exit 1
        elif [ -x '/usr/local/bin/dind' ]; then
		set -- '/usr/local/bin/dind' "$@"
	fi

else
        set -- docker-entrypoint.sh "$@"
fi

exec "$@"
