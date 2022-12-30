/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:3306
 Source Schema         : iot_smart_home

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 30/12/2022 22:18:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for iot_app
-- ----------------------------
DROP TABLE IF EXISTS `iot_app`;
CREATE TABLE `iot_app` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `app_id` varchar(32) NOT NULL,
  `app_secret` varchar(64) NOT NULL,
  `app_type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1代表设备，2代表对外开放的app',
  PRIMARY KEY (`id`,`app_id`),
  UNIQUE KEY `idx_iot_app_app_id` (`app_id`,`app_secret`),
  KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for iot_app_proxy
-- ----------------------------
DROP TABLE IF EXISTS `iot_app_proxy`;
CREATE TABLE `iot_app_proxy` (
  `id` int NOT NULL AUTO_INCREMENT,
  `app_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `component_app_id` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for iot_device
-- ----------------------------
DROP TABLE IF EXISTS `iot_device`;
CREATE TABLE `iot_device` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `user_id` bigint unsigned NOT NULL,
  `app_id` bigint unsigned NOT NULL,
  `type` tinyint unsigned DEFAULT '1',
  `type_name` varchar(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `device_status` tinyint unsigned NOT NULL DEFAULT '1',
  `device_cur_status` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `online_status` tinyint NOT NULL DEFAULT '0',
  `last_ip` varchar(24) DEFAULT '',
  `latest_visit` varchar(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `gateway_device_id` int DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `fk_iot_device_app` (`app_id`),
  CONSTRAINT `fk_iot_device_app` FOREIGN KEY (`app_id`) REFERENCES `iot_app` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for iot_device_operate_log
-- ----------------------------
DROP TABLE IF EXISTS `iot_device_operate_log`;
CREATE TABLE `iot_device_operate_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `device_gateway_id` int unsigned DEFAULT '0',
  `device_id` int NOT NULL,
  `device_type` varchar(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `source` varchar(12) NOT NULL,
  `operate_name` varchar(64) NOT NULL,
  `operate_number` varchar(64) NOT NULL,
  `operate_payload` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `operate_level` tinyint NOT NULL DEFAULT '0',
  `response_payload` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci,
  `operate_time` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `response_ip` varchar(24) DEFAULT NULL,
  `response_port` varchar(12) DEFAULT NULL,
  `response_time` varchar(24) DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=130 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for iot_device_report_log
-- ----------------------------
DROP TABLE IF EXISTS `iot_device_report_log`;
CREATE TABLE `iot_device_report_log` (
  `id` int NOT NULL AUTO_INCREMENT,
  `device_gateway_id` int unsigned DEFAULT '0',
  `device_id` int NOT NULL,
  `device_type` varchar(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `source` varchar(12) NOT NULL,
  `report_name` varchar(64) NOT NULL,
  `report_number` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `report_payload` varchar(500) NOT NULL,
  `report_level` tinyint NOT NULL DEFAULT '0',
  `report_time` datetime DEFAULT NULL,
  `report_ip` varchar(24) DEFAULT NULL,
  `report_port` varchar(12) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=154 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for iot_face_model
-- ----------------------------
DROP TABLE IF EXISTS `iot_face_model`;
CREATE TABLE `iot_face_model` (
  `id` int NOT NULL AUTO_INCREMENT,
  `device_id` int NOT NULL,
  `user_name` varchar(64) NOT NULL,
  `status` tinyint NOT NULL DEFAULT '0',
  `urls` text NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for iot_setting
-- ----------------------------
DROP TABLE IF EXISTS `iot_setting`;
CREATE TABLE `iot_setting` (
  `key` varchar(120) NOT NULL,
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for iot_user
-- ----------------------------
DROP TABLE IF EXISTS `iot_user`;
CREATE TABLE `iot_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `mobile` varchar(11) NOT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(64) NOT NULL,
  `salt` varchar(12) NOT NULL,
  `status` tinyint unsigned NOT NULL DEFAULT '1',
  `last_ip` varchar(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `register_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
  `latest_visit` varchar(24) DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `index2` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `token` char(43) NOT NULL,
  `data` blob NOT NULL,
  `expiry` timestamp(6) NOT NULL,
  PRIMARY KEY (`token`),
  KEY `sessions_expiry_idx` (`expiry`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
