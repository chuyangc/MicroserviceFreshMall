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

Kong+Konga+Postgres

```shell
docker run -d --name kong-database \
           --network=kong-net \
           -p 5432:5432 \
           -e "POSTGRES_USER=kong" \
           -e "POSTGRES_DB=kong" \
           -e "POSTGRES_PASSWORD=kong" \
           -e "POSTGRES_HOST_AUTH_METHOD=trust" \
           postgres:9.6

docker run --rm \
    --network=kong-net \
    -e "KONG_LOG_LEVEL=debug" \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    -e "KONG_PG_PASSWORD=kong" \
    -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
    kong:latest kong migrations bootstrap

docker run -d --name kong \
 -e "KONG_DATABASE=postgres" \
 -e "KONG_PG_HOST=192.168.178.140" \
 -e "KONG_DNS_RESOLVER=192.168.178.140:8600" \
 -e "KONG_PG_PASSWORD=kong" \
 -e "KONG_CASSANDRA_CONTACT_POINTS=kong-database" \
 -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
 -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
 -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
 -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
 -e "KONG_ADMIN_LISTEN=0.0.0.0:8001 reuseport backlog=16384, 0.0.0.0:8444 http2 ssl reuseport backlog=16384" \
 -p 8000:8000 \
 -p 8443:8443 \
 -p 8001:8001 \
 -p 8444:8444 \
 kong:latest

docker run -d -p 1337:1337 --name konga pantsel/konga

docker container update --restart=always kong konga kong-database
```

