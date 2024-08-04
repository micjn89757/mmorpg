CREATE DATABASE mmodb;
CREATE USER 'tester' identified by '123456';
GRANT ALL ON mmodb.* to tester;
USE mmodb;

CREATE TABLE `user` (
    `id` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户id, 自增',
    `account` varchar(15) NOT NULL COMMENT '用户名',
    `password` char(32) NOT NULL COMMENT '密码的md5',
    `created_time` DATETIME DEFAULT
    primary key (id)
    unique key idx_name(`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT '用户登录'



CREATE TABLE `actor` (
  `id` int NOT NULL AUTO_INCREMENT COMMENT '角色id',
  `account` varchar(15) NOT NULL COMMENT '所属账号',
  `scene_id` int DEFAULT NULL COMMENT '场景id',
  `created_time` datetime DEFAULT NULL COMMENT 'Create Time',
  `nickname` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名',
  PRIMARY KEY (`id`),
  KEY `account` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

