from peewee import *
from passlib.hash import pbkdf2_sha256

from user_srv.config import dbconfig


class BaseModel(Model):
    class Meta:
        database = dbconfig.DB


class User(BaseModel):
    GENDER_CHOICES = (
        ("female", "女"),
        ("male", "男")
    )

    ROLE_CHOICES = (
        (1, "普通用户"),
        (2, "管理员")
    )
    # 用户模型
    mobile = CharField(max_length=11, index=True, unique=True, verbose_name="手机号码")
    password = CharField(max_length=100, verbose_name="密码")
    nick_name = CharField(max_length=20, null=True, verbose_name="昵称")
    head_url = CharField(max_length=200, null=True, verbose_name="头像")
    birthday = DateField(null=True, verbose_name="生日")
    address = CharField(max_length=200, null=True, verbose_name="地址")
    desc = TextField(null=True, verbose_name="个人简介")
    gender = CharField(max_length=6, choices=GENDER_CHOICES, null=True, verbose_name="性别")
    role = IntegerField(default=1, choices=ROLE_CHOICES, verbose_name="用户角色")


if __name__ == '__main__':
    # dbconfig.DB.create_tables([User])
    # Use Message Digest Algorithm5
    # for i in range(10):
    #     user = User()
    #     user.nick_name = f"chuyangc{i}"
    #     user.mobile = f"1888888888{i}"
    #     user.password = pbkdf2_sha256.hash("admin123")
    #     user.save()

    for user in User.select():
        print(pbkdf2_sha256.verify("admin123", user.password))
