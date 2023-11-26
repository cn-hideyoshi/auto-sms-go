/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.5
 Source Server Type    : MySQL
 Source Server Version : 80013
 Source Host           : 192.168.1.5:3306
 Source Schema         : auto-sms-user

 Target Server Type    : MySQL
 Target Server Version : 80013
 File Encoding         : 65001

 Date: 26/11/2023 11:16:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for as_user
-- ----------------------------
DROP TABLE IF EXISTS `as_user`;
CREATE TABLE `as_user`  (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'user_id',
  `user_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'user_name',
  `user_password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'user_password',
  `company_id` int(11) NOT NULL DEFAULT 0 COMMENT 'company id',
  `create_time` timestamp(0) NULL DEFAULT NULL COMMENT 'create_time',
  `update_time` timestamp(0) NULL DEFAULT NULL COMMENT 'update_time',
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
