/*
 Navicat Premium Data Transfer

 Source Server         : 毕设mysql5.7
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 192.168.178.138:3306
 Source Schema         : shop_order_srv

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 12/02/2023 00:22:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ordergoods
-- ----------------------------
DROP TABLE IF EXISTS `ordergoods`;
CREATE TABLE `ordergoods`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(0) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  `update_time` datetime(0) NOT NULL,
  `order` int(11) NOT NULL,
  `goods` int(11) NOT NULL,
  `goods_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `goods_image` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `goods_price` decimal(10, 5) NOT NULL,
  `nums` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ordergoods
-- ----------------------------
INSERT INTO `ordergoods` VALUES (1, '2023-02-11 17:21:33', 0, '2023-02-11 17:21:33', 1, 421, '烟台红富士苹果12个 净重2.6kg以上', 'https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/df392d01993cdab9de740fe17798bda1', 44.90000, 5);
INSERT INTO `ordergoods` VALUES (2, '2023-02-11 17:21:33', 0, '2023-02-11 17:21:33', 1, 422, '西州蜜瓜25号哈密瓜 2粒装 单果1.2', 'https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/c3dee23a62efe14bbd4fc2c70046dc73', 36.90000, 3);
INSERT INTO `ordergoods` VALUES (3, '2023-02-11 17:23:05', 0, '2023-02-11 17:23:05', 2, 421, '烟台红富士苹果12个 净重2.6kg以上', 'https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/df392d01993cdab9de740fe17798bda1', 44.90000, 5);
INSERT INTO `ordergoods` VALUES (4, '2023-02-11 17:23:05', 0, '2023-02-11 17:23:05', 2, 422, '西州蜜瓜25号哈密瓜 2粒装 单果1.2', 'https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/c3dee23a62efe14bbd4fc2c70046dc73', 36.90000, 3);
INSERT INTO `ordergoods` VALUES (5, '2023-02-11 17:30:00', 0, '2023-02-11 17:30:00', 3, 421, '烟台红富士苹果12个 净重2.6kg以上', 'https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/df392d01993cdab9de740fe17798bda1', 44.90000, 5);
INSERT INTO `ordergoods` VALUES (6, '2023-02-11 17:30:00', 0, '2023-02-11 17:30:00', 3, 422, '西州蜜瓜25号哈密瓜 2粒装 单果1.2', 'https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/c3dee23a62efe14bbd4fc2c70046dc73', 36.90000, 3);
INSERT INTO `ordergoods` VALUES (7, '2023-02-11 23:26:11', 0, '2023-02-11 23:26:11', 4, 421, '烟台红富士苹果12个 净重2.6kg以上', 'https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/df392d01993cdab9de740fe17798bda1', 44.90000, 5);
INSERT INTO `ordergoods` VALUES (8, '2023-02-11 23:26:11', 0, '2023-02-11 23:26:11', 4, 423, '越南进口红心火龙果 4个装 红肉中果 单', 'https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/b39672c6abebe124b982250642cb9a0f', 27.90000, 5);

-- ----------------------------
-- Table structure for orderinfo
-- ----------------------------
DROP TABLE IF EXISTS `orderinfo`;
CREATE TABLE `orderinfo`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(0) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  `update_time` datetime(0) NOT NULL,
  `user` int(11) NOT NULL,
  `order_sn` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `pay_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `trade_no` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `order_amount` float NOT NULL,
  `pay_time` datetime(0) DEFAULT NULL,
  `address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `signer_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `singer_mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `post` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `orderinfo_order_sn`(`order_sn`) USING BTREE,
  UNIQUE INDEX `orderinfo_trade_no`(`trade_no`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of orderinfo
-- ----------------------------
INSERT INTO `orderinfo` VALUES (1, '2023-02-11 17:21:33', 0, '2023-02-11 17:21:33', 1, '20230211172133191', 'alipay', 'paying', NULL, 335.2, NULL, '清远市', 'chuyangc', '18488888888', '请尽快发货');
INSERT INTO `orderinfo` VALUES (2, '2023-02-11 17:23:05', 0, '2023-02-11 17:23:05', 1, '20230211172305121', 'alipay', 'paying', NULL, 335.2, NULL, '清远市', 'chuyangc', '18488888888', '请尽快发货');
INSERT INTO `orderinfo` VALUES (3, '2023-02-11 17:30:00', 0, '2023-02-11 17:30:00', 1, '20230211173000170', 'alipay', 'paying', NULL, 335.2, NULL, '清远市', 'chuyangc', '18488888888', '请尽快发货');
INSERT INTO `orderinfo` VALUES (4, '2023-02-11 23:26:11', 0, '2023-02-11 23:26:11', 1, '20230211232611131', 'alipay', 'paying', NULL, 364, NULL, '柯城区', 'Susan Smith', '18788989898', '请尽快发货');

-- ----------------------------
-- Table structure for shoppingcart
-- ----------------------------
DROP TABLE IF EXISTS `shoppingcart`;
CREATE TABLE `shoppingcart`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(0) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  `update_time` datetime(0) NOT NULL,
  `user` int(11) NOT NULL,
  `goods` int(11) NOT NULL,
  `nums` int(11) NOT NULL,
  `checked` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of shoppingcart
-- ----------------------------
INSERT INTO `shoppingcart` VALUES (1, '2023-02-11 17:14:33', 1, '2023-02-11 17:15:16', 1, 421, 5, 1);
INSERT INTO `shoppingcart` VALUES (2, '2023-02-11 17:17:20', 0, '2023-02-11 22:45:46', 1, 422, 11, 0);
INSERT INTO `shoppingcart` VALUES (3, '2023-02-11 22:18:45', 1, '2023-02-11 22:18:45', 1, 423, 5, 0);

SET FOREIGN_KEY_CHECKS = 1;
