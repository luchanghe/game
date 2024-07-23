# ws服务端基本脚手架
作为学习和练习维护，更新会很慢。
# 使用方式
## 路由规则
* 在`game/proto.route`的`enum RouteMap`枚举中定义路由
* 请求路由定义奇数,返回路由偶数
* 不指定返回路由默认使用DefaultResponse

## 使用方式
1. `game/proto`目录下定义proto结构，❕不允许命名为base.proto，它被脚本生成占用
2. `game/model`目录下基于User结构体去定义数据结构
3. 执行`go run develop.go` 将User结构体生成为`base.proto`文件并生成go的proto文件存储在`game/pb`目录下，并生成对应的操作方法文件

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

# 未来计划
- [x] 增加反向生成用户结构pb的脚本  
- [ ] 增加mongoDb的数据库支持
- [ ] 增加活动支持
- [ ] 增加单区服配置支持
- [ ] 提供网关支持
- [ ] 提供跨服支持
- [ ] 提供proto工具调试器





