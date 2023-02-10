import json
import grpc
import consul

from google.protobuf import empty_pb2

from inventory_srv.proto import inventory_pb2, inventory_pb2_grpc
from inventory_srv.config import settings


class InventroyTest:
    def __init__(self):
        # 连接grpc服务器
        c = consul.Consul(host="192.168.178.138", port=8500)
        # 获取所有的services
        services = c.agent.services()
        ip = ""
        port = ""
        for key, value in services.items():
            if value["Service"] == settings.SERVICE_NAME:
                ip = value["Address"]
                port = value["Port"]
                break
        if not ip:
            raise Exception()
        channel = grpc.insecure_channel(f"{ip}:{port}")
        self.inventory_stub = inventory_pb2_grpc.InventoryStub(channel)

    def set_inv(self):
        rsp = self.inventory_stub.SetInv(
            inventory_pb2.GoodsInvInfo(goodsId=10, num=110)
        )

    def get_inv(self):
        rsp = self.inventory_stub.InvDetail(
            inventory_pb2.GoodsInvInfo(goodsId=3)
        )
        print(rsp.num)

    def sell(self):
        goods_list = [(1, 10), (2, 20), (3, 30)]
        req = inventory_pb2.SellInfo()
        for goodsId, num in goods_list:
            req.goodsInfo.append(inventory_pb2.GoodsInvInfo(goodsId=goodsId, num=num))
        rsp = self.inventory_stub.Sell(req)

    def reback(self):
        goods_list = [(1, 6), (3, 3)]
        request = inventory_pb2.SellInfo()
        for goodsId, num in goods_list:
            request.goodsInfo.append(inventory_pb2.GoodsInvInfo(goodsId=goodsId, num=num))
        rsp = self.inventory_stub.Reback(request)


if __name__ == "__main__":
    inventory = InventroyTest()
    # inventory.set_inv()
    # inventory.get_inv()
    inventory.sell()
    # inventory.reback()
