# ws游戏服务器脚手架
作为学习和练习维护
# 使用方式
## 路由规则
* 在`game/proto.route`的`enum RouteMap`枚举中定义路由
* 请求路由定义奇数,返回路由偶数
* 不指定返回路由默认使用DefaultResponse

## 使用方式
1. 先定义路由和设计proto结构
2. 执行命令`go generate ./develop/createPb.go`生成go的proto文件
3. 执行命令`go generate ./develop/createAction.go`生成对应的操作方法文件

# 数据解析
目前这块还没做，目前粗糙设计为
1.客户端发送的二进制数据分为四个段，分别为
* `int32 包体长度`,
* `int32 请求自增ID`,
* `int32 请求的路由ID`
* `string proto数据串`

2.服务端返回的二进制数据分为四个段，分别为
* `int32 包体长度`,
* `int32 客户端请求的路由ID`,
* `int32 错误标识`
* `string proto数据串`

# 数据库支持
还没做，正在考虑选型redis，mysql，mongoDb




