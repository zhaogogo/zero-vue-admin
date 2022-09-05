/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : 127.0.0.1:3306
 Source Schema         : usercenter

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 04/09/2022 17:49:51
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `nick_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `mobile` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态  0：禁用   1：正常',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新人',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '删除人',
  `delete_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`),
  KEY `idx_delete_at` (`delete_at`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户管理';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` VALUES (1, 'admin', '超管', '', '$2a$05$pnqfACHHMlD3mFGabfUaxeF3k0IsgLlRs7rH1AQqj/ZilcrSfR1qO', 'admin@qq.com', '13612345678', 1, 'admin', '2017-08-14 11:11:11', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (22, 'liubei', '刘备', '', '123456', 'test@qq.com', '13889700023', 1, 'admin', '2018-09-23 19:43:00', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (23, 'zhaoyun', '赵云', '', '123456', 'test@qq.com', '13889700023', 1, 'admin', '2018-09-23 19:43:44', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (24, 'zhugeliang', '诸葛亮', '', '123456', 'test@qq.com', '13889700023', 7, 'admin', '2018-09-23 19:44:23', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (25, 'caocao', '曹操', '', '123456', 'test@qq.com', '13889700023', 1, 'admin', '2018-09-23 19:45:32', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (26, 'dianwei', '典韦', '', '123456', 'test@qq.com', '13889700023', 1, 'admin', '2018-09-23 19:45:48', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (27, 'xiahoudun', '夏侯惇', '', '123456', 'test@qq.com', '13889700023', 1, 'admin', '2018-09-23 19:46:09', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (28, 'xunyu', '荀彧', '', '123456', 'test@qq.com', '13889700023', 1, 'admin', '2018-09-23 19:46:38', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (29, 'sunquan', '孙权', '', '123456', 'test@qq.com', '13889700023', 1, 'admin', '2018-09-23 19:46:54', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (30, 'zhouyu', '周瑜', '', '123456', 'test@qq.com', '13889700023', 1, 'admin', '2018-09-23 19:47:28', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (31, 'luxun', '陆逊', '', '123456', 'test@qq.com', '13889700023', 1, 'admin', '2018-09-23 19:47:44', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (32, 'huanggai', '黄盖', '', '', 'test@qq.com', '13889700023', 1, '', '2018-09-23 19:48:38', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (33, '1', '1', '', '123456', '1', '1', 1, 'admin', '2021-04-26 17:57:50', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (35, '12', '1', '', '123456', '1', '1', 1, 'admin', '2021-04-26 18:01:53', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (36, '12313', '12', '', '123456', '1', '1', 1, 'admin', '2021-04-26 18:03:07', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (37, '324', '1', '', '123456', '1', '1', 1, 'admin', '2021-04-26 18:07:31', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (38, 'aa', 'aa', '', '123456', 'a', 'a', 1, 'admin', '2021-04-27 11:24:14', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (39, '133', '133', '', '', '1121', '1', 1, 'admin', '2021-04-27 12:30:15', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (40, 'liu', 'liu', '', '123456', '1002219331@qq.com', '18613030352', 1, 'admin', '2021-04-27 13:47:42', 'admin', '2022-09-04 17:49:04', NULL, NULL);
INSERT INTO `sys_user` VALUES (41, 'zhao', 'zhao', '', '1', '1', '1', 1, 'admin', '2022-09-04 15:31:29', 'admin', '2022-09-04 17:49:04', NULL, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
