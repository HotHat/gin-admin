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

 Date: 09/09/2025 14:37:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for logger
-- ----------------------------
DROP TABLE IF EXISTS `logger`;
CREATE TABLE `logger`  (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `level` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `trace_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `user_id` int(10) NOT NULL,
  `tag` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `message` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `stack` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `data` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_logger_level`(`level`) USING BTREE,
  INDEX `idx_logger_trace_id`(`trace_id`) USING BTREE,
  INDEX `idx_logger_user_id`(`user_id`) USING BTREE,
  INDEX `idx_logger_tag`(`tag`) USING BTREE,
  INDEX `idx_logger_created_at`(`created_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of logger
-- ----------------------------
INSERT INTO `logger` VALUES (1, 'info', '', 0, 'main', 'starting service ...', '', '{\"pid\":43580,\"workdir\":\"configs\",\"config\":\"dev\",\"static\":\"\",\"version\":\"v10.1.0\"}', '2025-09-05 12:33:18.176');
INSERT INTO `logger` VALUES (2, 'info', '', 0, 'main', 'HTTP server is listening on :8040', '', '', '2025-09-05 12:33:18.519');
INSERT INTO `logger` VALUES (3, 'info', 'TRACE-D2T6GIFINHHAKF66VA1G', 0, 'request', '[HTTP] /api/health-GET-401 (0ms)', '', '{\"path\":\"/api/health\",\"user_agent\":\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36\",\"remote_addr\":\"127.0.0.1:2509\",\"proto\":\"HTTP/1.1\",\"content_length\":0,\"status\":401,\"res_size\":119,\"host\":\"127.0.0.1:8040\",\"method\":\"GET\",\"referer\":\"\",\"res_time\":\"2025-09-05 12:34:17.548\",\"client_ip\":\"127.0.0.1\",\"uri\":\"/api/health\",\"content_type\":\"\",\"pragma\":\"\",\"cost\":0,\"res_body\":\"{\\\"success\\\":false,\\\"error\\\":{\\\"id\\\":\\\"com.invalid.token\\\",\\\"code\\\":401,\\\"detail\\\":\\\"Invalid access token\\\",\\\"status\\\":\\\"Unauthorized\\\"}}\"}', '2025-09-05 12:34:17.548');
INSERT INTO `logger` VALUES (4, 'info', 'TRACE-D2T6GP7INHHAKF66VA2G', 0, 'request', '[HTTP] /api/health-GET-401 (0ms)', '', '{\"client_ip\":\"127.0.0.1\",\"host\":\"127.0.0.1:8040\",\"remote_addr\":\"127.0.0.1:2539\",\"proto\":\"HTTP/1.1\",\"uri\":\"/api/health\",\"pragma\":\"no-cache\",\"status\":401,\"res_size\":119,\"res_body\":\"{\\\"success\\\":false,\\\"error\\\":{\\\"id\\\":\\\"com.invalid.token\\\",\\\"code\\\":401,\\\"detail\\\":\\\"Invalid access token\\\",\\\"status\\\":\\\"Unauthorized\\\"}}\",\"method\":\"GET\",\"referer\":\"\",\"content_type\":\"\",\"cost\":0,\"res_time\":\"2025-09-05 12:34:44.905\",\"path\":\"/api/health\",\"user_agent\":\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36\",\"content_length\":0}', '2025-09-05 12:34:44.905');
INSERT INTO `logger` VALUES (5, 'info', 'TRACE-D2T6GQ7INHHAKF66VA3G', 0, 'request', '[HTTP] /api/health-GET-401 (0ms)', '', '{\"uri\":\"/api/health\",\"content_length\":0,\"res_time\":\"2025-09-05 12:34:48.956\",\"client_ip\":\"127.0.0.1\",\"path\":\"/api/health\",\"host\":\"127.0.0.1:8040\",\"pragma\":\"no-cache\",\"res_size\":119,\"method\":\"GET\",\"user_agent\":\"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36\",\"proto\":\"HTTP/1.1\",\"content_type\":\"\",\"status\":401,\"remote_addr\":\"127.0.0.1:2544\",\"cost\":0,\"res_body\":\"{\\\"success\\\":false,\\\"error\\\":{\\\"id\\\":\\\"com.invalid.token\\\",\\\"code\\\":401,\\\"detail\\\":\\\"Invalid access token\\\",\\\"status\\\":\\\"Unauthorized\\\"}}\",\"referer\":\"\"}', '2025-09-05 12:34:48.956');
INSERT INTO `logger` VALUES (6, 'info', '', 0, 'main', 'Received signal', '', '{\"signal\":\"interrupt\"}', '2025-09-05 13:30:52.079');
INSERT INTO `logger` VALUES (7, 'info', '', 0, 'main', 'starting service ...', '', '{\"version\":\"v10.1.0\",\"pid\":41600,\"config\":\"dev\",\"static\":\"\",\"workdir\":\"configs\"}', '2025-09-05 17:14:28.149');
INSERT INTO `logger` VALUES (8, 'info', '', 0, 'main', 'HTTP server is listening on :8040', '', '', '2025-09-05 17:14:28.338');
INSERT INTO `logger` VALUES (9, 'info', '', 0, 'main', 'Received signal', '', '{\"signal\":\"interrupt\"}', '2025-09-05 17:15:16.802');

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
INSERT INTO `menu` VALUES (1, 'home', 'Home', '', 90, 'page', '/home', '', 'enabled', 0, '', '2025-09-05 12:33:18.448', '2025-09-05 12:33:18.448');
INSERT INTO `menu` VALUES (2, 'system', 'System', '', 10, 'page', '/system', '', 'enabled', 0, '', '2025-09-05 12:33:18.449', '2025-09-05 12:33:18.449');
INSERT INTO `menu` VALUES (3, 'menu', 'Menu', '', 90, 'page', '/system/menu', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.', '2025-09-05 12:33:18.461', '2025-09-05 12:33:18.461');
INSERT INTO `menu` VALUES (4, 'add', 'Add', '', 9, 'button', '', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.d2t6g3ninhhakf66v9g0.', '2025-09-05 12:33:18.476', '2025-09-05 12:33:18.476');
INSERT INTO `menu` VALUES (5, 'edit', 'Edit', '', 8, 'button', '', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.d2t6g3ninhhakf66v9g0.', '2025-09-05 12:33:18.478', '2025-09-05 12:33:18.478');
INSERT INTO `menu` VALUES (6, 'delete', 'Delete', '', 7, 'button', '', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.d2t6g3ninhhakf66v9g0.', '2025-09-05 12:33:18.479', '2025-09-05 12:33:18.479');
INSERT INTO `menu` VALUES (7, 'search', 'Search', '', 6, 'button', '', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.d2t6g3ninhhakf66v9g0.', '2025-09-05 12:33:18.480', '2025-09-05 12:33:18.480');
INSERT INTO `menu` VALUES (8, 'role', 'Role', '', 80, 'page', '/system/role', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.', '2025-09-05 12:33:18.480', '2025-09-05 12:33:18.480');
INSERT INTO `menu` VALUES (9, 'add', 'Add', '', 9, 'button', '', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.d2t6g3ninhhakf66v9l0.', '2025-09-05 12:33:18.483', '2025-09-05 12:33:18.483');
INSERT INTO `menu` VALUES (10, 'edit', 'Edit', '', 8, 'button', '', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.d2t6g3ninhhakf66v9l0.', '2025-09-05 12:33:18.484', '2025-09-05 12:33:18.484');
INSERT INTO `menu` VALUES (11, 'delete', 'Delete', '', 7, 'button', '', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.d2t6g3ninhhakf66v9l0.', '2025-09-05 12:33:18.485', '2025-09-05 12:33:18.485');
INSERT INTO `menu` VALUES (12, 'search', 'Search', '', 6, 'button', '', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.d2t6g3ninhhakf66v9l0.', '2025-09-05 12:33:18.486', '2025-09-05 12:33:18.486');
INSERT INTO `menu` VALUES (13, 'user', 'User', '', 70, 'page', '/system/user', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.', '2025-09-05 12:33:18.487', '2025-09-05 12:33:18.487');
INSERT INTO `menu` VALUES (14, 'add', 'Add', '', 9, 'button', '', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.d2t6g3ninhhakf66v9qg.', '2025-09-05 12:33:18.488', '2025-09-05 12:33:18.488');
INSERT INTO `menu` VALUES (15, 'edit', 'Edit', '', 8, 'button', '', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.d2t6g3ninhhakf66v9qg.', '2025-09-05 12:33:18.489', '2025-09-05 12:33:18.489');
INSERT INTO `menu` VALUES (16, 'delete', 'Delete', '', 7, 'button', '', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.d2t6g3ninhhakf66v9qg.', '2025-09-05 12:33:18.491', '2025-09-05 12:33:18.491');
INSERT INTO `menu` VALUES (17, 'search', 'Search', '', 6, 'button', '', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.d2t6g3ninhhakf66v9qg.', '2025-09-05 12:33:18.492', '2025-09-05 12:33:18.492');
INSERT INTO `menu` VALUES (18, 'logger', 'Logger', '', 10, 'page', '/system/logger', '', 'enabled', 0, 'd2t6g3ninhhakf66v9fg.', '2025-09-05 12:33:18.493', '2025-09-05 12:33:18.493');

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
INSERT INTO `menu_resource` VALUES (1, 1, 'GET', '/api/v1/menus', '2025-09-05 12:33:18.463', '2025-09-05 12:33:18.463');
INSERT INTO `menu_resource` VALUES (2, 1, 'GET', '/api/v1/menus/{id}', '2025-09-05 12:33:18.464', '2025-09-05 12:33:18.464');
INSERT INTO `menu_resource` VALUES (3, 2, 'POST', '/api/v1/menus', '2025-09-05 12:33:18.477', '2025-09-05 12:33:18.477');
INSERT INTO `menu_resource` VALUES (4, 2, 'PUT', '/api/v1/menus/{id}', '2025-09-05 12:33:18.478', '2025-09-05 12:33:18.478');
INSERT INTO `menu_resource` VALUES (5, 3, 'DELETE', '/api/v1/menus/{id}', '2025-09-05 12:33:18.479', '2025-09-05 12:33:18.479');
INSERT INTO `menu_resource` VALUES (6, 3, 'GET', '/api/v1/menus', '2025-09-05 12:33:18.481', '2025-09-05 12:33:18.481');
INSERT INTO `menu_resource` VALUES (7, 4, 'GET', '/api/v1/roles', '2025-09-05 12:33:18.482', '2025-09-05 12:33:18.482');
INSERT INTO `menu_resource` VALUES (8, 4, 'GET', '/api/v1/roles/{id}', '2025-09-05 12:33:18.482', '2025-09-05 12:33:18.482');
INSERT INTO `menu_resource` VALUES (9, 5, 'POST', '/api/v1/roles', '2025-09-05 12:33:18.483', '2025-09-05 12:33:18.483');
INSERT INTO `menu_resource` VALUES (10, 5, 'PUT', '/api/v1/roles/{id}', '2025-09-05 12:33:18.484', '2025-09-05 12:33:18.484');
INSERT INTO `menu_resource` VALUES (11, 5, 'DELETE', '/api/v1/roles/{id}', '2025-09-05 12:33:18.485', '2025-09-05 12:33:18.485');
INSERT INTO `menu_resource` VALUES (12, 5, 'GET', '/api/v1/roles', '2025-09-05 12:33:18.487', '2025-09-05 12:33:18.487');
INSERT INTO `menu_resource` VALUES (13, 6, 'GET', '/api/v1/users', '2025-09-05 12:33:18.488', '2025-09-05 12:33:18.488');
INSERT INTO `menu_resource` VALUES (14, 6, 'GET', '/api/v1/users/{id}', '2025-09-05 12:33:18.488', '2025-09-05 12:33:18.488');
INSERT INTO `menu_resource` VALUES (15, 6, 'POST', '/api/v1/users', '2025-09-05 12:33:18.489', '2025-09-05 12:33:18.489');
INSERT INTO `menu_resource` VALUES (16, 6, 'PUT', '/api/v1/users/{id}', '2025-09-05 12:33:18.490', '2025-09-05 12:33:18.490');
INSERT INTO `menu_resource` VALUES (17, 6, 'DELETE', '/api/v1/users/{id}', '2025-09-05 12:33:18.492', '2025-09-05 12:33:18.492');
INSERT INTO `menu_resource` VALUES (18, 6, 'GET', '/api/v1/loggers', '2025-09-05 12:33:18.493', '2025-09-05 12:33:18.493');

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
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
INSERT INTO `role_menu` VALUES (1, 1, 1, '2025-09-09 09:23:25.000', '2025-09-09 09:23:27.000');
INSERT INTO `role_menu` VALUES (2, 1, 2, '2025-09-09 09:24:15.000', '2025-09-09 09:24:20.000');
INSERT INTO `role_menu` VALUES (3, 1, 3, '2025-09-09 09:24:17.000', '2025-09-09 09:24:23.000');

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
