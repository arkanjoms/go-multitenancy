#!/usr/bin/env sh
set -e

psql -v ON_ERROR_STOP=1 --username "${POSTGRES_USER}" --dbname "${POSTGRES_DB}" <<-EOSQL
    CREATE USER tenant1 WITH ENCRYPTED PASSWORD 'tenant1';
    CREATE DATABASE tenant1db;
    GRANT ALL PRIVILEGES ON DATABASE tenant1db TO tenant1;
    CREATE USER tenant2 WITH ENCRYPTED PASSWORD 'tenant2';
    CREATE DATABASE tenant2db;
    GRANT ALL PRIVILEGES ON DATABASE tenant2db TO tenant2;
EOSQL
