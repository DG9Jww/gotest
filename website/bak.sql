-- MySQL dump 10.13  Distrib 5.7.26, for Win64 (x86_64)
--
-- Host: localhost    Database: dachuang2
-- ------------------------------------------------------
-- Server version	5.7.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Current Database: `dachuang2`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `dachuang2` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;

USE `dachuang2`;

--
-- Table structure for table `admins`
--

DROP TABLE IF EXISTS `admins`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `admins` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ad_name` varchar(32) COLLATE utf8_unicode_ci NOT NULL,
  `ad_password` varchar(32) COLLATE utf8_unicode_ci NOT NULL,
  `ad_right` tinyint(4) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `admins`
--

LOCK TABLES `admins` WRITE;
/*!40000 ALTER TABLE `admins` DISABLE KEYS */;
INSERT INTO `admins` VALUES (2,'admin','D764232FE344FC680A0F8BE0A51A012C',0),(3,'dale','6B44696917A612202E5EA680D4344BD2',0),(4,'jack','495AC10A83963D212BEF8BDC2D5E949D',0),(5,'jack','D3CABBFE663CB5CB9EED2780E0AAEF15',0);
/*!40000 ALTER TABLE `admins` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `articles`
--

DROP TABLE IF EXISTS `articles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `articles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `a_title` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `a_content` longtext COLLATE utf8_unicode_ci NOT NULL,
  `a_author` varchar(32) COLLATE utf8_unicode_ci NOT NULL,
  `a_publishTime` bigint(20) NOT NULL,
  `a_categoryID` tinyint(1) NOT NULL,
  `a_cover` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `a_status` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `articles`
--

LOCK TABLES `articles` WRITE;
/*!40000 ALTER TABLE `articles` DISABLE KEYS */;
INSERT INTO `articles` VALUES (1,'first','mycontentsssssssssssssssssssssssssssssssssss','dale',1644563655537,123,'',1),(3,'test','i was been changed','dale',1644569557740,1,'',1),(4,'test','i was been changed','dale',1644569562333,1,'http://xx.x..xxx.jpg',1),(5,'test','i was been changed','dale',1644569562915,1,'http://xx.x..xxx.jpg',1),(6,'test','i was been changed','dale',1644569563492,1,'',1),(8,'test','i was been changed','dale',1644569568228,1,'http://xx.x..xxx.jpg',1),(9,'test','i was been changed','dale',1644569568778,1,'',1),(10,'test','i was been changed','dale',1644569569260,1,'',1),(11,'test','i was been changed','dale',1644569569703,1,'',1),(12,'test','i was been changed','dale',1644569570163,1,'',1),(13,'test','i was been changed','dale',1644569570645,1,'',1),(14,'test','i was been changed','dale',1644569575513,1,'',1),(15,'test','i was been changed','dale',1644569576086,1,'',1),(16,'test','i was been changed','dale',1644569577783,1,'',1),(17,'test','i was been changed','dale',1644569582589,2,'',1),(18,'test','i was been changed','dale',1644569583101,2,'',1),(19,'test','i was been changed','dale',1644569583635,2,'',1),(20,'test','i was been changed','dale',1644569584148,2,'',1),(21,'test','i was been changed','dale',1644569584724,2,'',1),(22,'test','i was been changed','dale',1644569585268,2,'',1),(23,'test','i was been changed','dale',1644569585947,2,'',1),(24,'test','i was been changed','dale',1644569586685,2,'',1),(25,'test','i was been changed','dale',1644569587488,2,'',1),(26,'test','i was been changed','dale',1644569685366,2,'',1),(27,'test','i was been changed','dale',1644569686004,2,'',1),(28,'test','i was been changed','dale',1644569686495,2,'',1),(29,'test','i was been changed','dale',1644569686996,2,'',1),(30,'test','i was been changed','dale',1644569687457,2,'',1),(31,'test','i was been changed','dale',1644569687940,2,'',1),(32,'test','i was been changed','dale',1644569688412,2,'',1),(33,'test','i was been changed','dale',1644569688897,2,'',1),(34,'test','i was been changed','dale',1644569689373,2,'',1),(35,'test','i was been changed','dale',1644569689850,2,'',1),(36,'test','i was been changed','dale',1644569690332,2,'',1),(37,'test','i was been changed','dale',1644569692733,2,'',1),(38,'test','i was been changed','jack',1644569697380,2,'',1),(39,'test','i was been changed','jack',1644569697982,2,'',1),(40,'test','i was been changed','jack',1644569698456,2,'',1),(41,'','','',1644822455688,0,'',0);
/*!40000 ALTER TABLE `articles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `test`
--

DROP TABLE IF EXISTS `test`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `test` (
  `id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `test`
--

LOCK TABLES `test` WRITE;
/*!40000 ALTER TABLE `test` DISABLE KEYS */;
/*!40000 ALTER TABLE `test` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-02-17 16:59:45
