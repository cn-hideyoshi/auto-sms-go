SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for as_msg_group
-- ----------------------------
DROP TABLE IF EXISTS `as_msg_group`;
CREATE TABLE `as_msg_group`  (
  `group_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'msg group id',
  `group_name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'msg group name',
  `group_content` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'msg group content',
  `group_type` tinyint(3) NULL DEFAULT NULL COMMENT 'msg group type : 1. quick send 2. delay send 3. scheduled send',
  `template_id` int(11) NOT NULL COMMENT 'template_id',
  `create_time` datetime(0) NULL DEFAULT NULL COMMENT 'create time',
  `update_time` datetime(0) NULL DEFAULT NULL COMMENT 'update time',
  PRIMARY KEY (`group_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for as_msg_group_user
-- ----------------------------
DROP TABLE IF EXISTS `as_msg_group_user`;
CREATE TABLE `as_msg_group_user`  (
  `group_id` int(11) NOT NULL COMMENT 'msg group id',
  `user_id` int(11) NOT NULL COMMENT 'user id',
  `phone_id` int(11) NOT NULL COMMENT 'user phone id',
  `phone_no` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'phone no',
  PRIMARY KEY (`group_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
