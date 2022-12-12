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

 Date: 28/11/2022 04:43:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `nick_name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '中文名',
  `password` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `type` tinyint NOT NULL DEFAULT 1 COMMENT '账户类型 0-本地用户 1-ldap用户',
  `email` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '电话',
  `department` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '部门',
  `position` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '职位',
  `avatar` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `default_router` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'dashboard' COMMENT '默认路由',
  `side_mode` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '#fff' COMMENT 'side颜色',
  `text_color` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '#fff' COMMENT '文本颜色',
  `active_text_color` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '#1890ff' COMMENT '选中路由文本颜色',
  `create_by` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新人',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '删除人',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`) USING BTREE,
  KEY `idx_delete_time` (`delete_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
