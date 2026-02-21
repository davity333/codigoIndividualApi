-- MySQL dump 10.13  Distrib 8.0.36, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: metal
-- ------------------------------------------------------
-- Server version	8.0.37

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
-- Table structure for table `messages`
--

DROP TABLE IF EXISTS `messages`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `messages` (
  `idmessage` int NOT NULL AUTO_INCREMENT,
  `content` varchar(555) DEFAULT NULL,
  `receiveId` int DEFAULT NULL,
  `senderId` int DEFAULT NULL,
  `timeMessage` bigint DEFAULT NULL,
  PRIMARY KEY (`idmessage`),
  KEY `fk_sender` (`senderId`),
  KEY `fk_receiver` (`receiveId`),
  CONSTRAINT `fk_receiver` FOREIGN KEY (`receiveId`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_sender` FOREIGN KEY (`senderId`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `messages`
--

LOCK TABLES `messages` WRITE;
/*!40000 ALTER TABLE `messages` DISABLE KEYS */;
INSERT INTO `messages` VALUES (1,'Hola tilín, probando el chat',2,1,1708045200),(2,'enviando mensaje a dahomey',2,1,1708045200),(3,'enviando mensaje a david',1,2,1708045200),(4,'hola  david soy jose',1,3,1708045200),(5,'hola jose, te habla david',3,1,1708045200);
/*!40000 ALTER TABLE `messages` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reservations`
--

DROP TABLE IF EXISTS `reservations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reservations` (
  `idReservation` bigint unsigned NOT NULL AUTO_INCREMENT,
  `studentId` int NOT NULL,
  `teacherId` int NOT NULL,
  `reservationDate` date NOT NULL,
  `reservationTime` time NOT NULL,
  `attendance` tinyint(1) DEFAULT NULL,
  `topic` varchar(45) NOT NULL,
  PRIMARY KEY (`idReservation`),
  UNIQUE KEY `idReservation` (`idReservation`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reservations`
--

LOCK TABLES `reservations` WRITE;
/*!40000 ALTER TABLE `reservations` DISABLE KEYS */;
/*!40000 ALTER TABLE `reservations` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(45) DEFAULT NULL,
  `email` varchar(45) DEFAULT NULL,
  `password` varchar(200) DEFAULT NULL,
  `firstname` varchar(45) DEFAULT NULL,
  `lastname` varchar(45) DEFAULT NULL,
  `rol` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'davity333','davity@gmail.com','$2a$10$LUcIOVGoqRBCdjY5/c/gDOnyE6e6VjEQZDDs5AaQrNxAZ4Zzo8uW6','David Reynold','Guzman Castro','Estudiante'),(2,'mey metal','mey@gmail.com','$2a$10$qCexprIB7D9KT3PqGPymEOdfNqsO7hL4JtiN1YiQ8fIxhi6UO11ZK','Dahomey','Cervantes Sosa','Estudiante'),(3,'Jose','jose@gmail.com','$2a$10$cHnJO.Nk9OfxL/5f9NgavubZIU.M0pX.9SMuUJjepPmfo5XoBvRqO','Jose','Lopez Portillo','Estudiante'),(5,'Ali','ali@gmail.com','$2a$10$h7DHMDk8ek3wYsOjETcz8.ET6kIplOf77It5Cv01ZlzbDnSoyMwKi','ali','lopez zunun','Docente');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2026-02-21  9:41:06
