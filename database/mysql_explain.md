# EXPLAIN 调优索引

## EXPLAIN 的使用

EXPLAIN 关键字，可以模拟 MySQL 优化器执行语句，从而很好的分析SQL语句或表结构的性能瓶颈。

几种实用方法

```
EXPLAIN SELECT ...;

EXPLAIN EXTENDED SELECT ...; SHOW WARNING;

EXPLAIN FORMAT=JSON SELECT ...;
```

![preview](img/mysql_explain.png)





