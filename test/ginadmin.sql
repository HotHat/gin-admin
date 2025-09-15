/*
 Navicat Premium Data Transfer

 Source Server         : $$$
 Source Server Type    : MySQL
 Source Server Version : 100432
 Source Host           : localhost:3306
 Source Schema         : ginadmin

 Target Server Type    : MySQL
 Target Server Version : 100432
 File Encoding         : 65001

 Date: 15/09/2025 10:48:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for logger
-- ----------------------------
DROP TABLE IF EXISTS `logger`;
CREATE TABLE `logger`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `level` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `trace_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `tag` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `message` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `stack` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_logger_level`(`level`) USING BTREE,
  INDEX `idx_logger_trace_id`(`trace_id`) USING BTREE,
  INDEX `idx_logger_user_id`(`user_id`) USING BTREE,
  INDEX `idx_logger_tag`(`tag`) USING BTREE,
  INDEX `idx_logger_created_at`(`created_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 170 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of logger
-- ----------------------------
-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `sequence` bigint(20) NOT NULL,
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `properties` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `status` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `parent_id` int(10) NOT NULL DEFAULT 0,
  `parent_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_menu_type`(`type`) USING BTREE,
  INDEX `idx_menu_status`(`status`) USING BTREE,
  INDEX `idx_menu_parent_id`(`parent_id`) USING BTREE,
  INDEX `idx_menu_code`(`code`) USING BTREE,
  INDEX `idx_menu_sequence`(`sequence`) USING BTREE,
  INDEX `idx_menu_parent_path`(`parent_path`) USING BTREE,
  INDEX `idx_menu_created_at`(`created_at`) USING BTREE,
  INDEX `idx_menu_updated_at`(`updated_at`) USING BTREE,
  INDEX `idx_menu_name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu
-- ----------------------------
INSERT INTO `menu` VALUES (1, 'home', '首页', '', 90, 'page', 'dashboard_console', '', '1', 0, '', '2025-09-05 12:33:18.448', '2025-09-05 12:33:18.448');
INSERT INTO `menu` VALUES (2, 'system', '系统设置', '', 10, 'page', '/system', '', '1', 0, '', '2025-09-05 12:33:18.449', '2025-09-05 12:33:18.449');
INSERT INTO `menu` VALUES (3, 'menu', '菜单管理', '', 90, 'page', '/system/menu', '', '1', 2, '2.', '2025-09-05 12:33:18.461', '2025-09-05 12:33:18.461');
INSERT INTO `menu` VALUES (4, 'add', '添加', '', 9, 'button', '/system/menu/add', '', '1', 3, '2.3.', '2025-09-05 12:33:18.476', '2025-09-05 12:33:18.476');
INSERT INTO `menu` VALUES (5, 'edit', '编辑', '', 8, 'button', '/system/menu/edit', '', '1', 3, '2.3.', '2025-09-05 12:33:18.478', '2025-09-05 12:33:18.478');
INSERT INTO `menu` VALUES (6, 'delete', '删除', '', 7, 'button', '/system/menu/delete', '', '1', 3, '2.3.', '2025-09-05 12:33:18.479', '2025-09-05 12:33:18.479');
INSERT INTO `menu` VALUES (7, 'search', '列表', '', 6, 'button', '/system/menu/list', '', '1', 3, '2.3.', '2025-09-05 12:33:18.480', '2025-09-05 12:33:18.480');
INSERT INTO `menu` VALUES (8, 'role', '角色管理', '', 80, 'page', '/system/role', '', '1', 2, '2.', '2025-09-05 12:33:18.480', '2025-09-05 12:33:18.480');
INSERT INTO `menu` VALUES (9, 'add', '添加', '', 9, 'button', '/system/role/add', '', '1', 8, '2.8.', '2025-09-05 12:33:18.483', '2025-09-05 12:33:18.483');
INSERT INTO `menu` VALUES (10, 'edit', '编辑', '', 8, 'button', '/system/role/edit', '', '1', 8, '2.8.', '2025-09-05 12:33:18.484', '2025-09-05 12:33:18.484');
INSERT INTO `menu` VALUES (11, 'delete', '删除', '', 7, 'button', '/system/role/delete', '', '1', 8, '2.8.', '2025-09-05 12:33:18.485', '2025-09-05 12:33:18.485');
INSERT INTO `menu` VALUES (12, 'search', '列表', '', 6, 'button', '/system/role/list', '', '1', 8, '2.8.', '2025-09-05 12:33:18.486', '2025-09-05 12:33:18.486');
INSERT INTO `menu` VALUES (13, 'user', '用户管理', '', 70, 'page', '/system/user', '', '1', 2, '2.', '2025-09-05 12:33:18.487', '2025-09-05 12:33:18.487');
INSERT INTO `menu` VALUES (14, 'add', '添加', '', 9, 'button', '/system/user/add', '', '1', 13, '2.13.', '2025-09-05 12:33:18.488', '2025-09-05 12:33:18.488');
INSERT INTO `menu` VALUES (15, 'edit', '编辑', '', 8, 'button', '/system/user/edit', '', '1', 13, '2.13.', '2025-09-05 12:33:18.489', '2025-09-05 12:33:18.489');
INSERT INTO `menu` VALUES (16, 'delete', '删除', '', 7, 'button', '/system/user/delete', '', '1', 13, '2.13.', '2025-09-05 12:33:18.491', '2025-09-05 12:33:18.491');
INSERT INTO `menu` VALUES (17, 'search', '列表', '', 6, 'button', '/system/user/list', '', '1', 13, '2.13.', '2025-09-05 12:33:18.492', '2025-09-05 12:33:18.492');
INSERT INTO `menu` VALUES (18, 'logger', '日志', '', 10, 'page', '/system/logger', '', '1', 2, '2.', '2025-09-05 12:33:18.493', '2025-09-05 12:33:18.493');

-- ----------------------------
-- Table structure for menu_resource
-- ----------------------------
DROP TABLE IF EXISTS `menu_resource`;
CREATE TABLE `menu_resource`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `menu_id` int(10) NOT NULL DEFAULT 0,
  `method` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_menu_resource_menu_id`(`menu_id`) USING BTREE,
  INDEX `idx_menu_resource_created_at`(`created_at`) USING BTREE,
  INDEX `idx_menu_resource_updated_at`(`updated_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of menu_resource
-- ----------------------------
INSERT INTO `menu_resource` VALUES (1, 3, 'GET', '/api/v1/menus', '2025-09-05 12:33:18.463', '2025-09-05 12:33:18.463');
INSERT INTO `menu_resource` VALUES (2, 3, 'GET', '/api/v1/menus/{id}', '2025-09-05 12:33:18.464', '2025-09-05 12:33:18.464');
INSERT INTO `menu_resource` VALUES (3, 4, 'POST', '/api/v1/menus', '2025-09-05 12:33:18.477', '2025-09-05 12:33:18.477');
INSERT INTO `menu_resource` VALUES (4, 5, 'PUT', '/api/v1/menus/{id}', '2025-09-05 12:33:18.478', '2025-09-05 12:33:18.478');
INSERT INTO `menu_resource` VALUES (5, 6, 'DELETE', '/api/v1/menus/{id}', '2025-09-05 12:33:18.479', '2025-09-05 12:33:18.479');
INSERT INTO `menu_resource` VALUES (6, 8, 'GET', '/api/v1/menus', '2025-09-05 12:33:18.481', '2025-09-05 12:33:18.481');
INSERT INTO `menu_resource` VALUES (7, 8, 'GET', '/api/v1/roles', '2025-09-05 12:33:18.482', '2025-09-05 12:33:18.482');
INSERT INTO `menu_resource` VALUES (8, 8, 'GET', '/api/v1/roles/{id}', '2025-09-05 12:33:18.482', '2025-09-05 12:33:18.482');
INSERT INTO `menu_resource` VALUES (9, 9, 'POST', '/api/v1/roles', '2025-09-05 12:33:18.483', '2025-09-05 12:33:18.483');
INSERT INTO `menu_resource` VALUES (10, 10, 'PUT', '/api/v1/roles/{id}', '2025-09-05 12:33:18.484', '2025-09-05 12:33:18.484');
INSERT INTO `menu_resource` VALUES (11, 11, 'DELETE', '/api/v1/roles/{id}', '2025-09-05 12:33:18.485', '2025-09-05 12:33:18.485');
INSERT INTO `menu_resource` VALUES (12, 13, 'GET', '/api/v1/roles', '2025-09-05 12:33:18.487', '2025-09-05 12:33:18.487');
INSERT INTO `menu_resource` VALUES (13, 13, 'GET', '/api/v1/users', '2025-09-05 12:33:18.488', '2025-09-05 12:33:18.488');
INSERT INTO `menu_resource` VALUES (14, 13, 'GET', '/api/v1/users/{id}', '2025-09-05 12:33:18.488', '2025-09-05 12:33:18.488');
INSERT INTO `menu_resource` VALUES (15, 14, 'POST', '/api/v1/users', '2025-09-05 12:33:18.489', '2025-09-05 12:33:18.489');
INSERT INTO `menu_resource` VALUES (16, 15, 'PUT', '/api/v1/users/{id}', '2025-09-05 12:33:18.490', '2025-09-05 12:33:18.490');
INSERT INTO `menu_resource` VALUES (17, 16, 'DELETE', '/api/v1/users/{id}', '2025-09-05 12:33:18.492', '2025-09-05 12:33:18.492');
INSERT INTO `menu_resource` VALUES (18, 18, 'GET', '/api/v1/loggers', '2025-09-05 12:33:18.493', '2025-09-05 12:33:18.493');

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `code` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `description` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `sequence` bigint(20) NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT 1,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_role_code`(`code`) USING BTREE,
  INDEX `idx_role_name`(`name`) USING BTREE,
  INDEX `idx_role_sequence`(`sequence`) USING BTREE,
  INDEX `idx_role_status`(`status`) USING BTREE,
  INDEX `idx_role_created_at`(`created_at`) USING BTREE,
  INDEX `idx_role_updated_at`(`updated_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role
-- ----------------------------
INSERT INTO `role` VALUES (1, '111', 'admin', 'root user', 1, 1, '2025-09-09 09:23:02.000', '2025-09-09 09:23:05.000');

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `role_id` int(10) NOT NULL,
  `menu_id` int(10) NOT NULL,
  `created_at` datetime(3) NOT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_role_menu_role_id`(`role_id`) USING BTREE,
  INDEX `idx_role_menu_menu_id`(`menu_id`) USING BTREE,
  INDEX `idx_role_menu_created_at`(`created_at`) USING BTREE,
  INDEX `idx_role_menu_updated_at`(`updated_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `role_menu` VALUES (1, 1, 1, '2025-09-12 11:05:53.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (2, 1, 2, '2025-09-12 11:05:53.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (3, 1, 3, '2025-09-12 11:05:53.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (4, 1, 4, '2025-09-12 11:05:53.000', '2025-09-12 11:05:51.000');
INSERT INTO `role_menu` VALUES (5, 1, 5, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (6, 1, 6, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (7, 1, 7, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (8, 1, 8, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (9, 1, 9, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (10, 1, 10, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (11, 1, 11, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (12, 1, 12, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (13, 1, 13, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (14, 1, 14, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (15, 1, 15, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (16, 1, 16, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (17, 1, 17, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');
INSERT INTO `role_menu` VALUES (18, 1, 18, '2025-09-12 11:06:05.000', '2025-09-12 11:05:53.000');

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `phone` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `email` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `remark` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `status` tinyint(4) NOT NULL DEFAULT 1,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_name`(`name`) USING BTREE,
  INDEX `idx_user_status`(`status`) USING BTREE,
  INDEX `idx_user_created_at`(`created_at`) USING BTREE,
  INDEX `idx_user_updated_at`(`updated_at`) USING BTREE,
  INDEX `idx_user_username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'admin', 'admin', '123456', '1520000002', 'abc@qq.com', 'test', 1, '2025-09-08 15:46:40.000', '2025-09-08 15:46:42.000');
INSERT INTO `user` VALUES (5, 'hello', 'world', '$2a$10$do/nEfDTfIChxoeC3D1NAOTCIfVOY7PZ43T7LZ6H/QmIyzpZSzG6W', '13800138000', 'aaa', '', 1, '2025-09-09 13:46:46.062', '2025-09-09 13:46:46.118');

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` int(10) NOT NULL DEFAULT 0,
  `role_id` int(10) NOT NULL DEFAULT 0,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_role_user_id`(`user_id`) USING BTREE,
  INDEX `idx_user_role_role_id`(`role_id`) USING BTREE,
  INDEX `idx_user_role_created_at`(`created_at`) USING BTREE,
  INDEX `idx_user_role_updated_at`(`updated_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role` VALUES (1, 1, 1, '2025-09-09 09:22:38.000', '2025-09-09 09:22:40.000');
INSERT INTO `user_role` VALUES (2, 5, 1, '2025-09-09 13:46:46.123', '2025-09-09 13:46:46.061');

SET FOREIGN_KEY_CHECKS = 1;
