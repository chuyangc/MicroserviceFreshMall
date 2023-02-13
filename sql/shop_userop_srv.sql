/*
 Navicat Premium Data Transfer

 Source Server         : 毕设mysql5.7
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 192.168.178.138:3306
 Source Schema         : shop_userop_srv

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 13/02/2023 18:02:09
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for address
-- ----------------------------
DROP TABLE IF EXISTS `address`;
CREATE TABLE `address`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(0) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  `update_time` datetime(0) NOT NULL,
  `user` int(11) NOT NULL,
  `province` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `city` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `district` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `signer_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `signer_mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of address
-- ----------------------------
INSERT INTO `address` VALUES (1, '2023-02-13 16:00:01', 0, '2023-02-13 16:00:01', 1, '辽宁省', '焦作市', '华北', 'Maria Williams', 'Ronald Wilson', '18686868686');

-- ----------------------------
-- Table structure for leavingmessages
-- ----------------------------
DROP TABLE IF EXISTS `leavingmessages`;
CREATE TABLE `leavingmessages`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(0) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  `update_time` datetime(0) NOT NULL,
  `user` int(11) NOT NULL,
  `message_type` int(11) NOT NULL,
  `subject` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `message` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `file` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of leavingmessages
-- ----------------------------
INSERT INTO `leavingmessages` VALUES (1, '2023-02-13 16:02:29', 0, '2023-02-13 16:02:29', 1, 1, 'sit', 'veniam consequat nisi dolor laborum', 'mailto://lcrhbfxjmi.gb/lrtdwdx');

-- ----------------------------
-- Table structure for userfav
-- ----------------------------
DROP TABLE IF EXISTS `userfav`;
CREATE TABLE `userfav`  (
  `add_time` datetime(0) NOT NULL,
  `is_deleted` tinyint(1) NOT NULL,
  `update_time` datetime(0) NOT NULL,
  `user` int(11) NOT NULL,
  `goods` int(11) NOT NULL,
  PRIMARY KEY (`user`, `goods`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of userfav
-- ----------------------------
INSERT INTO `userfav` VALUES ('2023-02-13 15:59:13', 0, '2023-02-13 15:59:13', 1, 422);

SET FOREIGN_KEY_CHECKS = 1;
