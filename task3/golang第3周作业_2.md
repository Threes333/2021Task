#### 第3周任务

* 数据库学习,了解数据库的基本概念,熟练使用SQL语句,掌握MySQL的安装,使用。

	tips:初学数据库可以看看**<<SQL必知必会>>**,安装可以百度

* golang操作数据库,学习golang的sql包,使用golangCRUD

	**文档**:https://studygolang.com/static/pkgdoc/pkg/database_sql.htm

#### 第3周作业

1.在linux或windows上安装mysql,使用可视化工具操作(安装好就行)。

2.使用SQL语句在数据库中建表(**交建表SQL语句**)

**仓库信息表**: 

*字段*: 仓库编码(varchar),仓库容量(int)

*其他*: 仓库编码为主键

建表SQL语句:

```mysql
create table `store` (
     `id` varchar(20) primary key,
     `capacity` int
     );
```

**服装表**:

*字段*: 服装编码(varchar),服装尺码(varchar),销售价格(int),服装类型(varchar)

*其他*:  服装编码为主键

建表SQL语句:

```mysql
create table `clothings`(
    `id` varchar(20) primary key,
    `size` varchar(20),
    `price` int,
    `style` varchar(20)
     );
```

**供应商表**:

*字段*: 供应商编码(varchar),供应商名称(varchar)

*其他*:  供应商编码为主键

建表SQL语句:

```mysql
create table `suppliers` (
    `id` varchar(20) primary key,
    `name` varchar(20)
    );
```

**供应情况表**:

*字段*: 服装编码(varchar),供应商编码(int),服装质量等级(varchar)

*其他*:  服装编码、供应商编码为主键

建表SQL语句:

```mysql
create table `condition`(
     `id` varchar(20),
     `supplier_id` int,
     `quality` varchar(20),
     primary key(`id`,`supplier_id`)
     );
```

3.建好表后,使用golang的`database/sql`包连接数据库并执行下面几个SQL(**交程序源码**)

查询:

​	(1)查询服装尺码为'S'且销售价格在100以下的服装信息。

​	(2)查询仓库容量最大的仓库信息。 

（3）查询A类服装的库存总量。 

​	(4) 查询服装编码以‘A’开始开头的服装。

（5）查询服装质量等级有不合格的供应商信息。

更新:

​	(1)把服装尺寸为'S'的服装的销售价格均在原来基础上提高10%。

删除:

​	(1)删除所有服装质量等级不合格的供应情况。

插入:

​	(1)	向每张表插入一条记录。



#### 学习书籍:

* **《SQL必知必会》**

* **《MySQL技术内幕:InnoDB存储引擎》姜承尧 著**
* **《高性能MySQL》第3版**



#### 作业提交方式

在截止日期之前将代码上传到GitHub的代码仓库中（截止日期为8月8日）