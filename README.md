本项目引用了 https://github.com/percona/go-mysql 包

实际上，本项目只是对 https://github.com/percona/go-mysql 包的 sql 指纹功能,进行打包成可执行文件

#### 使用方法
只限于 cli 下使用
````
./sql-fingerprint 'select * from table where id=1'
#输出：
select * from table where id=?
````

解析多条 sql 指纹，用空格分开：
````
./sql-fingerprint 'select * from table where id=1' 'select * from table where id=2'
#输出
select * from table where id=?
select * from table where id=?
````

### English

This project references the https://github.com/percona/go-mysql package

In fact, this project is just packaging the sql fingerprint function of the https://github.com/percona/go-mysql package into an executable file

#### Instructions
Only available under cli
````
./sql-fingerprint 'select * from table where id=1'
# output:
select * from table where id=?
````

Parse multiple sql fingerprints, separated by spaces:
````
./sql-fingerprint 'select * from table where id=1' 'select * from table where id=2'
# output
select * from table where id=?
select * from table where id=?
````
