SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for as_user
-- ----------------------------
DROP TABLE IF EXISTS `as_user`;
CREATE TABLE `as_user`  (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'user id',
  `user_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'username',
  `user_password` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'user password',
  `department_id` int(11) NOT NULL DEFAULT 0 COMMENT 'department id',
  `company_id` int(11) NOT NULL DEFAULT 0 COMMENT 'company id',
  `create_time` timestamp(0) NULL DEFAULT NULL COMMENT 'create time',
  `update_time` timestamp(0) NULL DEFAULT NULL COMMENT 'update time',
  PRIMARY KEY (`user_id`) USING BTREE,
  UNIQUE INDEX `unique_username`(`user_name`) USING BTREE COMMENT 'unique_username'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for as_user_phone
-- ----------------------------
DROP TABLE IF EXISTS `as_user_phone`;
CREATE TABLE `as_user_phone`  (
  `phone_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `phone_no` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`phone_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
