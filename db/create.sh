#! /usr/bin/env bash

cd "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

psql postgres < create-db.sql
psql gorm-postgis < create-postgis.sql
