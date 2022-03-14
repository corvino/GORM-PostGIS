#! /usr/bin/env bash

psql postgres <<< '
  drop database if exists "gorm-postgis";
  drop role if exists "gorm-postgis";'
