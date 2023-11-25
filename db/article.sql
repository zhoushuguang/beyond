create database beyond_article;
use beyond_article;

CREATE TABLE `article` (
    `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
    `content` text COLLATE utf8_unicode_ci NOT NULL COMMENT '内容',
    `cover` varchar(255) NOT NULL DEFAULT '' COMMENT '封面',
    `description` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
    `author_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '作者ID',
    `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态 0:待审核 1:审核不通过 2:可见 3:用户删除',
    `comment_num` int(11) NOT NULL DEFAULT '0' COMMENT '评论数',
    `like_num` int(11) NOT NULL DEFAULT '0' COMMENT '点赞数',
    `collect_num` int(11) NOT NULL DEFAULT '0' COMMENT '收藏数',
    `view_num` int(11) NOT NULL DEFAULT '0' COMMENT '浏览数',
    `share_num` int(11) NOT NULL DEFAULT '0' COMMENT '分享数',
    `tag_ids` varchar(255) NOT NULL DEFAULT '' COMMENT '标签ID',
    `publish_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发布时间',
    `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
    PRIMARY KEY (`id`),
    KEY `ix_author_id` (`author_id`),
    KEY `ix_update_time` (`update_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='文章表';


insert into article(title, content, author_id, like_num, publish_time) values ('文章标题1', '文章内容1', 1, 1, '2023-11-25 17:01:01');
insert into article(title, content, author_id, like_num, publish_time) values ('文章标题2', '文章内容2', 1, 10, '2023-11-25 15:01:01');



