# TikTokLite
<center><img src="https://img-blog.csdnimg.cn/cc03d16ceea3494e8b38ce0f4a5eb0f6.png" alt="image-20221027202423203" width="65%" height="65%"  ></center>

***

**极简抖音**

[探索本项目相关文档](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145)

## 荣誉展示

## 上手指南

### 启动服务

将`config.yaml`中所有host改为本机地址后输入

```bash
docker-compose up
```

即可通过docker快速启动部署服务及相关依赖服务

### 相关环境

- **golang**>= 1.18
- **mysql**>=8.0：数据库
- **redis**>=7.0.0：缓存
- **minio**：对象存储
- **ffmpeg**：获取视频封面

## 技术选型

<img src="https://img-blog.csdnimg.cn/0d5cabef362d4f71a5051b44596745c1.png" width="50%" height="50%" >

## 实现功能

|    功能    |                             说明                             |
| :--------: | :----------------------------------------------------------: |
|  基础功能  |      视频feed流、视频投稿，个人信息、用户登录、用户注册      |
| 扩展功能一 | 视频点赞/取消点赞，点赞列表；用户评论/删除评论，视频评论列表 |
| 扩展功能二 |            用户关注/取关；用户关注列表、粉丝列表             |


## 目录结构

```c++
.
├── common
│   ├── AuthMiddleware.go
│   ├── cache.go
│   └── dbInit.go
├── config
│   └── config.go
├── config.yaml
├── controller
│   ├── commentController.go
│   ├── favortiteController.go
│   ├── feedController.go
│   ├── publishController.go
│   ├── relationController.go
│   └── userController.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── log
│   └── log.go
├── main.go
├── minioStore
│   └── minioClient.go
├── proto
│   ├── pkg
│   │   ├── comment.pb.go
│   │   ├── favorite.pb.go
│   │   ├── feed.pb.go
│   │   ├── login.pb.go
│   │   ├── publish.pb.go
│   │   ├── register.pb.go
│   │   ├── relation.pb.go
│   │   └── user.pb.go
│   └── proto
│       ├── comment.proto
│       ├── favorite.proto
│       ├── feed.proto
│       ├── login.proto
│       ├── publish.proto
│       ├── register.proto
│       ├── relation.proto
│       └── user.proto
├── README.md
├── redis.conf
├── repository
│   ├── commentModel.go
│   ├── favoriteModel.go
│   ├── relationModel.go
│   ├── userModel.go
│   └── videoModel.go
├── response
│   └── response.go
├── routes
│   ├── comment.go
│   ├── favorite.go
│   ├── publish.go
│   ├── relation.go
│   ├── routes.go
│   └── user.go
├── service
│   ├── commentService.go
│   ├── favoriteService.go
│   ├── feedService.go
│   ├── publishService.go
│   ├── relationService.go
│   └── userService.go
├── TikTokLite.sql
├── util
│   └── util.go
└── wait-for.sh
```

- `common`：中间件、数据库初始化
- `config`： 读取配置
- `controller`：视图层，处理前端消息
- `log`：zap日志组件进行封装
- `minioStore`：对象存储服务，生成视频对外访问连接
- `proto`：前端消息结构体，由`protobuf`文件自动生成
- `repository`：数据层，直接对数据库进行操作
- `response`：对返回消息进行封装
- `routes`：路由层
- `service`：逻辑层，执行业务操作，从数据层获取数据，封装后返回试图层
- `uitl`：工具函数
- `TikTokLite.sql`：数据库建表文件 
- `config.yaml`：配置文件
- `redis.conf`：redis配置文件
- `main.go`：服务入口

## 开发整体设计

### 整体架构图

![结构图](https://github.com/jhzol/test/blob/master/image/Tiktoklite.png?raw=true)






### 数据库设计

<img src="https://img-blog.csdnimg.cn/be4524a1a81e4a31a6699ea03c3466f2.png"   >

## 相关优化

- 使用redis作为缓存，通过减少对数据库的访问提升效率
- 使用redis锁、事务。
- SQL注入问题
- 使用jwt进行权限认证，由于未定义错误类型通知客户端权限过期重新登录，故暂未设置jwt过期时间
- 对数据表建立合理的外键来确保插入数据的准确性，查询时也可提升速度
- 对用户密码进行加密存储，保护用户账户安全
- 使用对象存储对视频及封面进行存储，生成对外访问url
- 使用docker整合所有相关依赖服务，便于用户快速部署服务，使用wait-for确保其他依赖服务启动后再启动后端服务

## 性能测试

通过命令 go tool pprof -http=:6060 "http://localhost:8080/debug/pprof/profile?seconds=120" 生成了两个版本的火焰图，左图为v1.0，右图为v1.2版本，通过对比两张详细火焰图，优化后的相同方法调用时间更短（添加了相应的中间件）

## 未来展望

- 分布式

  利用grpc作为分布式框架，etcd或zookeeper作为注册中心，将五个模块分别布置到不同的服务器上，通过RPC远程调用的方式，来调用相关的模块的方法，做到分布式处理与解耦

  ![image-20221027225809359](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20221027225809359.png)

## 线上地址

- http://112.74.109.70:8080/

## 贡献者

- 金浩哲
- 张建红
- 刘航
- 薛寅珊
- 吴志伟

*您也可以查阅仓库为该项目做出贡献的开发者*

## 版权说明

该项目签署了MIT 授权许可，详情请参阅 [LICENSE.txt](https://github.com/shaojintian/Best_README_template/blob/master/LICENSE.txt)

## 鸣谢

https://youthcamp.bytedance.com/
