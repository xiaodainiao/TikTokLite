# TikTokLite
<img src="C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20221027202423203.png" alt="image-20221027202423203" style="zoom: 50%;" />

​                                                           **极简版抖音**                                           

​                                                       [探索本项目相关文档](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145) 

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

###  目录结构

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

### 整体架构

![结构图](https://github.com/jhzol/test/blob/master/image/Tiktoklite.png?raw=true)

### 实现功能

|    功能    |                             说明                             |
| :--------: | :----------------------------------------------------------: |
|  基础功能  |      视频feed流、视频投稿，个人信息、用户登录、用户注册      |
| 扩展功能一 | 视频点赞/取消点赞，点赞列表；用户评论/删除评论，视频评论列表 |
| 扩展功能二 |            用户关注/取关；用户关注列表、粉丝列表             |

### 相关优化

- 使用redis作为缓存，通过减少对数据库的访问提升效率
- 使用redis锁、事务。
- SQL注入问题
- 使用jwt进行权限认证，由于未定义错误类型通知客户端权限过期重新登录，故暂未设置jwt过期时间
- 对数据表建立合理的外键来确保插入数据的准确性，查询时也可提升速度
- 对用户密码进行加密存储，保护用户账户安全
- 使用对象存储对视频及封面进行存储，生成对外访问url
- 使用docker整合所有相关依赖服务，便于用户快速部署服务，使用wait-for确保其他依赖服务启动后再启动后端服务

### 小组分工

- 金浩哲：

  整体框架、基础功能、扩展功能二、用户信息缓存、jwt鉴权鉴权、docker部署

- 张建红

  用户评论列表、评论信息缓存、密码加密、jwt鉴权、扩展功能二、点赞缓存

- 刘航、薛寅珊

  视频点赞、点赞列表

- 吴志伟

  用户评论

### 线上地址

http://112.74.109.70:8080/

