import nacos


# Both HTTP/HTTPS protocols are supported, if not set protocol prefix default is HTTP, and HTTPS with no ssl check(verify=False)
# "192.168.3.4:8848" or "https://192.168.3.4:443" or "http://192.168.3.4:8848,192.168.3.5:8848" or "https://192.168.3.4:443,https://192.168.3.5:443"
# 监听配置集变动函数
def listen_cfg(args):
    print("配置集发生变化")
    print(args)


SERVER_ADDRESSES = "127.0.0.1:8848"  # nacos的地址和端口
NAMESPACE = "public"  # 需要拉取配置集的名称空间

# 无密码拉取模式
# client = nacos.NacosClient(SERVER_ADDRESSES, namespace=NAMESPACE)
# 有密码拉取模式
client = nacos.NacosClient(SERVER_ADDRESSES, namespace=NAMESPACE, username="nacos", password="nacos")

# get config
# data_id = "xxx"  # 配置集的DataID
# group = "xxx"  # 配置集的组
# print(client.get_config(data_id, group))

# 设置变动监听
# if __name__ == "__main__":
#     # 注意这个小坑，下面这个代码不能在顶行写，在windows下会报错。
#     # client.add_config_watcher(data_id, group, listen_cfg)
#     print(client.add_naming_instance("user-web", "127.0.0.1", 50051))
#     client.remove_naming_instance("user-web","127.0.0.1",50051)
