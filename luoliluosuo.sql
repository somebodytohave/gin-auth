/*
 Navicat Premium Data Transfer

 Source Server         : mecm
 Source Server Type    : MySQL
 Source Server Version : 50553
 Source Host           : localhost:3306
 Source Schema         : luoliluosuo

 Target Server Type    : MySQL
 Target Server Version : 50553
 File Encoding         : 65001

 Date: 25/11/2018 15:13:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `created_by` int(11) NOT NULL,
  `topic_id` int(11) NULL DEFAULT NULL,
  `state` int(2) NULL DEFAULT 0,
  `created_at` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `modified_by` int(11) NULL DEFAULT NULL,
  `modified_at` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_l_article_l_user_1`(`created_by`) USING BTREE,
  INDEX `fk_l_article_l_topic_1`(`topic_id`) USING BTREE,
  CONSTRAINT `fk_l_article_l_topic_1` FOREIGN KEY (`topic_id`) REFERENCES `topic` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_l_article_l_user_1` FOREIGN KEY (`created_by`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '文章对应多个主题下的多个标签' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (1, '232', '433', 'ccc', 1, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `article` VALUES (3, '455', 'ttt', 'fff', 1, NULL, NULL, NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for article_tags
-- ----------------------------
DROP TABLE IF EXISTS `article_tags`;
CREATE TABLE `article_tags`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NULL DEFAULT NULL,
  `tag_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_article_tags_l_article_1`(`article_id`) USING BTREE,
  INDEX `fk_article_tags_l_tag_1`(`tag_id`) USING BTREE,
  CONSTRAINT `fk_article_tags_l_article_1` FOREIGN KEY (`article_id`) REFERENCES `article` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_article_tags_l_tag_1` FOREIGN KEY (`tag_id`) REFERENCES `tag` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of article_tags
-- ----------------------------
INSERT INTO `article_tags` VALUES (1, 1, 2);
INSERT INTO `article_tags` VALUES (2, 3, 1);

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '权限名称',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户权限表' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `state` int(2) NULL DEFAULT 0,
  `topic_id` int(11) NULL DEFAULT NULL,
  `created_by` int(11) NULL DEFAULT NULL,
  `created_at` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `modified_by` int(11) NULL DEFAULT NULL,
  `modified_at` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_l_tag_l_topic_1`(`topic_id`) USING BTREE,
  CONSTRAINT `fk_l_tag_l_topic_1` FOREIGN KEY (`topic_id`) REFERENCES `topic` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '主题分类下的标签' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of tag
-- ----------------------------
INSERT INTO `tag` VALUES (1, '123', NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `tag` VALUES (2, '45', NULL, NULL, NULL, NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for topic
-- ----------------------------
DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `state` int(2) NULL DEFAULT 0,
  `created_by` int(11) NULL DEFAULT NULL,
  `created_at` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `modified_by` int(11) NULL DEFAULT NULL,
  `modified_at` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of topic
-- ----------------------------
INSERT INTO `topic` VALUES (1, '主题名称', 1, 1, '1542680311', 1, '1543033071', NULL);
INSERT INTO `topic` VALUES (2, '主题名称', 0, 1, '1542680493', 1, '1542859370', NULL);
INSERT INTO `topic` VALUES (3, '主题名称', 0, 1, '1542681283', 0, '1542681283', NULL);
INSERT INTO `topic` VALUES (4, '主题名称', 0, 1, '1542681510', 0, '1542681510', NULL);
INSERT INTO `topic` VALUES (5, '主题名称', 1, 1, '1542681691', 0, '1542681691', NULL);
INSERT INTO `topic` VALUES (6, '1', 0, 1, '1542690742', 0, '1542690742', NULL);
INSERT INTO `topic` VALUES (7, '1', 0, 1, '1542690780', 0, '1542690780', NULL);
INSERT INTO `topic` VALUES (8, '主题名称', 0, 1, '1542699946', 0, '1542699946', NULL);
INSERT INTO `topic` VALUES (9, '主题名称', 0, 1, '1542779822', 0, '1542779822', NULL);
INSERT INTO `topic` VALUES (10, '主题名称', 0, 1, '1542780106', 0, '1542780106', NULL);
INSERT INTO `topic` VALUES (11, '主题名称', 0, 1, '1542780137', 0, '1542780137', NULL);
INSERT INTO `topic` VALUES (12, '主题名称', 0, 1, '1542780227', 0, '1542780227', NULL);
INSERT INTO `topic` VALUES (13, '主题名称aa', 0, 1, '1542780255', 0, '1542780255', NULL);
INSERT INTO `topic` VALUES (14, '主题名称', 0, 1, '1542859104', 0, '1542859104', NULL);
INSERT INTO `topic` VALUES (15, '主题名称', 0, 1, '1542859358', 0, '1542859358', NULL);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `phone` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户手机号',
  `role_id` int(11) NULL DEFAULT NULL,
  `state` int(2) NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unique_user_username`(`username`) USING BTREE,
  UNIQUE INDEX `unique_user_password`(`password`) USING BTREE,
  INDEX `fk_l_user_l_role_1`(`role_id`) USING BTREE,
  CONSTRAINT `fk_l_user_l_role_1` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户信息' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '123user', 'user', NULL, NULL, 1);

SET FOREIGN_KEY_CHECKS = 1;
