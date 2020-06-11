Terminal close -- exit!
trib 8.0.18, for osx10.15 (x86_64)
--
-- Host: 127.0.0.1    Database: test
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_api_cates`
--

LOCK TABLES `ek_server_api_cates` WRITE;
/*!40000 ALTER TABLE `ek_server_api_cates` DISABLE KEYS */;
INSERT INTO `ek_server_api_cates` VALUES (1,'2020-05-29 14:49:07','2020-06-05 10:06:32',NULL,'用户','',''),(2,'2020-06-05 10:08:13','2020-06-05 10:08:13',NULL,'商品','商品','');
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
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_apis`
--

LOCK TABLES `ek_server_apis` WRITE;
/*!40000 ALTER TABLE `ek_server_apis` DISABLE KEYS */;
INSERT INTO `ek_server_apis` VALUES (1,'2020-05-29 14:49:19','2020-06-05 10:06:23',NULL,'登陆','{\"body\":[{\"name\":\"username\",\"key\":\"username\",\"value\":\"suke\",\"type\":\"string\"},{\"name\":\"password\",\"key\":\"password\",\"value\":\"123456\",\"type\":\"string\"}],\"header\":[],\"name\":\"登陆\",\"url\":\"/admin/login\",\"method\":\"GET\"}','1'),(2,'2020-05-29 14:49:29','2020-05-29 14:55:40',NULL,'注册','{\"body\":[{\"name\":\"password\",\"key\":\"password\",\"value\":\"123456\",\"type\":\"string\"},{\"name\":\"name\",\"key\":\"name\",\"value\":\"suke\",\"type\":\"string\"}],\"header\":[],\"name\":\"注册\",\"url\":\"/admin/login/reg\",\"method\":\"POST\"}','1');
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_goods_cates`
--

LOCK TABLES `ek_server_goods_cates` WRITE;
/*!40000 ALTER TABLE `ek_server_goods_cates` DISABLE KEYS */;
/*!40000 ALTER TABLE `ek_server_goods_cates` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ek_server_posts`
--

DROP TABLE IF EXISTS `ek_server_posts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ek_server_posts` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `title` varchar(255) DEFAULT NULL,
  `desc` varchar(255) DEFAULT NULL,
  `author` varchar(255) DEFAULT NULL,
  `context` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_ek_server_posts_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_posts`
--

LOCK TABLES `ek_server_posts` WRITE;
/*!40000 ALTER TABLE `ek_server_posts` DISABLE KEYS */;
/*!40000 ALTER TABLE `ek_server_posts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `ek_server_users`
--

DROP TABLE IF EXISTS `ek_server_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `ek_server_users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `ua` varchar(255) DEFAULT NULL,
  `login_time` datetime DEFAULT NULL,
  `status` int DEFAULT NULL,
  `is_admin` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_ek_server_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_users`
--

LOCK TABLES `ek_server_users` WRITE;
/*!40000 ALTER TABLE `ek_server_users` DISABLE KEYS */;
INSERT INTO `ek_server_users` VALUES (1,'2020-05-29 14:49:58','2020-05-29 14:49:58',NULL,'6c3e89a67a57','','','172.21.0.1:38328','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36','2020-05-29 14:49:58',1,0),(5,'2020-05-29 14:55:41','2020-05-29 14:55:41',NULL,'6c3e89a67a57','suke','','172.21.0.1:38360','Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36','2020-05-29 14:55:41',1,0),(7,'2020-06-08 17:41:18','2020-06-08 17:41:18',NULL,'6c3e89','sue','','127.0.0.1:63716','PostmanRuntime/7.24.1','2020-06-08 17:41:18',1,0),(17,'2020-06-08 17:42:18','2020-06-08 17:42:18',NULL,'3c7fde','asdwerrtre','','127.0.0.1:63716','PostmanRuntime/7.24.1','2020-06-08 17:42:18',1,0),(19,'2020-06-08 17:42:23','2020-06-08 17:42:23',NULL,'3c7fde','asdwerrtre1','','127.0.0.1:63716','PostmanRuntime/7.24.1','2020-06-08 17:42:23',1,0),(20,'2020-06-08 17:42:27','2020-06-08 17:42:27',NULL,'3c7fde','erqwe','','127.0.0.1:63716','PostmanRuntime/7.24.1','2020-06-08 17:42:27',1,0),(21,'2020-06-08 17:42:45','2020-06-08 17:42:45',NULL,'3c7fde','新用户','','127.0.0.1:63716','PostmanRuntime/7.24.1','2020-06-08 17:42:45',1,0);
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
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `ek_server_wechat_oauths`
--

LOCK TABLES `ek_server_wechat_oauths` WRITE;
/*!40000 ALTER TABLE `ek_server_wechat_oauths` DISABLE KEYS */;
INSERT INTO `ek_server_wechat_oauths` VALUES (1,'2020-05-30 21:56:51','2020-05-30 21:56:51',NULL,0,'33_-jcaIABs7ankjH0PXom6JjKUAIaD8witT_qnenGgB9B42zSGMSxMwdjVo59jOIopbdQqFdE5PeA0BtyWnmECM-NVkh1VCMFLP8_lG_icmwc',1590854210,'33_3yQh4nY1pNmb0xEwCN454d22Dd4iXvkntMLZjEyX3DQOUFihR8nKVPu1vKDe8cy36jyLfGFrMt-zwtQD-R_Js9wU2f_j5TJi-qYWMHUN2OY','oUsta6PmPtCCs-XSuw02Q07p1OB0','snsapi_userinfo','小脏孩',1,'http://thirdwx.qlogo.cn/mmopen/vi_32/5RKP7fZBeXq7QzV6S9DpEEVe934sUdEOfV4rI3ia084Picayh9b2B5F3c9wefn8TRKOatUnq3ghb78DFlICOVicpw/132','河南','郑州','中国','',0);
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

-- Dump completed on 2020-06-11 17:06:46
