Use RestaurantDB;

CREATE TABLE `Restaurant` (
    `rest_id` INT NOT NULL AUTO_INCREMENT ,
    `name` VARCHAR(200) NULL DEFAULT NULL , 
    `rating` FLOAT NULL DEFAULT NULL , 
    `cuisines` VARCHAR(500) NULL DEFAULT NULL , 
    `address` VARCHAR(500) NULL DEFAULT NULL , 
    `timing` VARCHAR(200) NULL DEFAULT NULL , 
    `cft` INT NULL DEFAULT NULL , 
    PRIMARY KEY (`rest_id`)
);