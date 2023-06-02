# sql-fingerprint
这是一个 sql 指纹提取器，只提供输出 sql 指纹功能
## 编译
```
git clone git@github.com:anjingjingan/sql-fingerprint.git
cd sql-fingerprint
go mod download
go build .
```
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
#sql-fingerprint
This is a sql fingerprint extractor that only provides output sql fingerprint function
## compile
```
git clone git@github.com:anjingjingan/sql-fingerprint.git
cd sql-fingerprint
go mod download
go build .
```
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
