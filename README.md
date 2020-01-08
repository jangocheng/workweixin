# workweixin
企业微信开发

查看所有的 `tag`, `git tag -l` 

`git checkout v0.1` 初始化企业应用开发，开启API接收和发送消息

`git checkout v0.2` 初始化企业应用开发，开启API管理通讯录，简单实现用户的加入和退出企业，更新用户信息。

`git checkout v0.3` 加入 TODO 功能， 在 企业应用 增加 查看，新建TODO 任务等。

`git checkout v0.4` 增加定时任务。

程序运行：

1.  `./init.sh` 建立docker网络
2. `cd cores/dbs && docker-compose up -d` 运行 `MySQL` 服务.
3. `./run.sh` 运行企业应用服务.