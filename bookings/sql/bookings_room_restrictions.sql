-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: localhost    Database: bookings
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `room_restrictions`
--

DROP TABLE IF EXISTS `room_restrictions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `room_restrictions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `start_date` date DEFAULT NULL,
  `end_date` date DEFAULT NULL,
  `room_id` int NOT NULL,
  `reservation_id` int DEFAULT NULL,
  `restriction_id` int NOT NULL,
  `created_at` timestamp(6) NULL DEFAULT NULL,
  `updated_at` timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `key3_idx` (`room_id`),
  KEY `key2_idx` (`restriction_id`),
  KEY `key1_idx` (`reservation_id`),
  CONSTRAINT `key1` FOREIGN KEY (`reservation_id`) REFERENCES `reservations` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `key2` FOREIGN KEY (`restriction_id`) REFERENCES `restrictions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `key3` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `room_restrictions`
--

LOCK TABLES `room_restrictions` WRITE;
/*!40000 ALTER TABLE `room_restrictions` DISABLE KEYS */;
INSERT INTO `room_restrictions` VALUES (1,'2024-06-05','2024-06-19',1,1,1,'2024-07-29 01:27:54.214759','2024-07-29 01:27:54.214759'),(2,'2024-07-01','2024-07-04',1,2,1,'2024-07-29 01:28:51.919877','2024-07-29 01:28:51.919877'),(3,'2024-07-01','2024-07-04',2,3,1,'2024-07-29 01:29:28.730960','2024-07-29 01:29:28.730960'),(4,'2024-07-18','2024-07-20',1,4,1,'2024-07-29 01:30:37.703915','2024-07-29 01:30:37.703915'),(5,'2024-07-10','2024-07-11',1,NULL,2,'2024-07-29 01:31:05.955032','2024-07-29 01:31:05.955032'),(6,'2024-07-27','2024-07-28',1,NULL,2,'2024-07-29 01:31:05.959236','2024-07-29 01:31:05.959236'),(7,'2024-07-19','2024-07-20',2,NULL,2,'2024-07-29 01:31:05.961774','2024-07-29 01:31:05.961774'),(8,'2024-07-29','2024-07-30',2,NULL,2,'2024-07-29 01:31:05.963793','2024-07-29 01:31:05.963793'),(9,'2024-07-30','2024-07-31',2,NULL,2,'2024-07-29 01:31:05.965323','2024-07-29 01:31:05.965323'),(10,'2024-07-11','2024-07-12',1,NULL,2,'2024-07-29 01:31:05.965832','2024-07-29 01:31:05.965832'),(11,'2024-07-09','2024-07-10',2,NULL,2,'2024-07-29 01:31:05.968282','2024-07-29 01:31:05.968282'),(12,'2024-07-18','2024-07-19',2,NULL,2,'2024-07-29 01:31:05.969867','2024-07-29 01:31:05.969867'),(13,'2024-07-20','2024-07-21',2,NULL,2,'2024-07-29 01:31:05.972268','2024-07-29 01:31:05.972268');
/*!40000 ALTER TABLE `room_restrictions` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-07-29 12:36:42
