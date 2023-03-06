# FinalDesign

毕业设计——基于微服务架构的生鲜商城

### 目前架构图

[点击查看](https://www.processon.com/view/link/63e9c64cb9d870719cfc2d78)

### 环境准备（自行准备好Docker和Docker-Compose）

Go:1.15

Python:3.8.6

Nodejs:12.18.3

```shell
# 安装npmmirror镜像
npm install -g cnpm --registry=https://registry.npmmirror.com
```



[Yapi](http://yapi.smart-xwork.cn/)

MySQL

```shell
docker pull mysql:5.7
docker run -p 3306:3306 --name mymysql -v $PWD/conf:/etc/mysql/conf.d -v $PWD/logs:/logs -v $PWD/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7

docker exec -it mymysql /bin/bash
mysql -uroot -p123456
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY 'root' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO 'root'@'127.0.0.1' IDENTIFIED BY 'root' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost' IDENTIFIED BY 'root' WITH GRANT OPTION;
FLUSH PRIVILEGES;
exit;

docker container update --restart=always mymysql
```

Redis

```shell
docker run -p 6379:6379 -d redis:latest redis-server

docker container update --restart=always redis-server
```

Consul

```shell
docker run -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600/udp consul consul agent  -dev -client=0.0.0.0

docker container update --restart=always consul
```

Nacos

```shell
docker run --name nacos-standalone -e MODE=standalone -e JVM_XMS=256m -e JVM_XMX=256m -e JVM_XMN=128m -p 8848:8848 -d nacos/nacos-server:latest

docker container update --restart=always nacos
# JVM_XMS JVM_XMX JVM_XMN
# 上面这三个参数后面的值可以根据自己机器的配置
```

[RocketMQ](https://www.yuque.com/attachments/yuque/0/2021/zip/22466542/1638598031338-1e6dae98-328c-4efa-8484-c76c7c7f9c64.zip?from=https%3A%2F%2Fwww.yuque.com%2Fjintianjiandaoyibaikuaiqian%2Fgfmuhg%2Fdm56fe)

