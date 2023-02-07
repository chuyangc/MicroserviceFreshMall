/*
 Navicat Premium Data Transfer

 Source Server         : alicloud
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           : 39.108.131.245:3306
 Source Schema         : shop_user_srv

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 07/02/2023 12:52:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mobile` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `nick_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `head_url` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `birthday` date DEFAULT NULL,
  `address` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `desc` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `gender` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
  `role` int(11) NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_mobile`(`mobile`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '18888888880', '$pbkdf2-sha256$29000$jtEaY8yZUyplTOl9z3mvlQ$8gf2OZgvf7Qn9WvNoYcZ5cHFUG4Mgv97XvQOvkijoEc', 'chuyangc0', NULL, '2023-02-05', NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (2, '18888888881', '$pbkdf2-sha256$29000$1xpjjNFa6x2DEEIIAcAYQw$IV7M7xbjkTXfVsJ2by8zj2TXZ3ggdTQXhmzo.QGrffE', 'chuyangc1', NULL, '2023-02-05', NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (3, '18888888882', '$pbkdf2-sha256$29000$U6q1ljIG4LzXmnNOiXGOMQ$nS9Ywoimc3.0H4YpM8YE6FP/tV4NNVPhKFA4lSiHPO8', 'chuyangc2', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (4, '18888888883', '$pbkdf2-sha256$29000$OQcgxBgDIOScs9Z6z9nbmw$yydtBx.PpYMIAi0e5IgE2UXgc09E6jHFu3n0peWPO0k', 'chuyangc3', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (5, '18888888884', '$pbkdf2-sha256$29000$MWYshVAKAeA8hzBmDOH8vw$hmoMzaP5939PaNl4uYApltE1ml6ewwRUaIPtpBcebI0', 'chuyangc4', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (6, '18888888885', '$pbkdf2-sha256$29000$ydmbs3YuhTCGcM45J2RsTQ$GNgogRFhFdJzAdtnaPrEumHJwegC66KjpI17a27nRe4', 'chuyangc5', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (7, '18888888886', '$pbkdf2-sha256$29000$7v3fG8NYC2GsNaY0BkBobQ$z34cyM/uNmhCXVv9KUgyWYc0RJ1WNraeg.4TXuw3xU4', 'chuyangc6', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (8, '18888888887', '$pbkdf2-sha256$29000$f..d09obwxgDAOD8XwtBSA$oKC62NdzeqCzqFTLvv8AMzEmstayfi4F0mWjXpZChQA', 'chuyangc7', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (9, '18888888888', '$pbkdf2-sha256$29000$UEpJ6T1HCAHA.P.f8957Lw$D9QL3n/HHzEi6XnOhgACSUsN1RYxLqIfEpB2IpZTxMk', 'chuyangc8', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (10, '18888888889', '$pbkdf2-sha256$29000$R.hdi1GK0bpXitF6T8m5Fw$bkupkX3aVmJaaLdBX9iFYiSsk/VBhCI8WdrNbsnVXKI', 'chuyangc9', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (11, '18888888810', '$pbkdf2-sha256$29000$n3POufe.V2oNYWwNwTinVA$oHYi8MTu6CNiuasUq6824zo8FSXWyK0iK0vabMLAItk', '18888888810', NULL, NULL, NULL, NULL, NULL, 1);
INSERT INTO `user` VALUES (12, '18475812345', '$pbkdf2-sha256$29000$zdl77x2DcK6V0jqHsHYOoQ$/Q.L0Bit9UbkhG0Mnpcsai6myR7P.uth0Ko/gxBZJ24', '18475812345', NULL, NULL, NULL, NULL, NULL, 1);

SET FOREIGN_KEY_CHECKS = 1;
