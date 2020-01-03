CREATE TABLE `users` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` VARCHAR(256) NOT NULL COMMENT 'user id',
    `user_name` VARCHAR(256) NOT NULL DEFAULT '' COMMENT 'user name',
    `gender` TINYINT NOT NULL COMMENT '性别 1表示男性，2表示女性',
    `state` TINYINT NOT NULL DEFAULT '4' COMMENT '激活状态：1=激活或关注 2=禁用 4=未激活',
    `email` VARCHAR(128) NOT NULL DEFAULT '' COMMENT 'email',
    `mobile` VARCHAR(16) NOT NULL DEFAULT '' COMMENT '手机号码',
    `create_time` BIGINT NOT NULL DEFAULT '0' COMMENT '创建/更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_id` (`user_id`)) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户表';