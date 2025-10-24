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

CREATE TABLE employees (
   `id` INT PRIMARY KEY AUTO_INCREMENT,
   `name` varchar(63) NOT NULL COMMENT '员工名字',
   `department` varchar(63) NOT NULL COMMENT '部门名称',
   `salary` DECIMAL(10, 2) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE books (
   id INT PRIMARY KEY AUTO_INCREMENT,
   title VARCHAR(127) NOT NULL,  -- 书名（字符串类型，对应结构体 string）
   author VARCHAR(63) NOT NULL, -- 作者（字符串类型，对应结构体 string）
   price DECIMAL(10, 2) NOT NULL -- 价格（ decimal 类型，对应结构体 float64）
) ENGINE=InnoDB;