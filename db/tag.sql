create database beyond_tag;
use beyond_tag;

CREATE TABLE `tag` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `tag_name` varchar(32) NOT NULL DEFAULT '' COMMENT '标签名',
    `tag_desc` varchar(128) NOT NULL DEFAULT '' COMMENT '标签描述',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='标签表';

CREATE TABLE `tag_resource` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `biz_id` varchar(64) NOT NULL DEFAULT '' COMMENT '业务ID',
    `target_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '内容id',
    `tag_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '标签id',
    `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户ID',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='标签资源表';

