# gin-auth

### 主要实现功能

更多相关项目流程参考[流程图](https://github.com/sun-wenming/gin-auth/blob/master/user.xmind)

1. 建立通用用户表的增删改查
   - user 用户储存用户的相关信息，不包含敏感信息。(用户名,密码等)
   - user_login 管理用户密码的登录操作,最终关联到 `user` 表查询数据
   - user_oauth 管理用户第三方登录的信息,类似`user_login` (待完善)
2. 注册功能
   - 账号密码 (加密)
   - 第三方注册(待完善)
3. 登录功能
4. JWt 认证功能

### 运行步骤:

1. 将 [Sql 文件 user.sql](https://github.com/sun-wenming/gin-auth/blob/master/user.sql) 导入数据库 `user` 中<br>
   在项目的 conf 配置下。数据库用户、密码默认为 `root` 端口号：`3306`

2. 在项目的根目录运行: `swag init` 初始化文档 And 运行 `go run main.go` 跑起来程序

3. 之后在浏览器上运行 http://localhost:8000/swagger/index.html 开始测试功能

### 待完善

- 第三方登录与注册

### 今后打算

> 打算将此内容成为一个基于 Gin 的 web 框架相关系列 **gin-xxx**<br>
> 涉及到知识点如下

1. [用户操作-登录/注册/认证 gin-auth](https://github.com/sun-wenming/gin-auth)
2. 暂定

### 涉及到的知识点(框架)

> 部分功能会展开介绍,如遇问题 Issues

- web 框架 [gin](https://github.com/gin-gonic/gin)
- orm 数据库操作 [gorm](https://github.com/jinzhu/gorm)
- jwt 认证 [jwt-go](https://github.com/dgrijalva/jwt-go)
- 验证 [validator](https://github.com/go-playground/validator)
- 缓存 [redigo](https://github.com/gomodule/redigo)
- 自动生成开发文档
  - [gin-swagger](https://github.com/swaggo/gin-swagger)
  - [swag](https://github.com/swaggo/swag)

### 参考

1. 整体架构入门参考于 [go-gin-example](https://github.com/EDDYCJY/go-gin-example/blob/master/README_ZH.md)
2. 建表参考于 [廖老师的文章](https://www.liaoxuefeng.com/article/001437480923144e567335658cc4015b38a595bb006aa51000)

### 结语

我也属于刚入行`Golang`新星.如发现重大问题请直接指出. 小弟积极吸取. **Thx.**
