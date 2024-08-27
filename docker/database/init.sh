#!/usr/bin/env sh
set -e

psql -v ON_ERROR_STOP=1 --username "${POSTGRES_USER}" --dbname "${POSTGRES_DB}" <<-EOSQL
    CREATE DATABASE tenant1db;
    CREATE DATABASE tenant2db;
EOSQL
