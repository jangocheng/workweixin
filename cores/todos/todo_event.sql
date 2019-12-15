CREATE TABLE `todo_list` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id` varchar(256) NOT NULL DEFAULT '0' COMMENT 'users 表关联',
    `todo_name` varchar(512) NOT NULL DEFAULT '' COMMENT 'todo事件',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `finish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `active` tinyint NOT NULL DEFAULT '0' COMMENT '完成标志,0未完成，1已完成',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_todo` (`user_id`, `todo_name`)
                         ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'todo列表'