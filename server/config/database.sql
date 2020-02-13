-- MySQL dump 10.13  Distrib 8.0.19, for osx10.15 (x86_64)
--
-- Host: localhost    Database: test
-- ------------------------------------------------------
-- Server version	8.0.19

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
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
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `test_api_cates` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_test_api_cates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_api_cates`
--

LOCK TABLES `test_api_cates` WRITE;
/*!40000 ALTER TABLE `test_api_cates` DISABLE KEYS */;
INSERT INTO `test_api_cates` VALUES (1,'2020-02-12 17:26:43','2020-02-13 13:53:32',NULL,'微信配置API(test)',''),(2,'2020-02-12 17:30:52','2020-02-13 09:22:09','2020-02-13 09:22:09','后台用户管理1',''),(3,'2020-02-13 10:10:13','2020-02-13 10:10:13',NULL,'后台用户管理','asd'),(4,'2020-02-13 10:10:54','2020-02-13 13:56:31',NULL,'【后台】文章管理 😄',''),(5,'2020-02-13 10:11:11','2020-02-13 10:11:11',NULL,'da s','阿斯顿'),(6,'2020-02-13 10:12:10','2020-02-13 14:40:53',NULL,'test api','');
/*!40000 ALTER TABLE `test_api_cates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_apis`
--

DROP TABLE IF EXISTS `test_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `test_apis` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `data` longtext,
  `cid` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_test_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_apis`
--

LOCK TABLES `test_apis` WRITE;
/*!40000 ALTER TABLE `test_apis` DISABLE KEYS */;
INSERT INTO `test_apis` VALUES (1,'2020-02-12 18:13:36','2020-02-12 18:41:34','2020-02-12 18:41:34','微信','{\"body\":[{\"key\":\"阿斯顿\",\"value\":\"的\",\"type\":\"string\",\"name\":\"的\"},{\"key\":\"url\",\"value\":\"asd\",\"type\":\"string\",\"name\":\"asd \"}],\"header\":[],\"name\":\"微信\",\"url\":\"wechat/jsApiConfig\",\"method\":\"GET\"}','1'),(2,'2020-02-12 18:23:44','2020-02-12 18:41:41','2020-02-12 18:41:41','测试','{\"name\":\"测试\",\"url\":\"test\",\"method\":\"GET\"}','2'),(3,'2020-02-12 18:24:36','2020-02-13 09:09:32','2020-02-13 09:09:32','测试','{\"name\":\"测试\",\"url\":\"test\",\"method\":\"GET\"}','2'),(4,'2020-02-13 09:14:14','2020-02-13 09:14:19','2020-02-13 09:14:19','asdad','{\"name\":\"asdad\",\"url\":\"asd\",\"method\":\"GET\"}','2'),(5,'2020-02-13 09:15:43','2020-02-13 09:15:43',NULL,'asd','{\"name\":\"asd\",\"url\":\"as\",\"method\":\"GET\"}','2'),(6,'2020-02-13 09:22:23','2020-02-13 09:31:01','2020-02-13 09:31:01','asd','{\"body\":[],\"header\":[],\"name\":\"asd\",\"url\":\"asd\",\"method\":\"GET\"}','1'),(7,'2020-02-13 09:31:58','2020-02-13 09:32:02','2020-02-13 09:32:02','asd','{\"name\":\"asd\",\"url\":\"asd\",\"method\":\"POST\"}','1'),(8,'2020-02-13 09:32:37','2020-02-13 09:32:41','2020-02-13 09:32:41','asd','{\"name\":\"asd\",\"url\":\"asd\",\"method\":\"GET\"}','1'),(9,'2020-02-13 09:33:39','2020-02-13 09:33:44','2020-02-13 09:33:44','asd','{\"name\":\"asd\",\"url\":\"asd\",\"method\":\"GET\"}','1'),(10,'2020-02-13 09:39:09','2020-02-13 14:39:00',NULL,'jsSDK授权','{\"body\":[{\"key\":\"asd\",\"value\":\"asd\",\"type\":\"string\",\"name\":\"asd\"}],\"header\":[],\"name\":\"jsSDK授权\",\"url\":\"wechat/jsApiConfig\",\"method\":\"GET\",\"headers\":[{\"key\":\"sd\",\"value\":\"dsa\",\"type\":\"string\",\"name\":\"d\"}]}','1'),(11,'2020-02-13 10:12:53','2020-02-13 13:54:54',NULL,'阿斯顿','{\"body\":[{\"key\":\"asd\",\"value\":\"sd\",\"type\":\"string\",\"name\":\"d\"}],\"header\":[],\"name\":\"阿斯顿\",\"url\":\"阿斯顿\",\"method\":\"POST\"}','6'),(12,'2020-02-13 10:15:22','2020-02-13 13:49:47',NULL,'微信登陆','{\"body\":[{\"key\":\"阿斯顿\",\"value\":\"都是\",\"type\":\"string\",\"name\":\"速度\"}],\"header\":[],\"name\":\"微信登陆\",\"url\":\"wechat/wechat_redirect\",\"method\":\"GET\"}','1'),(13,'2020-02-13 13:55:10','2020-02-13 13:55:10',NULL,'ad','{\"name\":\"ad\",\"url\":\"dssdd\",\"method\":\"GET\"}','4');
/*!40000 ALTER TABLE `test_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_articles`
--

DROP TABLE IF EXISTS `test_articles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `test_articles` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
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
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `test_users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `password` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `ua` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `login_time` datetime DEFAULT NULL,
  `status` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_users`
--

LOCK TABLES `test_users` WRITE;
/*!40000 ALTER TABLE `test_users` DISABLE KEYS */;
INSERT INTO `test_users` VALUES (1,'6c3e89a67a57','suke','','[::1]:59624','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36','2020-02-11 11:45:32','2020-02-11 11:45:32',1),(2,'6c3e89a67a57','sukea','','127.0.0.1:49701','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.100 Safari/537.36','2020-02-13 11:18:15','2020-02-13 11:18:15',1),(3,'6c3e89a67a57','suk23ea','','127.0.0.1:49701','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.100 Safari/537.36','2020-02-13 11:18:21','2020-02-13 11:18:21',1),(4,'6c3e89a67a57','43223','','127.0.0.1:49701','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.100 Safari/537.36','2020-02-13 11:18:24','2020-02-13 11:18:24',1),(5,'6c3e89a67a57','asdadadfdda','','127.0.0.1:49701','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.100 Safari/537.36','2020-02-13 11:18:28','2020-02-13 11:18:28',1),(6,'6c3e89a67a57','啊的叫啊叫大会','','127.0.0.1:49701','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.100 Safari/537.36','2020-02-13 11:18:32','2020-02-13 11:18:32',1),(7,'6c3e89a67a57','啊的叫啊jd 大会','','127.0.0.1:49701','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.100 Safari/537.36','2020-02-13 11:18:35','2020-02-13 11:18:35',1),(8,'6c3e89a67a57','fasds会','','127.0.0.1:49701','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.100 Safari/537.36','2020-02-13 11:18:40','2020-02-13 11:18:40',1);
/*!40000 ALTER TABLE `test_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test_wechat_oauths`
--

DROP TABLE IF EXISTS `test_wechat_oauths`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `test_wechat_oauths` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` int DEFAULT NULL,
  `access_token` varchar(255) DEFAULT NULL,
  `expires_in` bigint DEFAULT NULL,
  `refresh_token` varchar(255) DEFAULT NULL,
  `openid` varchar(255) DEFAULT NULL,
  `scope` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test_wechat_oauths`
--

LOCK TABLES `test_wechat_oauths` WRITE;
/*!40000 ALTER TABLE `test_wechat_oauths` DISABLE KEYS */;
INSERT INTO `test_wechat_oauths` VALUES (1,0,'30_xtmEouKulDDDh1x9XAv-I7WT6iuWDiuxhQ2ATRDlXBuSZ5sBP9oiKggH5Tb44MsiJ73l_iCKA0ICu0Y92p4ifQ',1581417676,'30_8qZiQFrhL8rzTeZmOYTZIUkFrkM50JFdGXJHvPO5_wE9iq3bEZVKRS3NwmI5PFoaWcbn9pqqNvp3UKpv50MPTQ','oUsta6PmPtCCs-XSuw02Q07p1OB0','snsapi_userinfo');
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

-- Dump completed on 2020-02-13 14:41:35
