## 框架的由来
你好，很高兴我的code会在github和你们见面。

在几家游戏公司从事PHP服务端开发之后，发现一套完善的框架对整体的开发（服务端）的时间成本节省了很多。
于是这个框架就由此诞生了，尽管它还不完善。


## 路由规则
* 在`proto/proto.route`的`enum RouteMap`枚举中定义路由
* 请求路由定义奇数,返回路由偶数
* 不指定返回路由默认使用DefaultResponse

## 使用方式
1. `proto`目录下定义proto结构，❕不允许命名为base.proto，它被脚本生成占用
2. `server/model`目录下基于User结构体去定义数据结构
3. 在`server`目录下执行`go run develop.go` 生成对应的操作文件
4. 在`server/action`目录下找到操作方法并开始开发
5. 在`server`目录下`go run main.go` 启动服务

# 数据解析
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

# 更新进度
- [x] 增加反向生成用户结构的proto文件  
- [x] 增加mongoDb支持
- [x] 增加游戏环境配置文件
- [ ] 增加业务内异常返回
- [ ] 增加xlsx表转化为Go配置
- [ ] 增加游戏内玩家注册和初始化数据
- [ ] 增加活动支持
- [ ] 提供跨服支持
- [ ] 提供proto工具调试器





