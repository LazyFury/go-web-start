-- MySQL dump 10.13  Distrib 8.0.18, for osx10.15 (x86_64)
--
-- Host: 0.0.0.0    Database: test
-- ------------------------------------------------------
-- Server version	8.0.20

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
-- Table structure for table `ek_server_api_cates`
--

DROP TABLE IF EXISTS `ek_server_api_cates`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ek_server_api_cates` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `base_url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_ek_server_api_cates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_api_cates`
--

LOCK TABLES `ek_server_api_cates` WRITE;
/*!40000 ALTER TABLE `ek_server_api_cates` DISABLE KEYS */;
INSERT INTO `ek_server_api_cates` VALUES (1,'2020-05-09 14:36:14','2020-05-12 11:40:45',NULL,'asd ','','');
/*!40000 ALTER TABLE `ek_server_api_cates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ek_server_apis`
--

DROP TABLE IF EXISTS `ek_server_apis`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ek_server_apis` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `data` longtext,
  `cid` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_ek_server_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_apis`
--

LOCK TABLES `ek_server_apis` WRITE;
/*!40000 ALTER TABLE `ek_server_apis` DISABLE KEYS */;
INSERT INTO `ek_server_apis` VALUES (1,'2020-05-09 14:37:45','2020-05-09 14:37:45',NULL,'asd','{\"name\":\"asd\",\"url\":\"asd\",\"method\":\"GET\"}','1'),(2,'2020-05-09 14:37:58','2020-05-09 14:37:58',NULL,'asd','{\"name\":\"asd\",\"url\":\"asd\",\"method\":\"GET\",\"body\":[{\"name\":\"asd \",\"key\":\"token\",\"value\":\"d\",\"type\":\"string\"}]}','1'),(3,'2020-05-12 11:40:40','2020-05-12 11:40:48',NULL,'asdasd','{\"body\":[],\"header\":[],\"name\":\"asdasd\",\"url\":\"asd\",\"method\":\"GET\"}','1');
/*!40000 ALTER TABLE `ek_server_apis` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ek_server_articles`
--

DROP TABLE IF EXISTS `ek_server_articles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ek_server_articles` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
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
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ek_server_goods` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `cid` int DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `description` mediumtext,
  `cover` varchar(255) DEFAULT NULL,
  `images` mediumtext,
  `price` double NOT NULL,
  `count` int DEFAULT NULL,
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
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ek_server_goods_cates` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `parent_id` int DEFAULT NULL,
  `cover` varchar(255) DEFAULT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `level` int DEFAULT '1',
  PRIMARY KEY (`id`),
  KEY `idx_ek_server_goods_cates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_goods_cates`
--

LOCK TABLES `ek_server_goods_cates` WRITE;
/*!40000 ALTER TABLE `ek_server_goods_cates` DISABLE KEYS */;
/*!40000 ALTER TABLE `ek_server_goods_cates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ek_server_users`
--

DROP TABLE IF EXISTS `ek_server_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ek_server_users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `password` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `ua` varchar(255) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `login_time` datetime DEFAULT NULL,
  `status` int DEFAULT NULL,
  `add_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_users`
--

LOCK TABLES `ek_server_users` WRITE;
/*!40000 ALTER TABLE `ek_server_users` DISABLE KEYS */;
INSERT INTO `ek_server_users` VALUES (1,'6c3e89a67a57','sukeai','','172.17.0.1:41684','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36','2020-05-09 14:39:24','2020-05-09 14:39:24',0,'2020-05-09 14:39:24'),(3,'6c3e89a67a57','sukeai1','','172.17.0.1:41684','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36','2020-05-09 14:40:03','2020-05-09 14:40:03',1,'2020-05-09 14:40:03'),(4,'6c3e89a67a57','阿斯顿件阿迪','','172.17.0.1:41684','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36','2020-05-09 14:40:06','2020-05-09 14:40:06',1,'2020-05-09 14:40:06'),(5,'6c3e89a67a57','阿阿斯顿件阿迪','','172.17.0.1:41684','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36','2020-05-09 14:40:09','2020-05-09 14:40:09',1,'2020-05-09 14:40:09'),(6,'6c3e89a67a57','阿阿啊的件阿迪','','172.17.0.1:41684','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36','2020-05-09 14:40:11','2020-05-09 14:40:11',0,'2020-05-09 14:40:11'),(7,'6c3e89a67a57','dgdasd','','172.17.0.1:41684','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36','2020-05-09 14:40:17','2020-05-09 14:40:17',1,'2020-05-09 14:40:17');
/*!40000 ALTER TABLE `ek_server_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ek_server_wechat_oauths`
--

DROP TABLE IF EXISTS `ek_server_wechat_oauths`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ek_server_wechat_oauths` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `uid` int DEFAULT NULL,
  `access_token` varchar(255) DEFAULT NULL,
  `expires_in` bigint DEFAULT NULL,
  `refresh_token` varchar(255) DEFAULT NULL,
  `openid` varchar(255) DEFAULT NULL,
  `scope` varchar(255) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `sex` int DEFAULT NULL,
  `headimgurl` varchar(255) DEFAULT NULL,
  `province` varchar(255) DEFAULT NULL,
  `city` varchar(255) DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  `unionid` varchar(255) DEFAULT NULL,
  `subscribe` int DEFAULT NULL,
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

-- Dump completed on 2020-05-13 11:21:49
