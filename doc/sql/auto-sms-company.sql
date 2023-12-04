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
  PRIMARY KEY (`company_id`) USING BTREE,
  UNIQUE INDEX `company_name_unique`(`company_name`) USING BTREE COMMENT 'name unique'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for as_department
-- ----------------------------
DROP TABLE IF EXISTS `as_department`;
CREATE TABLE `as_department`  (
  `department_id` int(11) NOT NULL AUTO_INCREMENT,
  `department_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `department_parent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `company_id` int(11) NOT NULL DEFAULT 0,
  `is_root` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'is root',
  `create_time` datetime(0) NULL DEFAULT NULL,
  `update_time` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`department_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
