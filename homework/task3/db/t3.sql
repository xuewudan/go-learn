--学生表
CREATE TABLE `students` (
    `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
    `name` varchar(63) NOT NULL COMMENT '姓名',
    `age` int NOT NULL COMMENT '年龄',
    `gender` int NOT NULL COMMENT '性别',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 账户表
CREATE TABLE accounts (
      id INT PRIMARY KEY AUTO_INCREMENT,
      balance DECIMAL(10, 2) NOT NULL DEFAULT 0.00
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- 交易记录表
CREATE TABLE transactions (
      id INT PRIMARY KEY AUTO_INCREMENT,
      from_account_id INT NOT NULL,
      to_account_id INT NOT NULL,
      amount DECIMAL(10, 2) NOT NULL,
      FOREIGN KEY (from_account_id) REFERENCES accounts(id),
      FOREIGN KEY (to_account_id) REFERENCES accounts(id)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;