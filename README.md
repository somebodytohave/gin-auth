# gin-blog

运行步骤:
1. 数据库初始化sql文件 [luoliluosuo.sql](https://github.com/sun-wenming/gin-blog/blob/master/luoliluosuo.sql)  数据库名称为  `luoliluosuo`  
在项目的conf配置文件下。数据库 用户 密码默认为 `root`

2. 在项目的根目录运行:`go run main.go`

3. 之后在浏览器上运行`http://localhost:8000/swagger/index.html` 操作 topics下:
![put操作](https://github.com/sun-wenming/gin-blog/blob/master/put.jpg)
或者执行
```
curl -X PUT "http://localhost:8000/api/v1/topics/1" -H "accept: application/json" -H "Content-Type: application/json" -d "{ \"modified_by\": 1, \"name\": \"主题名称\", \"state\": 1}"
```
