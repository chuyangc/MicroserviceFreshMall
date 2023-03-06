import json
import redis
import nacos
from playhouse.pool import PooledMySQLDatabase
from playhouse.shortcuts import ReconnectMixin
from loguru import logger


# 使用peewee的连接池， 使用ReconnectMixin来防止出现连接断开查询失败
class ReconnectMysqlDatabase(ReconnectMixin, PooledMySQLDatabase):
    pass


NACOS = {
    "Host": "192.168.178.140",
    "Port": 8848,
    "NameSpace": "4e2dee4a-8132-45ab-8af1-4aa5f2d102dd",
    "User": "nacos",
    "Password": "nacos",
    "DataId": "inventory-srv.json",
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

# rocketmq的配置
# ROCKETMQ_HOST = data["rocketmq"]["host"]
# ROCKETMQ_PORT = data["rocketmq"]["port"]

# 服务相关的配置
SERVICE_NAME = data["name"]
SERVICE_TAGS = data["tags"]

REDIS_HOST = data["redis"]["host"]
REDIS_PORT = data["redis"]["port"]
REDIS_DB = data["redis"]["db"]

# 配置连接池
pool = redis.ConnectionPool(host=REDIS_HOST, port=REDIS_PORT, db=REDIS_DB)
REDIS_CLIENT = redis.StrictRedis(connection_pool=pool)

DB = ReconnectMysqlDatabase(data["mysql"]["db"], host=data["mysql"]["host"], port=data["mysql"]["port"],
                            user=data["mysql"]["user"], password=data["mysql"]["password"])

# TODO nacos的配置信息
NACOS_SERVERADDR = "192.168.178.138:8848"
NACOS_HOST = "192.168.178.138"
NACOS_PORT = 8848
NACOS_USERNAME = "nacos"
NACOS_PASSWORD = "nacos"

NACOS_NAMESPACE = "public"
NACOS_NAMESPACEID = "90bab2c3-dec3-4f07-87b8-09af866c0490"
NACOS_GROUP = "web"