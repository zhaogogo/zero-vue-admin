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

 Date: 02/09/2022 01:19:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for system_user
-- ----------------------------
DROP TABLE IF EXISTS `system_user`;
CREATE TABLE `system_user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '编号',
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `nick_name` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '昵称',
  `avatar` varchar(150) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
  `mobile` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '手机号',
  `status` tinyint(3) unsigned zerofill NOT NULL DEFAULT '001' COMMENT '状态  0：禁用   1：正常',
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '更新人',
  `update_time` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `del_flag` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除  0：已删除  1：正常',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户管理';

-- ----------------------------
-- Records of system_user
-- ----------------------------
BEGIN;
INSERT INTO `system_user` VALUES (1, 'admin', '超管', '', '123456', 'admin@qq.com', '13612345678', 001, 'admin', '2018-08-14 11:11:11', 'admin', '2018-08-14 11:11:11', 0);
INSERT INTO `system_user` VALUES (22, 'liubei', '刘备', '', '123456', 'test@qq.com', '13889700023', 001, 'admin', '2018-09-23 19:43:00', 'admin', '2019-01-10 11:41:13', 0);
INSERT INTO `system_user` VALUES (23, 'zhaoyun', '赵云', '', '123456', 'test@qq.com', '13889700023', 001, 'admin', '2018-09-23 19:43:44', 'admin', '2018-09-23 19:43:52', 0);
INSERT INTO `system_user` VALUES (24, 'zhugeliang', '诸葛亮', '', '123456', 'test@qq.com', '13889700023', 007, 'admin', '2018-09-23 19:44:23', 'admin', '2018-09-23 19:44:29', 0);
INSERT INTO `system_user` VALUES (25, 'caocao', '曹操', '', '123456', 'test@qq.com', '13889700023', 001, 'admin', '2018-09-23 19:45:32', 'admin', '2019-01-10 17:59:14', 0);
INSERT INTO `system_user` VALUES (26, 'dianwei', '典韦', '', '123456', 'test@qq.com', '13889700023', 001, 'admin', '2018-09-23 19:45:48', 'admin', '2018-09-23 19:45:57', 0);
INSERT INTO `system_user` VALUES (27, 'xiahoudun', '夏侯惇', '', '123456', 'test@qq.com', '13889700023', 001, 'admin', '2018-09-23 19:46:09', 'admin', '2018-09-23 19:46:17', 0);
INSERT INTO `system_user` VALUES (28, 'xunyu', '荀彧', '', '123456', 'test@qq.com', '13889700023', 001, 'admin', '2018-09-23 19:46:38', 'admin', '2018-11-04 15:33:17', 0);
INSERT INTO `system_user` VALUES (29, 'sunquan', '孙权', '', '123456', 'test@qq.com', '13889700023', 001, 'admin', '2018-09-23 19:46:54', 'admin', '2018-09-23 19:47:03', 0);
INSERT INTO `system_user` VALUES (30, 'zhouyu', '周瑜', '', '123456', 'test@qq.com', '13889700023', 001, 'admin', '2018-09-23 19:47:28', 'admin', '2018-09-23 19:48:04', 0);
INSERT INTO `system_user` VALUES (31, 'luxun', '陆逊', '', '123456', 'test@qq.com', '13889700023', 001, 'admin', '2018-09-23 19:47:44', 'admin', '2018-09-23 19:47:58', 0);
INSERT INTO `system_user` VALUES (32, 'huanggai', '黄盖', '', '', 'test@qq.com', '13889700023', 001, '', '2018-09-23 19:48:38', 'admin', '2021-04-03 15:43:52', 0);
INSERT INTO `system_user` VALUES (33, '1', '1', '', '123456', '1', '1', 001, 'admin', '2021-04-26 17:57:50', 'admin', '2021-04-26 17:57:50', 0);
INSERT INTO `system_user` VALUES (35, '12', '1', '', '123456', '1', '1', 001, 'admin', '2021-04-26 18:01:53', 'admin', '2021-04-26 18:01:54', 0);
INSERT INTO `system_user` VALUES (36, '12313', '12', '', '123456', '1', '1', 001, 'admin', '2021-04-26 18:03:07', 'admin', '2021-04-26 18:03:07', 0);
INSERT INTO `system_user` VALUES (37, '324', '1', '', '123456', '1', '1', 001, 'admin', '2021-04-26 18:07:31', 'admin', '2021-04-26 18:07:32', 0);
INSERT INTO `system_user` VALUES (38, 'aa', 'aa', '', '123456', 'a', 'a', 001, 'admin', '2021-04-27 11:24:14', 'admin', '2021-04-27 11:24:14', 0);
INSERT INTO `system_user` VALUES (39, '133', '133', '', '', '1121', '1', 001, 'admin', '2021-04-27 12:30:15', 'admin', '2021-04-27 13:53:40', 0);
INSERT INTO `system_user` VALUES (40, 'liu', 'liu', '', '123456', '1002219331@qq.com', '18613030352', 001, 'admin', '2021-04-27 13:47:42', 'admin', '2021-04-27 13:47:42', 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
