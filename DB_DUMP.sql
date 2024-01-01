/*
SQLyog Community v13.2.0 (64 bit)
MySQL - 10.4.28-MariaDB : Database - go_blogs
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
/*Table structure for table `blogs` */

DROP TABLE IF EXISTS `blogs`;

CREATE TABLE `blogs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `category_id` bigint(20) unsigned DEFAULT NULL,
  `title` varchar(256) DEFAULT NULL,
  `detail` text DEFAULT NULL,
  `image` varchar(256) DEFAULT NULL,
  `views` bigint(20) DEFAULT 0,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_blogs_deleted_at` (`deleted_at`)
) ENGINE=MyISAM AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `blogs` */

insert  into `blogs`(`id`,`category_id`,`title`,`detail`,`image`,`views`,`created_at`,`updated_at`,`deleted_at`) values 
(1,1,'Test Blog One','Here is some details of the blog.',NULL,23,'2023-11-07 16:49:20.000','2023-11-07 16:49:23.000',NULL),
(2,2,'Test Two Blog','<div>Here is some details of the blog</div>','',43,'2023-12-03 10:34:44.823','2023-12-03 10:34:44.823',NULL),
(3,3,'Third Blog Updated','<div><b>Here is Title</b></div><div>Here is some details</div>','',234,'2023-12-03 10:35:24.707','2023-12-03 10:47:00.613',NULL),
(4,6,'Image Blog','<div>Here is some details of the image blog</div>','168908.jpg',22,'2023-12-03 13:01:04.269','2023-12-16 07:46:33.076',NULL),
(5,5,'sdfsdf','<div>sdfsdf</div>','201517.png',323,'2023-12-03 13:04:46.292','2023-12-03 13:04:46.292',NULL),
(6,6,'dsdfsdf444','<div>sdfsdf dgdfgdfgf</div>','183968.jpg',43,'2023-12-03 13:06:49.115','2023-12-16 07:45:40.829',NULL),
(7,3,'Latest Blog','<div>Here is the details of the blog</div>','201518.jpg',1234,'2023-12-09 09:03:35.812','2023-12-09 09:03:35.812',NULL);

/*Table structure for table `categories` */

DROP TABLE IF EXISTS `categories`;

CREATE TABLE `categories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(256) DEFAULT NULL,
  `status` bigint(20) unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_categories_deleted_at` (`deleted_at`)
) ENGINE=MyISAM AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `categories` */

insert  into `categories`(`id`,`title`,`status`,`created_at`,`updated_at`,`deleted_at`) values 
(1,'Test Category',0,'2023-11-07 16:47:58.000','2023-11-07 16:48:00.000',NULL),
(2,'Test Category 2',1,'2023-11-07 16:47:58.000','2023-11-07 16:48:00.000',NULL),
(3,'Second Category',1,'2023-11-07 16:47:58.000','2023-11-07 16:48:00.000',NULL),
(4,'Fourth Category',0,'2023-11-07 16:47:58.000','2023-11-07 16:48:00.000',NULL),
(5,'Youth Category',1,'2023-11-07 16:47:58.000','2023-11-07 16:48:00.000',NULL),
(6,'Vegetable',1,'2023-11-07 16:47:58.000','2023-11-07 16:48:00.000',NULL),
(7,'Fruits',1,'2023-11-07 16:47:58.000','2023-11-07 16:48:00.000',NULL),
(8,'Mango',1,'2023-11-07 16:47:58.000','2023-11-07 16:48:00.000',NULL),
(9,'Design',1,'2023-11-07 16:47:58.000','2023-11-07 16:48:00.000',NULL),
(10,'Intetrior Design',1,'2023-11-07 16:47:58.000','2023-11-07 16:48:00.000',NULL),
(11,'Decorator 1234',1,'2023-11-06 18:53:05.000','2023-11-30 14:01:55.459',NULL),
(12,'Workshop',1,'2023-11-07 16:47:58.000','2023-11-07 16:48:00.000',NULL),
(13,'Test',1,'2023-11-30 11:55:58.981','2023-11-30 11:55:58.981',NULL),
(14,'Test 2',0,'2023-11-30 11:56:21.730','2023-11-30 11:56:21.730',NULL),
(15,'Test 2',0,'2023-11-30 11:57:31.536','2023-11-30 11:57:31.536',NULL),
(16,'Test Three',1,'2023-11-30 11:58:31.376','2023-11-30 11:58:31.376',NULL),
(17,'Test Three',1,'2023-11-30 11:59:03.737','2023-11-30 11:59:03.737',NULL),
(18,'Test 2',0,'2023-11-30 11:59:06.851','2023-11-30 11:59:06.851',NULL),
(19,'ttt',1,'2023-11-30 11:59:20.107','2023-11-30 11:59:20.107','2023-11-30 15:14:11.719'),
(20,'ttt 2',1,'2023-11-30 11:59:46.724','2023-11-30 11:59:46.724','2023-11-30 15:14:31.584'),
(21,'ttt 3',1,'2023-11-30 12:00:22.304','2023-11-30 12:00:22.304',NULL),
(22,'Decorator Updated 234',1,'2023-11-30 13:36:46.882','2023-11-30 13:45:11.916','2023-11-30 15:13:35.389');

/*Table structure for table `users` */

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) DEFAULT NULL,
  `email` varchar(256) DEFAULT NULL,
  `password` varchar(256) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`) USING HASH,
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=MyISAM AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*Data for the table `users` */

insert  into `users`(`id`,`name`,`email`,`password`,`created_at`,`updated_at`,`deleted_at`) values 
(1,'Test User','admin2@admin.com','$2a$10$5IdOoUwKDfNxyfYzJzMGruy8wyHxJWdLAZMgQfxcKgEVS32amMdq.','2023-11-16 10:56:37.231','2023-11-16 10:56:37.231',NULL),
(2,'Administrator','admin@admin.com','$2a$10$7Jo8yyESu3UwrTbS64eJ0.ENqA9nS3RKr1CiA92HoY8YuWz.Y0KFG','2023-11-16 10:58:05.191','2023-11-16 10:58:05.191',NULL),
(3,'Test User Updated','test@test.com','$2a$10$APaOjYV7SC84gYGQgScNt.HcXwlNizCllhbBFVhTM9M/.RtM0PEji','2023-12-03 09:18:43.838','2023-12-03 09:24:39.947','2023-12-17 14:23:05.872'),
(4,'Test User','sadiqse024@gmail.com','$2a$10$4Tomc8Rac7sgmcVZT1io/ugedsL28NJ5iPzSiIHDgPZUB5gcacZFW','2023-12-17 13:43:08.038','2023-12-17 13:43:08.038',NULL),
(5,'khan','sadiqse@gmail.com','$2a$10$ZBuhFR8mRYcuh.oLv2fTXel5aAGcp4impRFwAT87dg4erbFB7ZaWm','2023-12-17 13:43:39.295','2023-12-17 14:22:52.547',NULL),
(6,'Test User','admin123@admin.com','$2a$10$yDwI4jHe1UA.huXGWgvwFOhkAKCNzu6K3j1n7rotMEjwHiiZH4I3S','2023-12-17 14:09:45.724','2023-12-17 14:20:16.483',NULL);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
