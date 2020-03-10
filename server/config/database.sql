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
  PRIMARY KEY (`id`),
  KEY `idx_test_api_cates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_api_cates`
--

LOCK TABLES `test_api_cates` WRITE;
/*!40000 ALTER TABLE `test_api_cates` DISABLE KEYS */;
INSERT INTO `test_api_cates` VALUES (1,'2020-03-10 10:34:36','2020-03-10 10:42:55',NULL,'基础接口','');
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_apis`
--

LOCK TABLES `test_apis` WRITE;
/*!40000 ALTER TABLE `test_apis` DISABLE KEYS */;
INSERT INTO `test_apis` VALUES (1,'2020-03-10 10:34:48','2020-03-10 10:42:41','2020-03-10 10:42:41','asd','{\"name\":\"asd\",\"url\":\"/hello\",\"method\":\"GET\"}','1'),(2,'2020-03-10 10:35:06','2020-03-10 10:44:06',NULL,'首页','{\"body\":[{\"name\":\"test\",\"key\":\"key\",\"value\":\"value\",\"type\":\"string\",\"tips\":\"添加参数测试\"}],\"header\":[],\"name\":\"首页\",\"url\":\"/\",\"method\":\"GET\",\"headers\":[]}','1'),(3,'2020-03-10 10:53:26','2020-03-10 10:53:44',NULL,'wechat','{\"body\":[{\"name\":\"url\",\"key\":\"url\",\"value\":\"baidu.com\",\"type\":\"string\"}],\"header\":[],\"name\":\"wechat\",\"url\":\"/wechat/jsApiConfig\",\"method\":\"GET\"}','1');
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
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_users`
--

LOCK TABLES `test_users` WRITE;
/*!40000 ALTER TABLE `test_users` DISABLE KEYS */;
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
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
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

-- Dump completed on 2020-03-10 11:05:02
