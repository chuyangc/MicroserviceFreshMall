import json

import grpc
from loguru import logger
from peewee import DoesNotExist
from google.protobuf import empty_pb2

from inventory_srv.config import settings
from inventory_srv.model.models import Inventory
from inventory_srv.proto import inventory_pb2, inventory_pb2_grpc


class InventoryServicer(inventory_pb2_grpc.InventoryServicer):
    @logger.catch
    def SetInv(self, request: inventory_pb2.GoodsInvInfo, context):
        # 设置库存，后续如果需要修改库存，也可以使用该接口
        # 判断是插入数据还是更新数据
        force_insert = False
        invs = Inventory.select().where(Inventory.goods == request.goodsId)
        if not invs:
            inv = Inventory()
            inv.goods = request.goodsId
            force_insert = True
        else:
            inv = invs[0]
        inv.goods = request.goodsId
        inv.stocks = request.num
        inv.save(force_insert=force_insert)

        return empty_pb2.Empty()

    @logger.catch
    def InvDetail(self, request: inventory_pb2.GoodsInvInfo, context):
        # 获取某个商品的库存详情
        try:
            inv = Inventory.get(Inventory.goods == request.goodsId)
            return inventory_pb2.GoodsInvInfo(goodsId=inv.goods, num=inv.stocks)
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("没有库存记录")
            return inventory_pb2.GoodsInvInfo()

    @logger.catch
    def Sell(self, request: inventory_pb2.GoodsInvInfo, context):
        # TODO 未解决超卖问题
        # 扣减库存
        with settings.DB.atomic() as txn:
            for item in request.goodsInfo:
                # 查询库存
                try:
                    goods_inv = Inventory.get(Inventory.goods == item.goodsId)
                except DoesNotExist as e:
                    txn.rollback()
                    context.set_code(grpc.StatusCode.NOT_FOUND)
                    return empty_pb2.Empty()
                if goods_inv.stocks < item.num:
                    # 库存不足
                    context.set_code(grpc.StatusCode.RESOURCE_EXHAUSTED)
                    context.set_details("库存不足")
                    # 事务回滚
                    txn.rollback()
                    return empty_pb2.Empty()
                else:
                    # TODO 可能会引起数据不一致，需要加入分布式锁
                    # 库存充足
                    goods_inv.stocks -= item.num
                    goods_inv.save()
            return empty_pb2.Empty()

    @logger.catch
    def Reback(self, request: inventory_pb2.GoodsInvInfo, context):
        # 归还库存，此处做订单超时自动归还
        with settings.DB.atomic() as txn:
            for item in request.goodsInfo:
                # 查询库存
                try:
                    goods_inv = Inventory.get(Inventory.goods == item.goodsId)
                except DoesNotExist as e:
                    txn.rollback()
                    context.set_code(grpc.StatusCode.NOT_FOUND)
                    return empty_pb2.Empty()
                # TODO 可能会引起数据不一致，需要加入分布式锁
                # 库存充足
                goods_inv.stocks += item.num
                goods_inv.save()
            return empty_pb2.Empty()
