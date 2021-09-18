-- MySQL dump 10.13  Distrib 5.5.62, for Win64 (AMD64)
--
-- Host: localhost    Database: quiz_3
-- ------------------------------------------------------
-- Server version	5.5.5-10.4.18-MariaDB

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
-- Table structure for table `books`
--

DROP TABLE IF EXISTS `books`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `books` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `image_url` text NOT NULL,
  `release_year` int(11) NOT NULL,
  `price` varchar(255) NOT NULL,
  `total_page` varchar(255) NOT NULL,
  `kategori_ketebalan` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `books`
--

LOCK TABLES `books` WRITE;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` VALUES (1,'LOTR','Test Description','https://i2.wp.com/ytimg.googleusercontent.com/vi/NVuWw7xcIuM/hqdefault.jpg',2006,'Rp.0,-','189','Sedang','2021-09-19 01:24:22','2021-09-19 01:24:22'),(2,'Michael Jackson Heehee','The Transformation of Michael Jackson','https://i2.wp.com/ytimg.googleusercontent.com/vi/NVuWw7xcIuM/hqdefault.jpg',2006,'Rp.0,-','400','Tebal','2021-09-19 01:27:21','2021-09-19 01:27:21'),(3,'WTF MAN','The Transformation of Michael Jackson','https://i2.wp.com/ytimg.googleusercontent.com/vi/NVuWw7xcIuM/hqdefault.jpg',2006,'Rp.100.000,-','400','Tebal','2021-09-19 01:44:44','2021-09-19 01:44:44'),(4,'Kok begini man','Gimana cara convert Rp100.000 ke integer tanpa membuat go panic','https://i2.wp.com/ytimg.googleusercontent.com/vi/NVuWw7xcIuM/hqdefault.jpg',2006,'Rp.200.000,-','400','Tebal','2021-09-19 01:46:19','2021-09-19 02:37:07'),(5,'WTFsdsadasdasdsasss MAN','The Transformation of Michael Jackson','https://i2.wp.com/ytimg.googleusercontent.com/vi/NVuWw7xcIuM/hqdefault.jpg',2006,'Rp.10.000,-','400','Tebal','2021-09-19 01:56:07','2021-09-19 01:56:07'),(7,'sss ssssMAN','The Transformation of Michael Jackson','https://i2.wp.com/ytimg.googleusercontent.com/vi/NVuWw7xcIuM/hqdefault.jpg',2006,'Rp.101.000,-','400','Tebal','2021-09-19 02:18:35','2021-09-19 02:18:35'),(8,'Makan','The Transformation of Michael Jackson','https://i2.wp.com/ytimg.googleusercontent.com/vi/NVuWw7xcIuM/hqdefault.jpg',2006,'Rp.1.212,-','400','Tebal','2021-09-19 02:20:15','2021-09-19 02:20:15');
/*!40000 ALTER TABLE `books` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'quiz_3'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-09-19  3:25:28
