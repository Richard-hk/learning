### Union运算符
1）创建数据

```sql
create table union_t1 (name varchar(20), age int);

insert into union_t1 values
("zhangsan", 20),
("lisi", 21);

create table union_t1_back like union_t1;

insert into union_t1_back select * from union_t1;

update union_t1 set age =22 where name = "lisi"
```

2）union （会去掉重复的数据）

```sql
mysql> select * from union_t1 
union 
select * from union_t1_back;
+----------+------+
| name     | age  |
+----------+------+
| zhangsan |   20 |
| lisi     |   22 |
| lisi     |   21 |
+----------+------+
```

3）union all（会保留重复的数据）

```sql
mysql> select * from union_t1 
union all  
select * from union_t1_back;
+----------+------+
| name     | age  |
+----------+------+
| zhangsan |   20 |
| lisi     |   22 |
| zhangsan |   20 |
| lisi     |   21 |
+----------+------+
```

**项目实战**

比较**备份表**和**原表**的数据哪些数据不一致

```sql
mysql> select * from ( 
select * from union_t1 
union all 
select * from union_t1_back
) t 
group by name, age 
having count(*) =1;
+------+------+
| name | age  |
+------+------+
| lisi |   21 |
| lisi |   22 |
+------+------+
```
