CREATE TABLE `students` (
    `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `name` varchar(63) NOT NULL COMMENT '姓名',
    `age` int NOT NULL COMMENT '年龄',
    `gender` int NOT NULL COMMENT '性别',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;