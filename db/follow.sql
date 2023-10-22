create database beyond_follow;
use beyond_follow;

CREATE TABLE `follow` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户ID',
    `followed_user_id` bigint(20) UNSIGNED NOT NULL COMMENT '被关注用户ID',
    `follow_status` tinyint(1) UNSIGNED NOT NULL DEFAULT '1' COMMENT '关注状态：1-关注，2-取消关注',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id_followed_user_id` (`user_id`,`followed_user_id`),
    KEY `ix_followed_user_id` (`followed_user_id`),
    KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT '关注表';

CREATE TABLE `follow_count` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户ID',
    `follow_count` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '关注数',
    `fans_count` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '粉丝数',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id` (`user_id`),
    KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT '关注计数表';