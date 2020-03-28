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
-- Table structure for table `test_api_cates`
--

DROP TABLE IF EXISTS `test_api_cates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `test_api_cates` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `base_url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_test_api_cates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_api_cates`
--

LOCK TABLES `test_api_cates` WRITE;
/*!40000 ALTER TABLE `test_api_cates` DISABLE KEYS */;
INSERT INTO `test_api_cates` VALUES (1,'2020-03-10 10:34:36','2020-03-26 17:31:20',NULL,'商品管理','',''),(2,'2020-03-23 09:14:32','2020-03-27 18:17:39',NULL,'用户管理','','');
/*!40000 ALTER TABLE `test_api_cates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_apis`
--

DROP TABLE IF EXISTS `test_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `test_apis` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `data` longtext,
  `cid` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_test_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_apis`
--

LOCK TABLES `test_apis` WRITE;
/*!40000 ALTER TABLE `test_apis` DISABLE KEYS */;
INSERT INTO `test_apis` VALUES (1,'2020-03-10 10:34:48','2020-03-10 10:42:41','2020-03-10 10:42:41','asd','{\"name\":\"asd\",\"url\":\"/hello\",\"method\":\"GET\"}','1'),(2,'2020-03-10 10:35:06','2020-03-28 11:13:14',NULL,'添加商品','{\"body\":[{\"name\":\"title\",\"key\":\"title\",\"value\":\"商品标题\",\"type\":\"number\"},{\"name\":\"desc\",\"key\":\"desc\",\"value\":\"商品简介\",\"type\":\"string\"},{\"name\":\"cover\",\"key\":\"cover\",\"value\":\"http://asdasd.d.dd.jmp\",\"type\":\"string\"},{\"name\":\"price\",\"key\":\"price\",\"value\":3.012,\"type\":\"number\"},{\"name\":\"count\",\"key\":\"count\",\"value\":90,\"type\":\"number\"},{\"name\":\"token\",\"key\":\"token\",\"value\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODUzNzAwMDUsImlhdCI6MTU4NTI4MzYwNSwiaWQiOjE0LCJuYmYiOjE1ODUyODM2MDUsInVzZXJuYW1lIjoic3VrZWFpIn0.vSkrbLLj4jOrl9d-4n8rI2HRQTMEI9egMXITj0d-3yE\",\"type\":\"string\"},{\"name\":\"cid\",\"key\":\"cid\",\"value\":29,\"type\":\"number\"}],\"header\":[],\"name\":\"添加商品\",\"url\":\"/admin/product/add\",\"method\":\"POST\",\"headers\":[]}','1'),(3,'2020-03-10 10:53:26','2020-03-28 09:30:44',NULL,'商品列表','{\"body\":[{\"name\":\"page\",\"key\":\"page\",\"value\":1,\"type\":\"number\"},{\"name\":\"limit\",\"key\":\"limit\",\"value\":\"10\",\"type\":\"string\"},{\"name\":\"order\",\"key\":\"order\",\"value\":\"id_desc\",\"type\":\"string\"},{\"name\":\"key\",\"key\":\"key\",\"value\":\"比阿嚏\",\"type\":\"string\"},{\"name\":\"cid\",\"key\":\"cid\",\"value\":\"\",\"type\":\"string\"}],\"header\":[],\"name\":\"商品列表\",\"url\":\"/product/list\",\"method\":\"GET\"}','1'),(4,'2020-03-23 09:14:53','2020-03-26 15:10:23','2020-03-26 15:10:23','注册','{\"body\":[{\"name\":\"user\",\"key\":\"login\",\"value\":\"17638163264\",\"type\":\"string\"},{\"name\":\"pwd\",\"key\":\"password\",\"value\":\"123456\",\"type\":\"string\"},{\"name\":\"validCode\",\"key\":\"validCode\",\"value\":\"351125\",\"type\":\"string\"},{\"name\":\"inviteCode\",\"key\":\"inviteCode\",\"value\":\"\",\"type\":\"string\"},{\"name\":\"imageCode\",\"key\":\"imageCode\",\"value\":\"\",\"type\":\"string\"},{\"name\":\"type\",\"key\":\"type\",\"value\":1,\"type\":\"number\"},{\"name\":\"channel\",\"key\":\"channel\",\"value\":1,\"type\":\"number\"}],\"header\":[],\"name\":\"注册\",\"url\":\"/handleRegister\",\"method\":\"POST\",\"headers\":[{\"name\":\"Content-Type\",\"key\":\"Content-Type\",\"value\":\"application/x-www-form-urlencoded\",\"type\":\"string\"},{\"name\":\"Accept\",\"key\":\"Accept\",\"value\":\"application/json\",\"type\":\"string\"}]}','2'),(5,'2020-03-23 09:27:29','2020-03-26 15:10:35','2020-03-26 15:10:35','发送注册图片验证码','{\"name\":\"发送注册图片验证码\",\"url\":\"/base/getImageCode\",\"method\":\"GET\",\"body\":[{\"name\":\"login\",\"key\":\"login\",\"value\":\"17638163264\",\"type\":\"string\"}]}','2'),(6,'2020-03-23 09:34:52','2020-03-26 15:10:32','2020-03-26 15:10:32','发送注册验证码','{\"body\":[{\"name\":\"login\",\"key\":\"login\",\"value\":\"17638163264\",\"type\":\"string\"}],\"header\":[],\"name\":\"发送注册验证码\",\"url\":\"/sendRegisterCode\",\"method\":\"GET\"}','2'),(7,'2020-03-23 10:01:32','2020-03-26 15:10:38','2020-03-26 15:10:38','登陆','{\"body\":[{\"name\":\"login\",\"key\":\"login\",\"value\":\"17638163264\",\"type\":\"string\"},{\"name\":\"pwd\",\"key\":\"password\",\"value\":\"123456\",\"type\":\"string\"},{\"name\":\"type\",\"key\":\"type\",\"value\":1,\"type\":\"number\"},{\"name\":\"channel\",\"key\":\"channel\",\"value\":\"1\",\"type\":\"string\"},{\"name\":\"validCode\",\"key\":\"validCode\",\"value\":\"089047\",\"type\":\"string\"}],\"header\":[],\"name\":\"登陆\",\"url\":\"/login\",\"method\":\"GET\"}','2'),(8,'2020-03-23 10:04:25','2020-03-26 15:13:38',NULL,'登陆','{\"body\":[{\"name\":\"username\",\"key\":\"username\",\"value\":\"sukeai\",\"type\":\"string\"},{\"name\":\"password\",\"key\":\"password\",\"value\":\"123456\",\"type\":\"string\"}],\"header\":[],\"name\":\"登陆\",\"url\":\"/admin/login\",\"method\":\"GET\"}','2'),(9,'2020-03-26 15:12:29','2020-03-26 15:13:31',NULL,'注册','{\"body\":[{\"name\":\"username\",\"key\":\"username\",\"value\":\"sukeai\",\"type\":\"string\"},{\"name\":\"password\",\"key\":\"password\",\"value\":\"123456\",\"type\":\"string\"}],\"header\":[],\"name\":\"注册\",\"url\":\"/admin/login/reg\",\"method\":\"GET\"}','2'),(10,'2020-03-27 10:17:16','2020-03-27 18:14:41',NULL,'用户列表','{\"body\":[],\"header\":[],\"name\":\"用户列表\",\"url\":\"/admin/user/list\",\"method\":\"GET\",\"headers\":[{\"name\":\"token\",\"key\":\"token\",\"value\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODUzNzAwMDUsImlhdCI6MTU4NTI4MzYwNSwiaWQiOjE0LCJuYmYiOjE1ODUyODM2MDUsInVzZXJuYW1lIjoic3VrZWFpIn0.vSkrbLLj4jOrl9d-4n8rI2HRQTMEI9egMXITj0d-3yE\",\"type\":\"string\"}]}','2'),(11,'2020-03-27 15:20:04','2020-03-28 11:30:56',NULL,'添加分类','{\"body\":[{\"name\":\"name\",\"key\":\"name\",\"value\":\"普通商品\",\"type\":\"string\"},{\"name\":\"desc\",\"key\":\"desc\",\"value\":\"分类简介阿萨德的.as.\",\"type\":\"string\"},{\"name\":\"cover\",\"key\":\"cover\",\"value\":\"http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg\",\"type\":\"string\"},{\"name\":\"parent_id\",\"key\":\"parent_id\",\"value\":30,\"type\":\"number\"}],\"header\":[],\"name\":\"添加分类\",\"url\":\"/admin/product/cate/add\",\"method\":\"POST\",\"headers\":[{\"name\":\"token\",\"key\":\"token\",\"value\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODUzNzAwMDUsImlhdCI6MTU4NTI4MzYwNSwiaWQiOjE0LCJuYmYiOjE1ODUyODM2MDUsInVzZXJuYW1lIjoic3VrZWFpIn0.vSkrbLLj4jOrl9d-4n8rI2HRQTMEI9egMXITj0d-3yE\",\"type\":\"string\"}]}','1'),(12,'2020-03-27 15:56:53','2020-03-27 15:57:25',NULL,'分类','{\"body\":[],\"header\":[],\"name\":\"分类\",\"url\":\"/admin/product/cate/list\",\"method\":\"GET\",\"headers\":[{\"name\":\"token\",\"key\":\"token\",\"value\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODUzNzAwMDUsImlhdCI6MTU4NTI4MzYwNSwiaWQiOjE0LCJuYmYiOjE1ODUyODM2MDUsInVzZXJuYW1lIjoic3VrZWFpIn0.vSkrbLLj4jOrl9d-4n8rI2HRQTMEI9egMXITj0d-3yE\",\"type\":\"string\"}]}','1');
/*!40000 ALTER TABLE `test_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_articles`
--

DROP TABLE IF EXISTS `test_articles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `test_articles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_test_articles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_articles`
--

LOCK TABLES `test_articles` WRITE;
/*!40000 ALTER TABLE `test_articles` DISABLE KEYS */;
/*!40000 ALTER TABLE `test_articles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_goods`
--

DROP TABLE IF EXISTS `test_goods`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `test_goods` (
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
  KEY `idx_test_goods_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_goods`
--

LOCK TABLES `test_goods` WRITE;
/*!40000 ALTER TABLE `test_goods` DISABLE KEYS */;
INSERT INTO `test_goods` VALUES (4,'2020-03-28 10:23:51','2020-03-28 10:23:51',NULL,29,'商品标题','','http://asdasd.d.dd.jmp',NULL,3.01,90),(5,'2020-03-28 11:13:18','2020-03-28 11:13:18',NULL,29,'商品标题','','http://asdasd.d.dd.jmp',NULL,3.01,90);
/*!40000 ALTER TABLE `test_goods` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_goods_cates`
--

DROP TABLE IF EXISTS `test_goods_cates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `test_goods_cates` (
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
  KEY `idx_test_goods_cates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=56 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_goods_cates`
--

LOCK TABLES `test_goods_cates` WRITE;
/*!40000 ALTER TABLE `test_goods_cates` DISABLE KEYS */;
INSERT INTO `test_goods_cates` VALUES (28,'2020-03-28 10:09:13','2020-03-28 10:09:13',NULL,'普通商品','分类简介阿萨德的.as.',0,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',1),(29,'2020-03-28 10:09:17','2020-03-28 10:09:17',NULL,'普通商品','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(30,'2020-03-28 10:09:22','2020-03-28 10:09:22',NULL,'普通商品','分类简介阿萨德的.as.',29,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',3),(31,'2020-03-28 10:09:26','2020-03-28 10:09:26',NULL,'普通商品','分类简介阿萨德的.as.',30,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',4),(32,'2020-03-28 10:10:20','2020-03-28 10:10:20',NULL,'普通商品','分类简介阿萨德的.as.',29,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',3),(33,'2020-03-28 10:11:16','2020-03-28 10:11:16',NULL,'普通商品','分类简介阿萨德的.as.',30,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',4),(34,'2020-03-28 10:14:11','2020-03-28 10:14:11',NULL,'普通商品12','分类简介阿萨德的.as.',29,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',3),(35,'2020-03-28 10:14:12','2020-03-28 10:14:12',NULL,'普通商品12','分类简介阿萨德的.as.',29,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',3),(36,'2020-03-28 10:14:14','2020-03-28 10:14:14',NULL,'普通商品12','分类简介阿萨德的.as.',29,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',3),(37,'2020-03-28 10:19:56','2020-03-28 10:19:56',NULL,'普通商品12asdada','分类简介阿萨德的.as.',29,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',3),(38,'2020-03-28 10:21:55','2020-03-28 10:21:55',NULL,'普通商品7','分类简介阿萨德的.as.',30,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',4),(39,'2020-03-28 10:43:49','2020-03-28 10:43:49',NULL,'普通商品1','分类简介阿萨德的.as.',0,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',1),(40,'2020-03-28 10:43:53','2020-03-28 10:43:53',NULL,'普通商品2','分类简介阿萨德的.as.',0,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',1),(41,'2020-03-28 10:43:55','2020-03-28 10:43:55',NULL,'普通商品3','分类简介阿萨德的.as.',0,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',1),(42,'2020-03-28 11:01:18','2020-03-28 11:01:18',NULL,'2普通商品','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(43,'2020-03-28 11:01:21','2020-03-28 11:01:21',NULL,'2普通3商品','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(44,'2020-03-28 11:01:22','2020-03-28 11:01:22',NULL,'2普通3商品4','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(45,'2020-03-28 11:01:24','2020-03-28 11:01:24',NULL,'2普通3商品42','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(46,'2020-03-28 11:01:27','2020-03-28 11:01:27',NULL,'2普通3商品142','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(47,'2020-03-28 11:01:28','2020-03-28 11:01:28',NULL,'2普通34商品142','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(48,'2020-03-28 11:01:30','2020-03-28 11:01:30',NULL,'2普通34商4品142','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(49,'2020-03-28 11:01:33','2020-03-28 11:01:33',NULL,'21普通34商4品142','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(50,'2020-03-28 11:01:35','2020-03-28 11:01:35',NULL,'44商4品142','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(51,'2020-03-28 11:01:37','2020-03-28 11:01:37',NULL,'44商4品1423','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(52,'2020-03-28 11:01:39','2020-03-28 11:01:39',NULL,'44商4品14232','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(53,'2020-03-28 11:01:41','2020-03-28 11:01:41',NULL,'44商4品142324','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(54,'2020-03-28 11:01:42','2020-03-28 11:01:42',NULL,'44商4品1422324','分类简介阿萨德的.as.',28,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',2),(55,'2020-03-28 11:31:15','2020-03-28 11:31:15',NULL,'普通商;品1','分类简介阿萨德的.as.',30,'http://wx4.sinaimg.cn/mw600/0085KTY1gy1gd88xot5qdj30bd0dkgrc.jpg','',4);
/*!40000 ALTER TABLE `test_goods_cates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_users`
--

DROP TABLE IF EXISTS `test_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `test_users` (
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
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_users`
--

LOCK TABLES `test_users` WRITE;
/*!40000 ALTER TABLE `test_users` DISABLE KEYS */;
INSERT INTO `test_users` VALUES (1,'3c7fde','suke','','127.0.0.1:61261','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36','2020-03-11 09:36:30','2020-03-11 09:36:30',1,NULL),(3,'3c7fde','suke1','','127.0.0.1:61807','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36','2020-03-11 09:42:57','2020-03-11 09:42:57',1,NULL),(4,'3c7fde','suke2','','127.0.0.1:61807','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36','2020-03-11 09:43:00','2020-03-11 09:43:00',1,NULL),(5,'3c7fde','suke3','','127.0.0.1:61807','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36','2020-03-11 09:43:03','2020-03-11 09:43:03',1,NULL),(6,'3c7fde','suke4','','127.0.0.1:61807','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36','2020-03-11 09:43:06','2020-03-11 09:43:06',1,NULL),(7,'3c7fde','suke5','','127.0.0.1:61807','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36','2020-03-11 09:43:09','2020-03-11 09:43:09',1,NULL),(8,'3c7fde','suke6','','127.0.0.1:61807','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36','2020-03-11 09:44:15','2020-03-11 09:44:15',1,NULL),(9,'3c7fde','suke7','','127.0.0.1:61807','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36','2020-03-11 09:44:19','2020-03-11 09:44:19',1,NULL),(10,'3c7fde','suke71','','127.0.0.1:61807','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36','2020-03-11 09:44:25','2020-03-11 09:44:25',1,NULL),(11,'3c7fde','suke12','','127.0.0.1:61807','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36','2020-03-11 09:44:29','2020-03-11 09:44:29',1,NULL),(12,'3c7fde','suke13','','127.0.0.1:61807','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36','2020-03-11 09:44:33','2020-03-11 09:44:33',1,NULL),(14,'6c3e89a67a57','sukeai','','127.0.0.1:53124','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36','2020-03-26 15:13:28','2020-03-26 15:13:28',1,'2020-03-26 15:13:28'),(15,'6c3e89a67a57','sukeai1','','127.0.0.1:50923','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36','2020-03-27 10:18:48','2020-03-27 10:18:48',1,'2020-03-27 10:18:48');
/*!40000 ALTER TABLE `test_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_wechat_oauths`
--

DROP TABLE IF EXISTS `test_wechat_oauths`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `test_wechat_oauths` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) DEFAULT NULL,
  `access_token` varchar(255) DEFAULT NULL,
  `expires_in` bigint(20) DEFAULT NULL,
  `refresh_token` varchar(255) DEFAULT NULL,
  `openid` varchar(255) DEFAULT NULL,
  `scope` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `sex` int(11) DEFAULT NULL,
  `headimgurl` varchar(255) DEFAULT NULL,
  `province` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  `unionid` varchar(255) DEFAULT NULL,
  `subscribe` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `idx_test_wechat_oauths_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_wechat_oauths`
--

LOCK TABLES `test_wechat_oauths` WRITE;
/*!40000 ALTER TABLE `test_wechat_oauths` DISABLE KEYS */;
/*!40000 ALTER TABLE `test_wechat_oauths` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-03-28 13:39:24
