version: "0.1"
database:
  # 参考[https://gorm.io/docs/connecting_to_the_database.html]
  dsn: "root:123456@tcp(127.0.0.1:3306)/user_db?charset=utf8mb4&parseTime=true&loc=Local"
  # 输入 mysql 或 postgres 或 sqlite 或 sqlserver。参考[https://gorm.io/docs/connecting_to_the_database.html]
  db: "mysql"
  # 输入所需的数据表或留空。您可以输入：orders,users,goods
  tables:
    - "users"
    - "user_profiles"
    - "friends"
    - "friend_requests"
  # 指定输出目录
  outPath: ""
  # 查询代码文件名，默认：gen.go
  outFile: ""
  # 为查询代码生成单元测试
  withUnitTest: false
  # 生成模型代码的包名
  modelPkgName: ""
  # 当字段可为空时生成指针
  fieldNullable: false
  # 生成带有 gorm 索引标签的字段
  fieldWithIndexTag: true
  # 生成带有 gorm 列类型标签的字段
  fieldWithTypeTag: true
