create database beyond_reply;
use beyond_reply;

CREATE TABLE `reply_count` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `biz_id` varchar(64) NOT NULL DEFAULT '' COMMENT '业务ID',
    `target_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '评论目标id',
    `reply_num` int(11) NOT NULL DEFAULT '0' COMMENT '评论总数',
    `reply_root_num` int(11) NOT NULL DEFAULT '0' COMMENT '根评论总数',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='评论计数';

CREATE TABLE `reply` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `biz_id` varchar(64) NOT NULL DEFAULT '' COMMENT '业务ID',
    `target_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '评论目标id',
    `reply_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '评论用户ID',
    `be_reply_user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '被回复用户ID',
    `parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '父评论ID',
    `content` text COLLATE utf8_unicode_ci NOT NULL COMMENT '内容',
    `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0:正常 1:删除',
    `like_num` int(11) NOT NULL DEFAULT '0' COMMENT '点赞数',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='评论表';