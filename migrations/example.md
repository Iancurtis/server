migrate -url mysql://root@tcp\(127.0.0.1:3306\)/cms -path ./migrations create 20161015_CreateTablePagesAndInsertARecord

migrate writes a table named `schema_migrations` to the db to mark the version.
we could use `up` to apply all migrations and use `down` to roll back
`reset` === `down` then `up`
`redo` === `migrate -1` then `migrate +1`
`goto v` 