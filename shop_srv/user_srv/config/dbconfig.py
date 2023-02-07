from playhouse.pool import PooledMySQLDatabase
from playhouse.shortcuts import ReconnectMixin


# 使用peewee的连接池， 使用ReconnectMixin来防止出现连接断开查询失败
class ReconnectMysqlDatabase(ReconnectMixin, PooledMySQLDatabase):
    pass


MYSQL_DB = "shop_user_srv"
MYSQL_HOST = "39.108.131.245"
MYSQL_PORT = 3306
MYSQL_USER = "root"
MYSQL_PASSWORD = "123456"

DB = ReconnectMysqlDatabase(database=MYSQL_DB, host=MYSQL_HOST, port=MYSQL_PORT, user=MYSQL_USER,
                            password=MYSQL_PASSWORD)
