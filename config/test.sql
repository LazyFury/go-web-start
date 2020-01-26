/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80019
 Source Host           : localhost:3306
 Source Schema         : test

 Target Server Type    : MySQL
 Target Server Version : 80019
 File Encoding         : 65001

 Date: 25/01/2020 22:00:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `ua` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `login_time` datetime DEFAULT NULL,
  `status` int NOT NULL,
  `Email` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `id` (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=252 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (208, 'sukdajsd', '', '', '', NULL, NULL, 0, '2568597007@qq.com');
INSERT INTO `users` VALUES (209, '喜欢邻家小姐姐', 'sukeaiz', '127.0.0.1:59555', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36', '2020-01-16 11:02:19', '2020-01-16 11:02:19', 1, '2568597007@qq.com');
INSERT INTO `users` VALUES (216, '一秒钟见不到', 'sukeaiz', '127.0.0.1:56246', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36', '2020-01-16 11:30:48', '2020-01-16 11:30:48', 0, '2568597007@qq.com');
INSERT INTO `users` VALUES (220, '就难过的不行那种', 'sukeaiz', '127.0.0.1:57689', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36', '2020-01-16 12:35:04', '2020-01-16 12:35:04', 1, '2568597007@qq.com');
INSERT INTO `users` VALUES (221, 'sukeai78', 'sukeaiz', '127.0.0.1:57689', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36', '2020-01-16 12:35:09', '2020-01-16 12:35:09', 1, NULL);
INSERT INTO `users` VALUES (222, 'sukeai7h', 'sukeaiz', '127.0.0.1:57689', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36', '2020-01-16 12:35:13', '2020-01-16 12:35:13', 1, NULL);
INSERT INTO `users` VALUES (226, '123123', 'sukeaiz', '127.0.0.1:59657', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36', '2020-01-16 13:44:34', '2020-01-16 13:44:34', 1, NULL);
INSERT INTO `users` VALUES (229, 'sukeai', 'sukeaiz', '127.0.0.1:55742', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36', '2020-01-16 13:54:43', '2020-01-16 13:54:43', 1, '');
INSERT INTO `users` VALUES (239, '苏可爱', 'sukeaiz', '125.93.252.188', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36', '2020-01-16 14:48:53', '2020-01-16 14:48:53', 1, '');
INSERT INTO `users` VALUES (240, '苏可爱?', 'sukeaiz', '125.93.252.188', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.117 Safari/537.36', '2020-01-16 14:48:57', '2020-01-16 14:48:57', 0, '');
INSERT INTO `users` VALUES (241, 'suke', '6c3e89a67a57', '127.0.0.1:54974', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36', '2020-01-24 20:23:37', '2020-01-24 20:23:37', 1, '');
INSERT INTO `users` VALUES (244, 'sukeai2', '2e79d1f72e08e8', '127.0.0.1:58641', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36', '2020-01-24 20:55:23', '2020-01-24 20:55:23', 1, '');
INSERT INTO `users` VALUES (245, 'sukeai21', '2e79d1f72e08e8', '127.0.0.1:58641', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36', '2020-01-24 20:55:25', '2020-01-24 20:55:25', 1, '');
INSERT INTO `users` VALUES (246, 'sukeai21123', '2e79d1f72e08e8', '127.0.0.1:58641', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36', '2020-01-24 20:55:27', '2020-01-24 20:55:27', 1, '');
INSERT INTO `users` VALUES (247, 'sukeai2f1123', '2e79d1f72e08e8', '127.0.0.1:58641', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36', '2020-01-24 20:55:30', '2020-01-24 20:55:30', 1, '');
INSERT INTO `users` VALUES (248, 'sukedai2f1123', '2e79d1f72e08e8', '127.0.0.1:58641', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36', '2020-01-24 20:55:32', '2020-01-24 20:55:32', 1, '');
INSERT INTO `users` VALUES (249, 'sukedagi2f1123', '2e79d1f72e08e8', '127.0.0.1:58641', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36', '2020-01-24 20:55:34', '2020-01-24 20:55:34', 1, '');
INSERT INTO `users` VALUES (250, 'suksedagi2f1123', '2e79d1f72e08e8', '127.0.0.1:58641', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36', '2020-01-24 20:55:36', '2020-01-24 20:55:36', 1, '');
INSERT INTO `users` VALUES (251, 'suksesdagi2f1123', '2e79d1f72e08e8', '127.0.0.1:58641', 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36', '2020-01-24 20:55:39', '2020-01-24 20:55:39', 1, '');
COMMIT;

-- ----------------------------
-- Table structure for wechat_oauths
-- ----------------------------
DROP TABLE IF EXISTS `wechat_oauths`;
CREATE TABLE `wechat_oauths` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int DEFAULT NULL,
  `access_token` varchar(255) DEFAULT NULL,
  `expires_in` varchar(255) DEFAULT NULL,
  `refresh_token` varchar(255) DEFAULT NULL,
  `openid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `scope` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of wechat_oauths
-- ----------------------------
BEGIN;
INSERT INTO `wechat_oauths` VALUES (1, 0, '29_3i_ey4fVg2WAu_de8_U4F9cB5eCj1HdCIwFnbdRW-zr03TbpDRKw1GGTFkHLXl6S5vtfRqbZDgsZxJJjQ1uQRg', '1579877572', '29_D-kdDsZ_PghqDl6SDMwlnuEzPTjdJUDW_5zG-d05-vrHjmsSKqfV0XExCcHQnKtMntRRu02yOfh6P4sEKlDM9A', 'oUsta6PmPtCCs-XSuw02Q07p1OB0', 'snsapi_userinfo');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
