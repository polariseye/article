/*
Navicat MySQL Data Transfer

Source Server         : 本地数据库
Source Server Version : 50712
Source Host           : localhost:3306
Source Database       : operation_i_dzz

Target Server Type    : MYSQL
Target Server Version : 50712
File Encoding         : 65001

Date: 2017-03-01 20:32:13
*/

CREATE DATABASE operation_i_dzz;
go;

-- ----------------------------
-- Table structure for tt
-- ----------------------------
DROP TABLE IF EXISTS `tt`;
CREATE TABLE `tt` (
  `Id` int(11) NOT NULL,
  `Name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
