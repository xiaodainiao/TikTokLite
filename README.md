<div align=center><img src="https://img-blog.csdnimg.cn/f56709e579474f4ba755541c6a29e920.jpeg"  width="60%" height="30%"  ></div>


<p align = 'center'>
<a href="https://github.com/xiaodainiao"><img src="https://img.shields.io/badge/GitHub-xiaodainiao-blueviolet?logo=github">
</a>
 <a href="https://blog.csdn.net/weixin_45043334?spm=1010.2135.3001.5343"><img src="https://img.shields.io/badge/博客-小呆鸟_coding-orange?logo=blogger">
</a>
<a href="http://www.xiaodainiao.xyz/"><img src="https://img.shields.io/badge/%E4%B8%AA%E4%BA%BA%E7%BD%91%E7%AB%99-www.xiaodainiao.xyz-blue?logo=googlehome">
  </a>
<a href=""><img src="https://img.shields.io/badge/%E5%85%AC%E4%BC%97%E5%8F%B7-%E5%B0%8F%E5%91%86%E9%B8%9F%E5%93%87-critical?logo=wechat">
</a>
<a href="https://www.zhihu.com/people/axiao-hong-jia-meng-200"><img src="https://img.shields.io/badge/%E7%9F%A5%E4%B9%8E-%E5%B0%8F%E5%91%86%E9%B8%9F-ff69b4?logo=zhihu">
</a>  
</p>

&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**You just work hard, the rest is up to time**

**个人简介**
- 学生一枚，热爱编程，从0学c++，自己深有体会，目前学过c++,go,python,数据库,分布式等。个人热爱开源，希望自己学的知识可以和大家共享，并且一起不断进步。
- 后续打算学分布式框架，以及相关的技术。一些文章也会同步到博客，如果有需要的小伙伴，可以去博客自取。




# TikTokLite

<div align=center><img src="https://img-blog.csdnimg.cn/cc03d16ceea3494e8b38ce0f4a5eb0f6.png" alt="image-20221027202423203" width="65%" height="65%"  ></div>



&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;**极简抖音**


&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;[探索本项目相关文档](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145)

## 目录



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

<img src="https://img-blog.csdnimg.cn/cc6070c6e54a40dc95ea9b34cd855aa8.png"  width="65%" height="65%"  >



### 数据库设计

<img src="https://img-blog.csdnimg.cn/be4524a1a81e4a31a6699ea03c3466f2.png" width="65%" height="65%"  >

## 优化

### 1. 安全

1. 引入JWT，进行`全局Token管理`，高效管理用户Token，并且设置过期时间。
2. Redis引入`redsync锁`机制，防止俩个线程同时修改用户信息(例如关注)
3. Redis引入`事务`机制，防止多表操作时，只修改一张表。最终导致失败。
4. 使用参数占位符来构造SQL语句，不使用字符串拼接，`避免SQL注入`
5. 用户密码进行`MD5加密`处理，返回用户基本信息时进行`脱敏`。
6. 实现`鉴权中间件`，将鉴权和实际业务分离，对不同的接口设置不同的访问权限
7. 使用`docker`整合所有相关依赖服务，便于用户快速部署服务，使用wait-for确保其他依赖服务启动后再启动后端服务

### 2. 性能
1. 根据实际业务,Querry语句的需求，合理`设置相关索引`，保证索引高命中
2. 引入`Redis`作为中间件，用来实现对象缓存，提升响应速度，减少IO操作，减少服务器压力
3. 通过`Minio`自己搭建对象存储，来存储上传视频，并且将上传的视频生成URL，并将URL放在数据库中，避免存储冗余。
4. 通过`pprof`进行性能测试，引入缓存与无缓存之间的性能

### 3. 项目维护

1. 项目`Git协同`,严格遵循成员PR->Review->Merge三步走流程，避免错误代码扩散到其他成员库
2. `多次迭代目录结构`。目录结构清晰，配置单元、日志单元、各模块单元条理分明

3. 文档管理。修改代码前后，要随时记录文档，跟进开发流程，字最后测试出现问题时，可以找到具体负责人快速解决


## 性能测试

通过命令 go tool pprof -http=:6060 "http://localhost:8080/debug/pprof/profile?seconds=120" 生成了两个版本的火焰图，左图为v1.0，右图为v1.2版本，通过对比两张详细火焰图，优化后的相同方法调用时间更短（添加了相应的中间件）

<img src="https://img-blog.csdnimg.cn/51fba6bd3cac43a4b4591a093f32a73f.png" width="65%" height="65%"  >
<img src="https://img-blog.csdnimg.cn/5c0a9e2f47544d7b9af7bba073de60d1.png" width="65%" height="65%"  >

## 未来展望

- 分布式

  利用grpc作为分布式框架，etcd或zookeeper作为注册中心，将五个模块分别布置到不同的服务器上，通过RPC远程调用的方式，来调用相关的模块的方法，做到分布式处理与解耦

<img src="https://img-blog.csdnimg.cn/f8e7445378f04f8ba77772a774c2afc0.png"  width="65%" height="65%" >

## 线上地址

- http://112.74.109.70:8080/



## 版本控制
- 该项目使用Git进行版本管理。您可以在repository参看当前可用版本。

## 贡献者

- 金浩哲
- 张建红
- 刘航
- 薛寅珊
- 吴志伟

*您也可以查阅仓库为该项目做出贡献的开发者*


## 鸣谢

[字节跳动青训营](https://youthcamp.bytedance.com/)
