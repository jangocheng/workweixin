CREATE TABLE `todo_list` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`todo_name` varchar(1024) NOT NULL COMMENT 'todo事件',
	`create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`finish_time` datetime DEFAULT NULL,
	`active` tinyint (1) NOT NULL DEFAULT '0' COMMENT '完成标志,0未完成，1已完成',
	PRIMARY KEY (`id`)) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = 'todo列表';