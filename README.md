# workweixin

企业微信开发

查看所有的 `tag`, `git tag -l` 

`git checkout v0.1` 初始化企业应用开发，开启API接收和发送消息

`git checkout v0.2` 初始化企业应用开发，开启API管理通讯录，简单实现用户的加入和退出企业，更新用户信息。

`git checkout v0.3` 加入 TODO 功能， 在 企业应用 增加 查看，新建TODO 任务等。

`git checkout v0.4` 增加定时任务。

`git checkout v0.5` 启用 GRPC

`git checkout v0.6` 微服务拆分已经项目结构重构

`git checkout v0.7` 加入Redis缓存

`git checkout v0.8` 加入链路跟踪

程序运行：

1. `cd databases/  && ./init.sh && docker-compose up -d` 建立docker网络，启动 MySQL 和 Redis 服务

2. `./jaeger/.init.sh` 起 tracing 服务，通过 16686 端口进行查看信息。（用的是测试docker）

3. `./run.sh` 构建基础镜像

4. `make todosrv action=up ` 启动 `todosrv` ，需要第一个启动

5. `make appsrv action=up` 启动 `appsrv`, 开启 微信企业应用服务

6. `make contactsrv action=up` 开启 `contactsrv` 通信录管理服务
