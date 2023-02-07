import argparse
import os.path
import socket
import sys
from concurrent import futures
import logging
import signal
from functools import partial

import consul
import grpc
import nacos
from loguru import logger

from user_srv.proto import user_pb2_grpc
from user_srv.handler.user import UserServicer
from common.grpc_health.v1 import health, health_pb2, health_pb2_grpc
from common.register import consul
from user_srv.config import consulconfig
from user_srv.config import srvconfig
from user_srv.config import nacosconfig

BASE_DIR = os.path.dirname(os.path.abspath(os.path.dirname(__file__)))
sys.path.insert(0, BASE_DIR)
REGISTER_MODE = "consul"  # consul || nacos


def on_exit(signum, frame, service_id):
    register = consul.ConsulRegister(consulconfig.CONSUL_HOST, consulconfig.CONSUL_PORT)
    logger.info(f"注销 {service_id} 服务")
    register.deregister(service_id=service_id)
    logger.info("注销成功")
    sys.exit(0)


def get_free_tcp_port():
    # 动态获取可用TCP的端口
    tcp = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    tcp.bind(("", 0))
    _, port = tcp.getsockname()
    tcp.close()
    return port


def serve():
    # 增加参数配置
    parser = argparse.ArgumentParser()
    parser.add_argument('--ip',
                        nargs="?",
                        type=str,
                        default="192.168.178.1",
                        help="binding ip"
                        )
    parser.add_argument('--port',
                        nargs="?",
                        type=int,
                        default=0,
                        help="listening port"
                        )
    args = parser.parse_args()

    if args.port == 0:
        port = get_free_tcp_port()
    else:
        port = args.port

    logger.add("logs/user_srv_{time}.log")

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    # 注册用户服务
    user_pb2_grpc.add_UserServicer_to_server(UserServicer(), server)

    # 注册健康检查
    health_pb2_grpc.add_HealthServicer_to_server(health.HealthServicer(), server)

    server.add_insecure_port(f'{args.ip}:{port}')

    import uuid
    service_id = str(uuid.uuid1())

    # 主进程退出信息监听
    """
        Windows Supported Signal
            SIGINT ctrl+C 终端
            SIGTERM kill发出的软件终止
    """
    signal.signal(signal.SIGINT, partial(on_exit, service_id=service_id))
    signal.signal(signal.SIGTERM, partial(on_exit, service_id=service_id))

    logger.info(f"启动服务: {args.ip}:{port}")
    server.start()

    logger.info(f"服务注册中心为:{REGISTER_MODE}")
    if REGISTER_MODE == "consul":
        # Consul
        register = consul.ConsulRegister(consulconfig.CONSUL_HOST, consulconfig.CONSUL_PORT)
        if not register.register(name=srvconfig.SERVICE_NAME,
                                 id=service_id,
                                 address=args.ip,
                                 port=port,
                                 tags=srvconfig.SERVICE_TAGS,
                                 check=None):
            logger.info(f"服务注册失败")
            sys.exit(0)
    elif REGISTER_MODE == "nacos":
        # Nacos
        client = nacos.NacosClient(server_addresses=nacosconfig.NACOS_SERVERADDR,
                                   namespace=nacosconfig.NACOS_NAMESPACEID,
                                   username=nacosconfig.NACOS_USERNAME,
                                   password=nacosconfig.NACOS_PASSWORD)
        # client.current_server(args.ip, args.port)
        succ = client.add_naming_instance("user-srv", args.ip, port)
        if not succ:
            logger.warning(f"服务注册失败")
            sys.exit(0)

    logger.success(f"服务注册成功-> user-srv")
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()
