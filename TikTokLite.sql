/*
 Navicat Premium Data Transfer

 Source Server         : docker
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : 192.168.1.12:3306
 Source Schema         : TikTokLite

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 09/06/2022 14:52:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
SET global log_bin_trust_function_creators = 1;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `comment_id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  `comment` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`comment_id`) USING BTREE,
  INDEX `commentUser`(`user_id`) USING BTREE,
  INDEX `commentVideo`(`video_id`) USING BTREE,
  CONSTRAINT `commentuser` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `commentvideo` FOREIGN KEY (`video_id`) REFERENCES `videos` (`video_id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for favorites
-- ----------------------------
DROP TABLE IF EXISTS `favorites`;
CREATE TABLE `favorites`  (
  `favorite_id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `video_id` bigint NOT NULL,
  PRIMARY KEY (`favorite_id`) USING BTREE,
  INDEX `favoriteUser`(`user_id`) USING BTREE,
  INDEX `favoriteVideo`(`video_id`) USING BTREE,
  CONSTRAINT `favoriteuser` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `favoritevideo` FOREIGN KEY (`video_id`) REFERENCES `videos` (`video_id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for relations
-- ----------------------------
DROP TABLE IF EXISTS `relations`;
CREATE TABLE `relations`  (
  `relation_id` bigint NOT NULL AUTO_INCREMENT,
  `follow_id` bigint NOT NULL,
  `follower_id` bigint NOT NULL,
  PRIMARY KEY (`relation_id`) USING BTREE,
  INDEX `FollowId`(`follow_id`) USING BTREE,
  INDEX `FollowerId`(`follower_id`) USING BTREE,
  CONSTRAINT `followerid` FOREIGN KEY (`follower_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE RESTRICT,
  CONSTRAINT `followid` FOREIGN KEY (`follow_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `user_id` bigint NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `follow_count` bigint NULL DEFAULT NULL,
  `follower_count` bigint NULL DEFAULT NULL,
  `total_favorited` bigint NULL DEFAULT NULL,
  `favorite_count` bigint NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `background_image` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `signature` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `video_id` bigint NOT NULL AUTO_INCREMENT,
  `author_id` bigint NOT NULL,
  `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `favorite_count` bigint NULL DEFAULT NULL,
  `comment_count` bigint NULL DEFAULT NULL,
  `publish_time` bigint NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`video_id`) USING BTREE,
  INDEX `user`(`author_id`) USING BTREE,
  CONSTRAINT `authorid` FOREIGN KEY (`author_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 0 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Triggers structure for table comments
-- ----------------------------
DROP TRIGGER IF EXISTS `comment_action`;
delimiter ;;
CREATE TRIGGER `comment_action` AFTER INSERT ON `comments` FOR EACH ROW update videos set comment_count = comment_count + 1 where videos.video_id = new.video_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table comments
-- ----------------------------
DROP TRIGGER IF EXISTS `delete_comment`;
delimiter ;;
CREATE TRIGGER `delete_comment` AFTER DELETE ON `comments` FOR EACH ROW update videos set comment_count = comment_count - 1 where videos.video_id = old.video_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table favorites
-- ----------------------------
DROP TRIGGER IF EXISTS `like_action`;
delimiter ;;
CREATE TRIGGER `like_action` AFTER INSERT ON `favorites` FOR EACH ROW update videos set favorite_count = favorite_count + 1 where videos.video_id = new.video_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table favorites
-- ----------------------------
DROP TRIGGER IF EXISTS `unlike_action`;
delimiter ;;
CREATE TRIGGER `unlike_action` AFTER DELETE ON `favorites` FOR EACH ROW update videos set favorite_count = favorite_count - 1 where videos.video_id = old.video_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table favorites
-- ----------------------------
DROP TRIGGER IF EXISTS `faved_count`;
delimiter ;;
CREATE TRIGGER `faved_count` AFTER INSERT ON `favorites` FOR EACH ROW update users,(select author_id from videos,favorites where videos.video_id = new.video_id) a set total_favorited = total_favorited + 1 where  a.author_id = users.user_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table favorites
-- ----------------------------
DROP TRIGGER IF EXISTS `unfaved_count`;
delimiter ;;
CREATE TRIGGER `unfaved_count` AFTER DELETE ON `favorites` FOR EACH ROW update users,(select author_id from videos,favorites where videos.video_id = old.video_id) a set total_favorited = total_favorited - 1 where  a.author_id = users.user_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table favorites
-- ----------------------------
DROP TRIGGER IF EXISTS `fav_count`;
delimiter ;;
CREATE TRIGGER `fav_count` AFTER INSERT ON `favorites` FOR EACH ROW update users set users.favorite_count = users.favorite_count + 1 where users.user_id = new.user_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table favorites
-- ----------------------------
DROP TRIGGER IF EXISTS `unfav_count`;
delimiter ;;
CREATE TRIGGER `unfav_count` AFTER DELETE ON `favorites` FOR EACH ROW update users set users.favorite_count = users.favorite_count - 1 where users.user_id = old.user_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table relations
-- ----------------------------
DROP TRIGGER IF EXISTS `false_follow_action`;
delimiter ;;
CREATE TRIGGER `false_follow_action` AFTER DELETE ON `relations` FOR EACH ROW update users set follow_count = follow_count - 1 where users.user_id = old.follow_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table relations
-- ----------------------------
DROP TRIGGER IF EXISTS `false_follower_action`;
delimiter ;;
CREATE TRIGGER `false_follower_action` AFTER DELETE ON `relations` FOR EACH ROW update users set follower_count = follower_count - 1 where users.user_id = old.follower_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table relations
-- ----------------------------
DROP TRIGGER IF EXISTS `follow_action`;
delimiter ;;
CREATE TRIGGER `follow_action` AFTER INSERT ON `relations` FOR EACH ROW update users set follow_count = follow_count + 1 where users.user_id = new.follow_id
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table relations
-- ----------------------------
DROP TRIGGER IF EXISTS `follower_action`;
delimiter ;;
CREATE TRIGGER `follower_action` AFTER INSERT ON `relations` FOR EACH ROW update users set follower_count = follower_count + 1 where users.user_id = new.follower_id
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
