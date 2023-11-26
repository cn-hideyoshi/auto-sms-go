/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.5
 Source Server Type    : MySQL
 Source Server Version : 80013
 Source Host           : 192.168.1.5:3306
 Source Schema         : auto-sms-company

 Target Server Type    : MySQL
 Target Server Version : 80013
 File Encoding         : 65001

 Date: 26/11/2023 11:16:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for as_company
-- ----------------------------
DROP TABLE IF EXISTS `as_company`;
CREATE TABLE `as_company`  (
  `company_id` int(11) NOT NULL AUTO_INCREMENT,
  `company_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `company_password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`company_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
