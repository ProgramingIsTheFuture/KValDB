# KValDB

Simple key value database that use json files to store the database, the key and the respective value.

This simple database have two gRPC methods:
	- Get
	- Set

The `GET` method will receive the database name and the key. With this, the server will return the respective key and value.

The `SET` method will receive the database name, with this it will create or open the file <database-name>.json and store the key and value you insert.


