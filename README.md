# goapi
API 模拟器

使用方法：

go build api.go

api -p 端口  -s 延时秒数

将需要模拟的接口放在api可执行文件所在的目录。以.json作为文件名为后缀

然后前端发请求:
例如  http://127.0.0.1:8099/api/user/list.json
