import json

from playhouse.pool import PooledMySQLDatabase
from playhouse.shortcuts import ReconnectMixin

import nacos
from loguru import logger


# 使用peewee的连接池， 使用ReconnectMixin来防止出现连接断开查询失败
class ReconnectMysqlDatabase(ReconnectMixin, PooledMySQLDatabase):
    pass


NACOS = {
    "Host": "192.168.178.138",
    "Port": 8848,
    "NameSpace": "f7f4d603-9f3d-4146-bb9c-83eac2f07233",
    "User": "nacos",
    "Password": "nacos",
    "DataId": "user-srv.json",
    "Group": "dev"
}

client = nacos.NacosClient(f'{NACOS["Host"]}:{NACOS["Port"]}', namespace=NACOS["NameSpace"],
                           username=NACOS["User"],
                           password=NACOS["Password"])

# 获取配置
data = client.get_config(NACOS["DataId"], NACOS["Group"])
data = json.loads(data)
logger.success(f"成功从Nacos配置中心加载配置 -> {data}")


def check_update_cfg(args):
    logger.warning(f"配置产生变化 -> {args}")


# consul的配置
CONSUL_HOST = data["consul"]["host"]
CONSUL_PORT = data["consul"]["port"]

# 服务相关的配置
SERVICE_NAME = data["name"]
SERVICE_TAGS = data["tags"]

# TODO nacos的配置信息
NACOS_SERVERADDR = "127.0.0.1:8848"
NACOS_HOST = "127.0.0.1"
NACOS_PORT = 8848
NACOS_USERNAME = "nacos"
NACOS_PASSWORD = "nacos"

NACOS_NAMESPACE = "public"
NACOS_NAMESPACEID = "838b806f-20e6-4c3b-abf2-0e3b3bc73dcc"
NACOS_GROUP = "web"

DB = ReconnectMysqlDatabase(data["mysql"]["db"], host=data["mysql"]["host"], port=data["mysql"]["port"],
                            user=data["mysql"]["user"], password=data["mysql"]["password"])
