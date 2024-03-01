-- Create database on MySQL
CREATE DATABASE IF NOT EXISTS `test_data`
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

USE test_data;

-- Create table on MySQL storing people's information. 
-- name with string of max 50 characters.
-- age with integer.
-- state with string of max 2 characters.
CREATE TABLE IF NOT EXISTS `people` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(50) NOT NULL,
    `age` INT,
    `state` CHAR(2),
    -- created time with datetime type
    `created_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    -- updated time with datetime type
    `updated_time` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

-- Insert some data.
INSERT INTO `people` (`name`, `age`, `state`) VALUES ('John Doe', 30, 'NY');
INSERT INTO `people` (`name`, `age`, `state`) VALUES ('Tim Armstrong', 25, 'CA');
INSERT INTO `people` (`name`, `age`, `state`) VALUES ('Ian MacKaye', 40, 'DC');
INSERT INTO `people` (`name`, `age`, `state`) VALUES ('Billy Armstrong', 22, 'CA');
INSERT INTO `people` (`name`, `age`, `state`) VALUES ('Jeff Tweedy', 34, 'IL');
INSERT INTO `people` (`name`, `age`, `state`) VALUES ('Ben Gibbard', 33, 'WA');