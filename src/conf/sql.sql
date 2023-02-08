package mysql

CREATE TABLE `users` (
`id`  varchar(64) CHARACTER SET utf8mb4 NOT NULL ,
`username`  varchar(20) CHARACTER SET utf8mb4  NOT NULL COMMENT '用户名' ,
`password`  varchar(64) CHARACTER SET utf8mb4  NOT NULL COMMENT '密码' ,
`face_image`  varchar(255) CHARACTER SET utf8mb4 NULL DEFAULT NULL COMMENT '我的头像，如果没有默认给一张' ,
`nickname`  varchar(20) CHARACTER SET utf8mb4  NOT NULL COMMENT '昵称' ,
`fans_counts`  int(11) NULL DEFAULT 0 COMMENT '我的粉丝数量' ,
`follow_counts`  int(11) NULL DEFAULT 0 COMMENT '我关注的人总数' ,
PRIMARY KEY (`id`),
UNIQUE INDEX `id` (`id`) USING BTREE ,
UNIQUE INDEX `username` (`username`) USING BTREE
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4
ROW_FORMAT=DYNAMIC

CREATE TABLE `videos` (
                          `id`  varchar(64) CHARACTER SET utf8mb4   NOT NULL ,
                          `user_id`  varchar(64) CHARACTER SET utf8mb4  NOT NULL COMMENT '发布者id' ,
                          `video_desc`  varchar(128) CHARACTER SET utf8mb4   NULL DEFAULT NULL COMMENT '视频描述' ,
                          `video_url`  varchar(255) CHARACTER SET utf8mb4   NOT NULL COMMENT '视频链接' ,
                          `cover_path`  varchar(255) CHARACTER SET utf8mb4  NULL DEFAULT NULL COMMENT '视频封面图链接' ,
                          `like_counts`  bigint(20) NOT NULL DEFAULT 0 COMMENT '喜欢/赞美的数量' ,
                          `create_time`  datetime NOT NULL COMMENT '创建时间' ,
                          `is_delete` tinyint null comment '是否删除（软删除）',
                          PRIMARY KEY (`id`)
)
    ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4
COMMENT='视频信息表'
ROW_FORMAT=DYNAMIC


CREATE TABLE `comments` (
                            `id`  varchar(20) CHARACTER SET utf8mb4 NOT NULL ,
                            `father_comment_id`  varchar(20) CHARACTER SET utf8mb4  NULL DEFAULT NULL ,
                            `to_user_id`  varchar(20) CHARACTER SET utf8mb4  NULL DEFAULT NULL ,
                            `video_id`  varchar(20) CHARACTER SET utf8mb4 NOT NULL COMMENT '视频id' ,
                            `from_user_id`  varchar(20) CHARACTER SET utf8mb4  NOT NULL COMMENT '留言者，评论的用户id' ,
                            `comment`  text CHARACTER SET utf8mb4 NOT NULL COMMENT '评论内容' ,
                            `create_time`  datetime NOT NULL ,
                            `is_delete` tinyint null comment '是否删除（软删除）',
                            PRIMARY KEY (`id`)
)
    ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4
COMMENT='评论表'
ROW_FORMAT=DYNAMIC


create table `favorite` (
                            `id`  varchar(64) CHARACTER SET utf8mb4  NOT NULL ,
                            `user_id`  varchar(64) CHARACTER SET utf8mb4 NOT NULL COMMENT '用户' ,
                            `video_id`  varchar(64) CHARACTER SET utf8mb4 NOT NULL COMMENT '视频' ,
                            `is_delete` tinyint null comment '是否删除/取消点赞（软删除）',
                            PRIMARY KEY (`id`),
                            UNIQUE INDEX `user_id` (`user_id`, `video_id`) USING BTREE
)
    ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4
COMMENT='用户点赞视频关联表'
ROW_FORMAT=DYNAMIC
;


CREATE TABLE `users_fans` (
                              `id`  varchar(64) CHARACTER SET utf8mb4  NOT NULL ,
                              `user_id`  varchar(64) CHARACTER SET utf8mb4 NOT NULL COMMENT '用户' ,
                              `fan_id`  varchar(64) CHARACTER SET utf8mb4 NOT NULL COMMENT '粉丝' ,
                              `is_delete` tinyint null comment '是否删除/取消关注（软删除）',
                              PRIMARY KEY (`id`),
                              UNIQUE INDEX `user_id` (`user_id`, `fan_id`) USING BTREE
)
    ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4
COMMENT='用户粉丝关联关系表'
ROW_FORMAT=DYNAMIC
;