#!/bin/sh
docker-compose exec db bash -c "chmod 0775 /init-database.sh"
docker-compose exec db bash -c "./init-database.sh"