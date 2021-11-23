This repo demonstrates an issue with Upserts for SQL Server in Gorm.

To test set `SQLSERVER_DSN` env variable to a valid DSN for SQL Server.

Expectation: If I run this example the `Data` column for the `TestModel` with Matcher `foobar` should be updated to `foobaz`.

What happens: It throws `mssql: Cannot insert duplicate key row in object 'dbo.test_models' with unique index 'idx_test_models_matcher'. The duplicate key value is (foobar).` for SQL Server.

Observation: The expectation holds for the sqlite database, so the implementation seems correct.