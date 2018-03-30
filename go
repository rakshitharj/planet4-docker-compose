#!/usr/bin/env bash
set -eax

APP_HOSTNAME=${APP_HOSTNAME:-test.planet4.dev}

SCALE_OPENRESTY=${SCALE_OPENRESTY:-1}
SCALE_APP=${SCALE_APP:-1}

[[ -f "acme.json" ]] && chmod 600 acme.json

docker-compose -f "${DOCKER_COMPOSE_FILE:-docker-compose.yml}" up -d \
  --scale openresty=$SCALE_OPENRESTY \
  --scale php-fpm=$SCALE_APP

[[ "$1" = "-f" ]] && docker-compose logs -f ${2:-php-fpm}