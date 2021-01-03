#!/usr/bin/env bash
#wait for the MySQL Server to come up
#sleep 90s

#run the setup script to create the DB and the schema in the DB
mysql --defaults-extra-file=/etc/mysql/conf.d/my.cnf < "/docker-entrypoint-initdb.d/dump.sql"
#mysql --defaults-extra-file=/etc/mysql/conf.d/my.cnf -h localhost mydb < "/var/lib/mysql/mydb_user.sql"