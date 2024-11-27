# api-gen
### 根据数据库表生成api文件


- **-dir**

生成的api文件存放路径（默认为`./desc`）
- **-home**

生成代码使用的模板路径（默认使用内置模板）
- **-prefix**

api文件路由前缀（默认为`/prefix`）
- **-table***

要生成的数据表名
- **-url***

数据库连接url，形如`root:password@tcp(127.0.0.1:3306)/database`
- **-service**

api服务名（默认为数据库名的小驼峰）
- **-y**

是否强制覆盖（即使开启强制覆盖，也不会覆盖service.api和common.api文件） 

### 根据数据库表生成proto文件


- **-dir**

生成的api文件存放路径（默认为`./pb`）
- **-home**

生成代码使用的模板路径（默认使用内置模板）
- **-package**

proto文件所在的包名（默认为数据库名的小驼峰）
- **-table***

要生成的数据表名
- **-url***

数据库连接url，形如`root:password@tcp(127.0.0.1:3306)/database`
- **-y**

是否强制覆盖

### 注意事项

- 为了代码安全，生成的api和proto文件默认不会覆盖同名文件  
- 生成的api和proto文件中的注释来源有数据库的comment，所以我们推荐任何字段和表都要加comment
