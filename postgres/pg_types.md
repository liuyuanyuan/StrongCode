# PostgreSQL数据类型

[Chapter 8. Data Types](https://www.postgresql.org/docs/12/datatype.html)

[51.63. `pg_type`](https://www.postgresql.org/docs/12/catalog-pg-type.html)

| 数据类型         | 别名      | 存储大小 | PreparedStatement setNull          | PreparedStatement setXxx                  |      |
| ---------------- | --------- | -------- | ---------------------------------- | ----------------------------------------- | ---- |
| smallint         | int2      | 2bytes   |                                    |                                           |      |
| int              | int4      | 4bytes   |                                    |                                           |      |
| bigint           | int8，oid | 8bytes   |                                    |                                           |      |
| real             | float4    | 4bytes   |                                    |                                           |      |
| double precision | float8    | 8bytes   |                                    |                                           |      |
| numeric          | decimal？ | 变长     |                                    |                                           |      |
| smallserial      |           | 2bytes   |                                    |                                           |      |
| serial           |           | 4bytes   |                                    |                                           |      |
| bigserial        |           | 8bytes   |                                    |                                           |      |
|                  |           |          |                                    |                                           |      |
| char             |           | 定长     |                                    |                                           |      |
| varchar          | name      | 有限变长 |                                    |                                           |      |
| text             |           | 无限变长 |                                    |                                           |      |
|                  |           |          |                                    |                                           |      |
| date             |           | 4bytes   |                                    |                                           |      |
| time             |           | 8bytes   |                                    |                                           |      |
| timetz           |           | 16bytes  |                                    |                                           |      |
| timestamp        |           | 8bytes   |                                    |                                           |      |
| timestamptz      |           | 8bytes   |                                    |                                           |      |
| interval         |           | 16bytes  |                                    |                                           |      |
|                  |           |          |                                    |                                           |      |
| bytea            |           |          | setNull(i, java.sql.Types.BINARY); | .setBinaryStream(i, (InputStream) value); |      |
|                  |           |          |                                    |                                           |      |
| boolean          |           |          |                                    |                                           |      |
|                  |           |          |                                    |                                           |      |
| money            |           |          |                                    |                                           |      |
|                  |           |          |                                    |                                           |      |
| json             |           |          |                                    |                                           |      |
|                  |           |          |                                    |                                           |      |
| sqlxml           |           |          |                                    |                                           |      |
|                  |           |          |                                    |                                           |      |
|                  |           |          |                                    |                                           |      |

