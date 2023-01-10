/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80030
 Source Host           : 127.0.0.1:3306
 Source Schema         : chaos_system

 Target Server Type    : MySQL
 Target Server Version : 80030
 File Encoding         : 65001

 Date: 09/01/2023 23:14:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api
-- ----------------------------
DROP TABLE IF EXISTS `api`;
CREATE TABLE `api` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `api` varchar(255) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'api接口',
  `group` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'api分组',
  `describe` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'api描述信息',
  `method` varchar(50) COLLATE utf8mb4_general_ci NOT NULL COMMENT '方法',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_api_methos` (`api`,`method`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=128 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of api
-- ----------------------------
BEGIN;
INSERT INTO `api` VALUES (1, '/api/v1/system/menu/usermenus', '默认', '获取用户动态路由', 'GET');
INSERT INTO `api` VALUES (2, '/api/v1/system/user/currentset', '默认', '用户当前设置', 'GET');
INSERT INTO `api` VALUES (3, '/api/v1/system/user/password', '默认', '修改登陆用户密码', 'PUT');
INSERT INTO `api` VALUES (4, '/api/v1/system/user/page', '默认', '修改用户页面配置', 'PUT');
INSERT INTO `api` VALUES (5, '/api/v1/system/user/changerole', '默认', '切换用户角色', 'POST');
INSERT INTO `api` VALUES (6, '/api/v1/system/user/paging', '用户管理', '分页用户详情(read)', 'POST');
INSERT INTO `api` VALUES (7, '/api/v1/system/user/[0-9]+', '用户管理', '用户详情(read)', 'GET');
INSERT INTO `api` VALUES (8, '/api/v1/system/user/create', '用户管理', '创建用户(write)', 'POST');
INSERT INTO `api` VALUES (9, '/api/v1/system/user/[0-9]+', '用户管理', '更新用户(write)', 'PUT');
INSERT INTO `api` VALUES (10, '/api/v1/system/user/[0-9]+', '用户管理', '删除用户(write)', 'DELETE');
INSERT INTO `api` VALUES (11, '/api/v1/system/user/[0-9]+/soft', '用户管理', '禁用用户(write)', 'DELETE');
INSERT INTO `api` VALUES (12, '/api/v1/system/user/[0-9]+/role', '用户管理', '修改用户角色(write)', 'PUT');
INSERT INTO `api` VALUES (13, '/api/v1/system/user/[0-9]+/password', '用户管理', '修改用户密码(write)', 'PUT');
INSERT INTO `api` VALUES (14, '/api/v1/system/role/all', '角色管理', '全部角色(read)', 'GET');
INSERT INTO `api` VALUES (15, '/api/v1/system/role/[0-9]+', '角色管理', '角色详情(read)', 'GET');
INSERT INTO `api` VALUES (16, '/api/v1/system/api/all', '角色管理', '全部接口列表', 'GET');
INSERT INTO `api` VALUES (17, '/api/v1/system/role/menupermission/[0-9]+', '角色管理', '角色菜单权限(read)', 'GET');
INSERT INTO `api` VALUES (18, '/api/v1/system/role/apipermission/[0-9]+', '角色管理', '角色API权限(read)', 'GET');
INSERT INTO `api` VALUES (19, '/api/v1/system/role/create', '角色管理', '创建角色(write)', 'POST');
INSERT INTO `api` VALUES (20, '/api/v1/system/role/[0-9]+', '角色管理', '更新角色(write)', 'PUT');
INSERT INTO `api` VALUES (21, '/api/v1/system/role/[0-9]+', '角色管理', '删除角色(write)', 'DELETE');
INSERT INTO `api` VALUES (22, '/api/v1/system/role/[0-9]+/soft', '角色管理', '禁用角色(write)', 'DELETE');
INSERT INTO `api` VALUES (23, '/api/v1/system/role/refreshpermission', '角色管理', '刷新权限表(write)', 'GET');
INSERT INTO `api` VALUES (24, '/api/v1/system/role/menupermission/[0-9]+', '角色管理', '更新角色菜单(write)', 'PUT');
INSERT INTO `api` VALUES (25, '/api/v1/system/role/apipermission/[0-9]+', '角色管理', '更新角色API权限(write)', 'PUT');
INSERT INTO `api` VALUES (26, '/api/v1/system/menu/all', '菜单管理', '获取全部菜单(read)', 'GET');
INSERT INTO `api` VALUES (27, '/api/v1/system/menu/[0-9]+', '菜单管理', '菜单详情(read)', 'GET');
INSERT INTO `api` VALUES (28, '/api/v1/system/user/all', '菜单管理', '全部用户(read)', 'GET');
INSERT INTO `api` VALUES (29, '/api/v1/system/menu/create', '菜单管理', '创建菜单(wirte)', 'POST');
INSERT INTO `api` VALUES (30, '/api/v1/system/menu/[0-9]+', '菜单管理', '更新菜单(write)', 'PUT');
INSERT INTO `api` VALUES (31, '/api/v1/system/menu/[0-9]+', '菜单管理', '删除菜单(write)', 'DELETE');
INSERT INTO `api` VALUES (32, '/api/v1/system/menu/[0-9]+/userparam', '菜单管理', '更新菜单用户参数(write)', 'PUT');
INSERT INTO `api` VALUES (33, '/api/v1/system/api/paging', 'API管理', '获取分页API(read)', 'POST');
INSERT INTO `api` VALUES (34, '/api/v1/system/api/[0-9]+', 'API管理', 'API详情(read)', 'GET');
INSERT INTO `api` VALUES (35, '/api/v1/system/api/create', 'API管理', '创建API(write)', 'POST');
INSERT INTO `api` VALUES (36, '/api/v1/system/api/[0-9]+', 'API管理', '更新API(write)', 'PUT');
INSERT INTO `api` VALUES (37, '/api/v1/system/api/[0-9]+', 'API管理', '删除API(write)', 'DELETE');
INSERT INTO `api` VALUES (38, '/api/v1/system/api/multiple', 'API管理', '批量删除API(write)', 'DELETE');
INSERT INTO `api` VALUES (126, '/api/v1/esmanager/conn/paging', '连接管理', 'ES_连接管理(read)', 'POST');
INSERT INTO `api` VALUES (127, '/api/v1/esmanager/conn/ping/[0-9]+', '连接管理', '测试es连接(read)', 'GET');
COMMIT;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `p_type` varchar(32) NOT NULL DEFAULT '',
  `v0` varchar(255) NOT NULL DEFAULT '',
  `v1` varchar(255) NOT NULL DEFAULT '',
  `v2` varchar(255) NOT NULL DEFAULT '',
  `v3` varchar(255) NOT NULL DEFAULT '',
  `v4` varchar(255) NOT NULL DEFAULT '',
  `v5` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `idx_casbin_rule` (`p_type`,`v0`,`v1`)
) ENGINE=InnoDB AUTO_INCREMENT=238 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES (198, 'p', '1', '/api/v1/system/menu/usermenus', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (199, 'p', '1', '/api/v1/system/user/currentset', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (200, 'p', '1', '/api/v1/system/user/password', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (201, 'p', '1', '/api/v1/system/user/page', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (202, 'p', '1', '/api/v1/system/user/changerole', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (203, 'p', '1', '/api/v1/system/user/paging', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (204, 'p', '1', '/api/v1/system/user/[0-9]+', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (205, 'p', '1', '/api/v1/system/user/create', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (206, 'p', '1', '/api/v1/system/user/[0-9]+', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (207, 'p', '1', '/api/v1/system/user/[0-9]+', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (208, 'p', '1', '/api/v1/system/user/[0-9]+/soft', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (209, 'p', '1', '/api/v1/system/user/[0-9]+/role', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (210, 'p', '1', '/api/v1/system/user/[0-9]+/password', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (211, 'p', '1', '/api/v1/system/role/all', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (212, 'p', '1', '/api/v1/system/role/[0-9]+', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (213, 'p', '1', '/api/v1/system/api/all', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (214, 'p', '1', '/api/v1/system/role/menupermission/[0-9]+', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (215, 'p', '1', '/api/v1/system/role/apipermission/[0-9]+', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (216, 'p', '1', '/api/v1/system/role/create', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (217, 'p', '1', '/api/v1/system/role/[0-9]+', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (218, 'p', '1', '/api/v1/system/role/[0-9]+', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (219, 'p', '1', '/api/v1/system/role/[0-9]+/soft', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (220, 'p', '1', '/api/v1/system/role/refreshpermission', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (221, 'p', '1', '/api/v1/system/role/menupermission/[0-9]+', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (222, 'p', '1', '/api/v1/system/role/apipermission/[0-9]+', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (223, 'p', '1', '/api/v1/system/menu/all', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (224, 'p', '1', '/api/v1/system/menu/[0-9]+', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (225, 'p', '1', '/api/v1/system/user/all', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (226, 'p', '1', '/api/v1/system/menu/create', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (227, 'p', '1', '/api/v1/system/menu/[0-9]+', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (228, 'p', '1', '/api/v1/system/menu/[0-9]+', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (229, 'p', '1', '/api/v1/system/menu/[0-9]+/userparam', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (230, 'p', '1', '/api/v1/system/api/paging', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (231, 'p', '1', '/api/v1/system/api/[0-9]+', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (232, 'p', '1', '/api/v1/system/api/create', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (233, 'p', '1', '/api/v1/system/api/[0-9]+', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (234, 'p', '1', '/api/v1/system/api/[0-9]+', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (235, 'p', '1', '/api/v1/system/api/multiple', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (236, 'p', '1', '/api/v1/esmanager/conn/paging', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (237, 'p', '1', '/api/v1/esmanager/conn/ping/[0-9]+', 'GET', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for casbin_rule_copy1
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule_copy1`;
CREATE TABLE `casbin_rule_copy1` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `p_type` varchar(32) NOT NULL DEFAULT '',
  `v0` varchar(255) NOT NULL DEFAULT '',
  `v1` varchar(255) NOT NULL DEFAULT '',
  `v2` varchar(255) NOT NULL DEFAULT '',
  `v3` varchar(255) NOT NULL DEFAULT '',
  `v4` varchar(255) NOT NULL DEFAULT '',
  `v5` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `idx_casbin_rule` (`p_type`,`v0`,`v1`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of casbin_rule_copy1
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule_copy1` VALUES (1, 'p', '1', '/api/v1/system/menu/usermenus', 'GET', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (3, 'p', '1', '/api/v1/system/role/all', 'GET', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (6, 'p', '1', '/api/v1/system/user/create', 'POST', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (7, 'p', '1', '/api/v1/system/menu/all', 'GET', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (8, 'p', '1', '/api/v1/system/menu/create', 'POST', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (9, 'p', '1', '/api/v1/system/menu/[0-9]+', 'DELETE', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (12, 'p', '1', '/api/v1/system/menu/[0-9]+', 'GET', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (13, 'p', '1', '/api/v1/system/menu/[0-9]+', 'PUT', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (14, 'p', '1', '/api/v1/system/menu/[0-9]+/userparam', 'POST', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (15, 'p', '1', '/api/v1/system/api/paging', 'POST', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (16, 'p', '1', '/api/v1/system/api/[0-9]+', 'DELETE', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (17, 'p', '1', '/api/v1/system/api/create', 'POST', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (18, 'p', '1', '/api/v1/system/api/[0-9]+', 'PUT', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (19, 'p', '1', '/api/v1/system/api/[0-9]+', 'GET', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (20, 'p', '1', '/api/v1/system/user/[0-9]+', 'GET', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (21, 'p', '1', '/api/v1/system/user/[0-9]+', 'PUT', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (22, 'p', '1', '/api/v1/system/api/multiple', 'DELETE', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (23, 'p', '1', '/api/v1/system/user/all', 'GET', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (24, 'p', '1', '/api/v1/system/role/create', 'POST', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (25, 'p', '1', '/api/v1/system/role/refreshpermission', 'GET', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (26, 'p', '1', '/api/v1/system/user/page', 'PUT', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (27, 'p', '1', '/api/v1/system/menu/usermenus', 'GET', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (28, 'p', '1', '/api/v1/system/user/currentset', 'GET', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (29, 'p', '1', '/api/v1/system/user/paging', 'POST', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (30, 'p', '1', '/api/v1/system/role/all', 'GET', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (31, 'p', '1', '/api/v1/system/user/[0-9]+/role', 'PUT', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (32, 'p', '1', '/api/v1/system/user/[0-9]+/soft', 'DELETE', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (33, 'p', '1', '/api/v1/system/role/[0-9]+/soft', 'DELETE', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (34, 'p', '1', '/api/v1/system/role/[0-9]+', 'PUT', '', '', '');
INSERT INTO `casbin_rule_copy1` VALUES (35, 'p', '1', '/api/v1/system/role/[0-9]+', 'GET', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父菜单ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由name',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '路由path',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '对应前端文件路径',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '附加属性',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '附加属性',
  `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
  `hidden` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否隐藏 0 false/1 true',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_name` (`name`) USING BTREE,
  UNIQUE KEY `uq_path` (`path`) USING BTREE,
  KEY `idx_delete_at` (`delete_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (1, 0, 'dashboard', 'dashboard', 'view/dashboard/index.vue', '仪表盘', 'setting', 0, 0, '2022-11-28 10:58:26', '2022-12-20 10:21:58', NULL);
INSERT INTO `menu` VALUES (2, 0, 'superAdmin', 'superAdmin', 'view/superAdmin/index.vue', '超级管理员', 'set-up', 1, 0, '2022-09-04 21:32:00', '2022-12-21 02:22:37', NULL);
INSERT INTO `menu` VALUES (3, 2, 'user', 'user', 'view/superAdmin/user/user.vue', '用户管理', 'user-solid', 0, 0, '2022-09-04 21:33:08', '2022-12-21 02:22:56', NULL);
INSERT INTO `menu` VALUES (4, 2, 'menu', 'menu', 'view/superAdmin/menu/menu.vue', '菜单管理', 'menu', 2, 0, '2022-09-08 03:29:49', '2022-12-21 02:16:22', NULL);
INSERT INTO `menu` VALUES (5, 2, 'systemsetup', 'systemsetup', 'view/superAdmin/systemsetup/index.vue', '系统管理', 'setting', 7, 0, '2022-12-16 12:48:20', '2022-12-21 02:21:42', NULL);
INSERT INTO `menu` VALUES (6, 5, 'ldapsetup', 'ldapsetup', 'view/superAdmin/systemsetup/ldapsetup/ldapsetup.vue', 'ldap配置', 's-order', 1, 0, '2022-12-16 12:53:22', '2022-12-16 12:58:05', NULL);
INSERT INTO `menu` VALUES (7, 2, 'api', 'api', 'view/superAdmin/api/api.vue', 'api管理', 'c-scale-to-original', 3, 0, '2022-12-21 02:27:36', '2022-12-21 02:36:36', NULL);
INSERT INTO `menu` VALUES (8, 2, 'role', 'role', 'view/superAdmin/role/role.vue', '角色管理', 'user', 1, 0, '2022-12-23 10:54:48', '2023-01-03 10:11:57', NULL);
INSERT INTO `menu` VALUES (15, 0, 'persionhome', 'persionhome', 'view/home/index.vue', '家', 'toilet-paper', 99, 1, '2023-01-03 04:54:09', '2023-01-03 10:16:35', NULL);
INSERT INTO `menu` VALUES (16, 0, 'esManager', 'esManager', 'view/esManager/index.vue', 'ES管理', 's-help', 2, 0, '2023-01-08 21:35:23', '2023-01-08 21:35:39', NULL);
INSERT INTO `menu` VALUES (17, 16, 'esConn', 'esConn', 'view/esManager/connManager/connManager.vue', '连接管理', 'connection', 0, 0, '2023-01-08 21:40:48', '2023-01-08 21:40:48', NULL);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `role` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户角色',
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `create_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `delete_by` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `delete_time` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_role` (`role`) USING BTREE,
  KEY `idx_delete_at` (`delete_time`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, 'admin', '超级管理员', 'admin', '2019-01-18 11:11:11', 'admin', '2022-12-23 17:35:55', '', NULL);
INSERT INTO `role` VALUES (2, 'mng', '项目经理', 'admin', '2019-01-19 11:11:11', 'admin', '2022-12-25 12:28:26', 'admin', NULL);
INSERT INTO `role` VALUES (3, 'dev', '开发人员', 'admin', '2019-01-19 11:11:11', 'admin', '2022-11-28 10:45:52', '', NULL);
INSERT INTO `role` VALUES (4, 'test', '测试人员', 'admin', '2019-01-19 11:11:11', 'admin', '2022-11-28 10:45:54', '', NULL);
INSERT INTO `role` VALUES (5, 'demo', 'demo', 'admin', '2020-11-26 14:52:20', 'admin', '2022-12-25 12:17:35', '', NULL);
COMMIT;

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `menu_id` bigint unsigned NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_m_id` (`menu_id`),
  KEY `idx_menu_id` (`menu_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
BEGIN;
INSERT INTO `role_menu` VALUES (1, 1, 1);
INSERT INTO `role_menu` VALUES (2, 2, 1);
INSERT INTO `role_menu` VALUES (3, 3, 1);
INSERT INTO `role_menu` VALUES (4, 4, 1);
INSERT INTO `role_menu` VALUES (7, 7, 1);
INSERT INTO `role_menu` VALUES (8, 8, 1);
INSERT INTO `role_menu` VALUES (19, 1, 2);
INSERT INTO `role_menu` VALUES (24, 2, 2);
INSERT INTO `role_menu` VALUES (27, 13, 2);
INSERT INTO `role_menu` VALUES (31, 5, 2);
INSERT INTO `role_menu` VALUES (32, 6, 2);
INSERT INTO `role_menu` VALUES (33, 15, 1);
INSERT INTO `role_menu` VALUES (34, 16, 1);
INSERT INTO `role_menu` VALUES (35, 17, 1);
COMMIT;

-- ----------------------------
-- Table structure for system_setting
-- ----------------------------
DROP TABLE IF EXISTS `system_setting`;
CREATE TABLE `system_setting` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `type` varchar(50) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `setting` text COLLATE utf8mb4_general_ci,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_type` (`type`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of system_setting
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `nick_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '中文名',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `type` tinyint NOT NULL COMMENT '账户类型 0-本地用户 1-ldap用户',
  `email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '电话',
  `department` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '部门',
  `position` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '职位',
  `current_role` bigint unsigned NOT NULL DEFAULT '0' COMMENT '当前用户角色',
  `create_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '创建人',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '更新人',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '删除人',
  `delete_time` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  `page_set_id` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`) USING BTREE,
  KEY `idx_delete_time` (`delete_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 'admin', '超级管理员', '$2a$05$x0ATWmRq9yd.oAIO8T8DS.KB332PX17nT0MK5fVh6R/fYvQAFliOy', 0, 'admin@xxx.com', '10000000000', '', '', 1, 'admin', '2022-12-15 16:57:51', 'admin', '2023-01-08 16:59:39', 'test1', NULL, NULL);
INSERT INTO `user` VALUES (4, 'qiang.zhao', 'qiang.zhao', '$2a$05$pShRxnBlMwj3IN7NCVl80.x3805b3Dkxg5tV5W3yHhumTzKzMzqhW', 0, 'qiang.zhao@qq.com', '19903413813', '运维部', '', 1, 'admin', '2022-12-16 11:55:17', 'admin', '2022-12-27 00:30:07', 'admin', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for user_menu_params
-- ----------------------------
DROP TABLE IF EXISTS `user_menu_params`;
CREATE TABLE `user_menu_params` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL DEFAULT '0',
  `menu_id` bigint unsigned NOT NULL,
  `type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`) USING BTREE,
  KEY `idx_menu_id` (`menu_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user_menu_params
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_page_set
-- ----------------------------
DROP TABLE IF EXISTS `user_page_set`;
CREATE TABLE `user_page_set` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL DEFAULT '0',
  `avatar` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `default_router` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT 'dashboard' COMMENT '默认路由',
  `side_mode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '#191a23' COMMENT 'side颜色',
  `text_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '#fff' COMMENT '文本颜色',
  `active_text_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '#1890ff' COMMENT '选中路由文本颜色',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user_page_set
-- ----------------------------
BEGIN;
INSERT INTO `user_page_set` VALUES (1, 1, '', 'dashboard', '#212341', '#FFFFFF', '#1890FF');
INSERT INTO `user_page_set` VALUES (43, 2, '', 'dashboard', '#1226DD', '#ED0B0B', '#ECFF18');
COMMIT;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL,
  `role_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_role_id` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
INSERT INTO `user_role` VALUES (38, 2, 1);
INSERT INTO `user_role` VALUES (95, 1, 1);
INSERT INTO `user_role` VALUES (96, 1, 2);
INSERT INTO `user_role` VALUES (98, 4, 1);
INSERT INTO `user_role` VALUES (99, 4, 2);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
