-- MySQL dump 10.13  Distrib 8.0.11, for macos10.13 (x86_64)
--
-- Host: localhost    Database: test
-- ------------------------------------------------------
-- Server version	8.0.11

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `ek_server_api_cates`
--

DROP TABLE IF EXISTS `ek_server_api_cates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `ek_server_api_cates` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `base_url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_ek_server_api_cates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_api_cates`
--

LOCK TABLES `ek_server_api_cates` WRITE;
/*!40000 ALTER TABLE `ek_server_api_cates` DISABLE KEYS */;
INSERT INTO `ek_server_api_cates` VALUES (1,'2020-03-28 14:17:21','2020-03-31 17:54:50',NULL,'商品管理','',''),(2,'2020-03-28 14:17:31','2020-03-28 14:17:31',NULL,'用户管理','用户管理',''),(3,'2020-03-28 14:17:40','2020-03-28 14:17:40',NULL,'微信接口','微信接口','');
/*!40000 ALTER TABLE `ek_server_api_cates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ek_server_apis`
--

DROP TABLE IF EXISTS `ek_server_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `ek_server_apis` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `data` longtext,
  `cid` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_ek_server_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_apis`
--

LOCK TABLES `ek_server_apis` WRITE;
/*!40000 ALTER TABLE `ek_server_apis` DISABLE KEYS */;
INSERT INTO `ek_server_apis` VALUES (1,'2020-03-28 14:18:08','2020-03-31 17:54:28',NULL,'添加商品','{\"body\":[],\"header\":[],\"name\":\"添加商品\",\"url\":\"/admin/product/add\",\"method\":\"POST\",\"headers\":[{\"name\":\"token\",\"key\":\"token\",\"value\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU3MzQ4NDMsImlhdCI6MTU4NTY0ODQ0MywiaWQiOjEsIm5iZiI6MTU4NTY0ODQ0MywidXNlcm5hbWUiOiJzdWtlIn0.Z1TspgrcG92qpkg6MDmliKmHSr_EvCf-ugY3ijNkUfU\",\"type\":\"string\"}]}','1'),(2,'2020-03-28 14:52:33','2020-03-28 14:53:05',NULL,'用户登陆','{\"body\":[{\"name\":\"username\",\"key\":\"username\",\"value\":\"suke\",\"type\":\"string\"},{\"name\":\"password\",\"key\":\"password\",\"value\":\"123456\",\"type\":\"string\"}],\"header\":[],\"name\":\"用户登陆\",\"url\":\"/admin/login\",\"method\":\"GET\"}','2'),(3,'2020-03-28 14:53:42','2020-03-28 14:53:42',NULL,'用户注册','{\"name\":\"用户注册\",\"url\":\"/admin/login/reg\",\"method\":\"GET\",\"body\":[{\"name\":\"username\",\"key\":\"username\",\"value\":\"suke\",\"type\":\"string\"},{\"name\":\"password\",\"key\":\"password\",\"value\":\"123456\",\"type\":\"string\"}]}','2'),(4,'2020-03-28 14:54:07','2020-03-31 17:54:10',NULL,'用户列表','{\"body\":[],\"header\":[],\"name\":\"用户列表\",\"url\":\"/admin/user/list\",\"method\":\"GET\",\"headers\":[{\"name\":\"token\",\"key\":\"token\",\"value\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU3MzQ4NDMsImlhdCI6MTU4NTY0ODQ0MywiaWQiOjEsIm5iZiI6MTU4NTY0ODQ0MywidXNlcm5hbWUiOiJzdWtlIn0.Z1TspgrcG92qpkg6MDmliKmHSr_EvCf-ugY3ijNkUfU\",\"type\":\"string\"}]}','2'),(5,'2020-03-28 14:58:11','2020-03-28 14:58:11',NULL,'微信配置','{\"name\":\"微信配置\",\"url\":\"/wechat/jsApiConfig\",\"method\":\"GET\",\"body\":[{\"name\":\"url\",\"key\":\"url\",\"value\":\"123\",\"type\":\"string\"}]}','3'),(6,'2020-03-31 17:55:48','2020-03-31 17:58:08',NULL,'添加分类','{\"body\":[{\"name\":\"parent_id\",\"key\":\"parent_id\",\"value\":1,\"type\":\"number\"},{\"name\":\"name\",\"key\":\"name\",\"value\":\"asd3\",\"type\":\"string\"}],\"header\":[],\"name\":\"添加分类\",\"url\":\"/admin/product/cate/add\",\"method\":\"POST\",\"headers\":[{\"name\":\"token\",\"key\":\"token\",\"value\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU3MzQ4NDMsImlhdCI6MTU4NTY0ODQ0MywiaWQiOjEsIm5iZiI6MTU4NTY0ODQ0MywidXNlcm5hbWUiOiJzdWtlIn0.Z1TspgrcG92qpkg6MDmliKmHSr_EvCf-ugY3ijNkUfU\",\"type\":\"string\"}]}','1'),(7,'2020-03-31 17:56:47','2020-03-31 17:56:47',NULL,'分类列表','{\"name\":\"分类列表\",\"url\":\"/admin/product/cate/list\",\"method\":\"GET\",\"headers\":[{\"name\":\"token\",\"key\":\"token\",\"value\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODU3MzQ4NDMsImlhdCI6MTU4NTY0ODQ0MywiaWQiOjEsIm5iZiI6MTU4NTY0ODQ0MywidXNlcm5hbWUiOiJzdWtlIn0.Z1TspgrcG92qpkg6MDmliKmHSr_EvCf-ugY3ijNkUfU\",\"type\":\"string\"}]}','1');
/*!40000 ALTER TABLE `ek_server_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ek_server_articles`
--

DROP TABLE IF EXISTS `ek_server_articles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `ek_server_articles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_ek_server_articles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_articles`
--

LOCK TABLES `ek_server_articles` WRITE;
/*!40000 ALTER TABLE `ek_server_articles` DISABLE KEYS */;
/*!40000 ALTER TABLE `ek_server_articles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ek_server_goods`
--

DROP TABLE IF EXISTS `ek_server_goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `ek_server_goods` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `cid` int(11) DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` mediumtext,
  `cover` varchar(255) DEFAULT NULL,
  `images` mediumtext,
  `price` double NOT NULL,
  `count` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_ek_server_goods_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_goods`
--

LOCK TABLES `ek_server_goods` WRITE;
/*!40000 ALTER TABLE `ek_server_goods` DISABLE KEYS */;
/*!40000 ALTER TABLE `ek_server_goods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ek_server_goods_cates`
--

DROP TABLE IF EXISTS `ek_server_goods_cates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `ek_server_goods_cates` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `parent_id` int(11) DEFAULT NULL,
  `cover` varchar(255) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `level` int(11) DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `idx_ek_server_goods_cates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_goods_cates`
--

LOCK TABLES `ek_server_goods_cates` WRITE;
/*!40000 ALTER TABLE `ek_server_goods_cates` DISABLE KEYS */;
INSERT INTO `ek_server_goods_cates` VALUES (1,'2020-03-31 17:56:22','2020-03-31 17:56:22',NULL,'','',0,'','',1),(2,'2020-03-31 17:57:39','2020-03-31 17:57:39',NULL,'asd','',0,'','',1),(3,'2020-03-31 17:58:05','2020-03-31 17:58:05',NULL,'asd3','',1,'','',2),(4,'2020-03-31 17:59:37','2020-03-31 17:59:37',NULL,' ','',1,'','',2),(5,'2020-03-31 18:02:34','2020-03-31 18:02:34',NULL,'-','',1,'','',2),(6,'2020-03-31 18:02:41','2020-03-31 18:02:41',NULL,'-*&/}}','',1,'','',2),(7,'2020-03-31 18:31:32','2020-03-31 18:31:32',NULL,'asd23','',1,'','',2);
/*!40000 ALTER TABLE `ek_server_goods_cates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ek_server_users`
--

DROP TABLE IF EXISTS `ek_server_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `ek_server_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `password` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `ua` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `login_time` datetime DEFAULT NULL,
  `status` int(11) DEFAULT NULL,
  `add_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_users`
--

LOCK TABLES `ek_server_users` WRITE;
/*!40000 ALTER TABLE `ek_server_users` DISABLE KEYS */;
INSERT INTO `ek_server_users` VALUES (1,'6c3e89a67a57','suke','','127.0.0.1:59350','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36','2020-03-28 14:53:45','2020-03-28 14:53:45',1,'2020-03-28 14:53:45');
/*!40000 ALTER TABLE `ek_server_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ek_server_wechat_oauths`
--

DROP TABLE IF EXISTS `ek_server_wechat_oauths`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `ek_server_wechat_oauths` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `uid` int(11) DEFAULT NULL,
  `access_token` varchar(255) DEFAULT NULL,
  `expires_in` bigint(20) DEFAULT NULL,
  `refresh_token` varchar(255) DEFAULT NULL,
  `openid` varchar(255) DEFAULT NULL,
  `scope` varchar(255) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `sex` int(11) DEFAULT NULL,
  `headimgurl` varchar(255) DEFAULT NULL,
  `province` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  `unionid` varchar(255) DEFAULT NULL,
  `subscribe` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_ek_server_wechat_oauths_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_wechat_oauths`
--

LOCK TABLES `ek_server_wechat_oauths` WRITE;
/*!40000 ALTER TABLE `ek_server_wechat_oauths` DISABLE KEYS */;
/*!40000 ALTER TABLE `ek_server_wechat_oauths` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-04-08 10:02:48
