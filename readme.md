# GORM+PostGIS

A sample project to demonstrate using the Go ORM [GORM](https://gorm.io) to interact with [PostGIS](https://postgis.net).

Use  [GORM customized data types](https://gorm.io/docs/data_types.html) to read and write geography/geometry columns, which can be specified struct tag. Values are turned into [geos](https://github.com/paulsmith/gogeos) types.

# Example

`create.sh` and `destroy.sh` scripts in `db` can be used to create and destroy the database running on localhost.

The go program writes some sample cities to the database, and then queries for cities around Seattle.

# References

- [gogeos](https://github.com/paulsmith/gogeos)
- [GORM](https://gorm.io)
- [PostGIS](https://postgis.net)
- [Stackoverflow: Inserting and selecting PostGIS Geometry with Gorm](https://stackoverflow.com/questions/54602557/inserting-and-selecting-postgis-geometry-with-gorm)
- [Gorm-PostGIS](https://github.com/OscarStack/Gorm-PostGIS)
