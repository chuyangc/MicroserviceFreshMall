from datetime import datetime

from peewee import *

from inventory_srv.config.settings import DB


class BaseModel(Model):
    add_time = DateTimeField(default=datetime.now, verbose_name="添加时间")
    is_deleted = BooleanField(default=False, verbose_name="是否删除")
    update_time = DateTimeField(verbose_name="更新时间", default=datetime.now)

    def save(self, *args, **kwargs):
        # 判断这是一个新添加的数据还是更新的数据
        if self._pk is not None:
            # 这是一个新数据
            self.update_time = datetime.now()
        return super().save(*args, **kwargs)

    @classmethod
    def delete(cls, permanently=False):  # permanently表示是否永久删除
        if permanently:
            return super().delete()
        else:
            return super().update(is_deleted=True)

    def delete_instance(self, permanently=False, recursive=False, delete_nullable=False):
        if permanently:
            return self.delete(permanently).where(self._pk_expr()).execute()
        else:
            self.is_deleted = True
            self.save()

    @classmethod
    def select(cls, *fields):
        return super().select(*fields).where(cls.is_deleted == False)

    class Meta:
        database = DB


# TODO 如果有多个仓库的话要添加stock，此处做简单处理
# class Stock(BaseModel):
#     #仓库表
#     name = CharField(verbose_name="仓库名")
#     address = CharField(verbose_name="仓库地址")


class Inventory(BaseModel):
    # 商品的库存表
    # stock = PrimaryKeyField(Stock)
    goods = IntegerField(verbose_name="商品id", unique=True)
    stocks = IntegerField(verbose_name="库存数量", default=0)
    version = IntegerField(verbose_name="版本号", default=0)  # 分布式锁的乐观锁


class InventoryNew(BaseModel):
    # 商品的库存表
    # stock = PrimaryKeyField(Stock)
    goods = IntegerField(verbose_name="商品id", unique=True)
    stocks = IntegerField(verbose_name="库存数量", default=0)
    version = IntegerField(verbose_name="版本号", default=0)  # 分布式锁的乐观锁
    freeze = IntegerField(verbose_name="冻结数量", default=0)


# class Delivery(BaseModel):
#     goods = IntegerField(verbose_name="商品id", unique=True)
#     goods_name = CharField(verbose_name="商品名")
#     nums = IntegerField(verbose_name="数量", unique=True)
#     order_sn = CharField(verbose_name="订单号")
#     status = CharField(verbose_name="订单号", default="waitting")

# class InventoryHistory(BaseModel):
#     # 出库历史表
#     order_sn = CharField(verbose_name="订单编号", max_length=20, unique=True)
#     order_inv_detail = CharField(verbose_name="订单详情", max_length=200)
#     status = IntegerField(choices=((1, "已扣减"), (2, "已归还")), default=1, verbose_name="出库状态")


if __name__ == "__main__":
    # db.create_tables([Inventory])

    # for i in range(5):
    #     goods_inv = Inventory(goods=i, stocks=100)
    #     goods_inv.save()
    goods_info = ((1, 2), (2, 3), (3, 19))
    # #
    with DB.atomic() as txn:
        for goods_id, num in goods_info:
            # 查询库存
            goods_inv = Inventory.get(Inventory.goods == goods_id)
            if goods_inv.stocks < num:
                # 库存不足
                print(f"{goods_id}:库存不足")
                txn.rollback()  # 回滚
                break
            else:
                goods_inv.stocks -= num
                goods_inv.save()

    # for i in range(421, 841):
    #     try:
    #         inv = Inventory.get(Inventory.goods == i)
    #         inv.stocks = 100
    #         inv.save()
    #     except DoesNotExist as e:
    #         inv = Inventory(goods=i, stocks=100)
    #         inv.save(force_insert=True)
