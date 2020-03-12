# SELECT 语法

[TOC]

## JOIN

### LEFT  [OUTER]  JOIN 

以左边表为基础，连接右边表中符合连接条件的结果（无符合结果的用NULL补齐）；

### RIGHT  [OUTER]  JOIN

以右边表为基础，连接左边表中符合连接条件的结果（无符合结果的用NULL补齐）；

### [INNNER]  JOIN

取左边表和右边表中符合连接条件的结果。

### FULL  [OUTER]  JOIN

返回所有联接的行，再为不匹配的左手行（在右边扩展为空）加一行，再为不匹配的右手行（扩展在左边为空）加一行。

### CROSS JOIN

CROSS JOIN 等价于 INNER JOIN ON (TRUE)。这种联接类型只是一种符号上的方便，因为它完成的工作用 FROM 和 WHERE 也能完成。



## GROUP BY grouping_element [, ...] ]

分组表达式中字段的值相同的一行或多行压缩为一行。

```sql
--查询会三种语言技能的人
select name, count(1) as num
from person_skill 
where course in (select lang from skill)
group by name
having count(course)=3::bigint
```



## HAVING condition [, ...] ]

HAVING条件 与 WHERE条件的声明形式是一样的；HAVING条件中的字段必须是GROUP BY的字段或者在查询中使用了聚合函数的字段；

WHERE 过滤的是使用 GROUP BY 之前的独立行，而 HAVING 过滤的是  GROUP BY 创建的组行；

