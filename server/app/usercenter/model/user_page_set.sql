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

 Date: 28/11/2022 17:27:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_page_set
-- ----------------------------
DROP TABLE IF EXISTS `user_page_set`;
CREATE TABLE `user_page_set` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL DEFAULT 0,
  `avatar` varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `default_router` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'dashboard' COMMENT '默认路由',
  `side_mode` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '#191a23' COMMENT 'side颜色',
  `text_color` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '#fff' COMMENT '文本颜色',
  `active_text_color` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '#1890ff' COMMENT '选中路由文本颜色',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
