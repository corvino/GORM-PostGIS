create schema if not exists postgis;
grant usage on schema postgis to public;
create extension if not exists postgis schema postgis;
alter database "gorm-postgis" set search_path=public,postgis,contrib;
