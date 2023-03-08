# [**MicroserviceFreshMall**](https://github.com/chuyangc/MicroserviceFreshMall)

基于微服务架构的生鲜商城

### 目前架构图

[点击查看](https://www.processon.com/view/link/63e9c64cb9d870719cfc2d78)

### 环境准备（自行准备好Docker和Docker-Compose）

Go:1.15

```shell
# 准备安装目录
mkdir ~/go && cd ~/go
# 下载
wget https://dl.google.com/go/go1.15.10.linux-amd64.tar.gz

# 执行`tar`解压到`/usr/local`目录下（官方推荐）
tar -C /usr/local -zxvf  go1.15.10.linux-amd64.tar.gz

# 添加/usr/loacl/go/bin目录到PATH变量中。添加到/etc/profile或$HOME/.profile都可以
vi /etc/profile
# 在最后一行添加
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin

source /etc/profile

# 开启go module
# Global Proxy
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct

# 七牛云
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```



Python:3.8.6

```shell
# 提前安装好系统依赖包：
#CentOS: 
sudo yum install -y openssl-devel bzip2-devel expat-devel gdbm-devel readline-devel sqlite-devel gcc gcc-c++  openssl-devel libffi-devel python-devel mariadb-devel
#Ubuntu:
sudo apt-get install zlib1g-dev libbz2-dev libssl-dev libncurses5-dev default-libmysqlclient-dev libsqlite3-dev libreadline-dev tk-dev libgdbm-dev libdb-dev libpcap-dev xz-utils libexpat1-dev liblzma-dev libffi-dev libc6-dev
   
#1. 获取
wget https://www.python.org/ftp/python/3.8.6/Python-3.8.6.tgz
tar -xzvf Python-3.8.6.tgz -C  /tmp
cd  /tmp/Python-3.8.6/
#2. 把Python3.8安装到 /usr/local 目录
./configure --prefix=/usr/local
make
make altinstall
#3. 更改/usr/bin/python链接
ln -s /usr/local/bin/python3.8 /usr/bin/python3
ln -s /usr/local/bin/pip3.8 /usr/bin/pip3

# 通过豆瓣镜像下载包
pip3 install xxx -i http://pypi.douban.com/simple/
```



Nodejs:12.18.3

```shell
cd ~

wget https://nodejs.org/dist/v12.18.3/node-v12.18.3-linux-x64.tar.xz

tar -xvf node-v12.18.3-linux-x64.tar.xz

ln -s /root/node-v12.18.3-linux-x64/bin/node /usr/bin/node
ln -s /root/node-v12.18.3-linux-x64/bin/npm /usr/bin/npm

# 测试一下
node -v

# 安装npmmirror镜像
npm install -g cnpm --registry=https://registry.npmmirror.com
```

---

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

Jaeger

```shell
docker run \
      --rm \
      --name jaeger \
      -p6831:6831/udp \
      -p16686:16686 \
      jaegertracing/all-in-one:latest
```

