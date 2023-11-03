SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';


ALTER SCHEMA `mercadofarma`  DEFAULT CHARACTER SET utf8  DEFAULT COLLATE utf8_general_ci;

CREATE TABLE IF NOT EXISTS `mercadofarma`.`users` (
`user_id` VARCHAR(255) NOT NULL,
`email` VARCHAR(255) NOT NULL,
`first_name` VARCHAR(255) NOT NULL,
`last_name` VARCHAR(255) NULL,
`hash` VARCHAR(255) NOT NULL,
`role` VARCHAR(45) NOT NULL COMMENT 'it can be shopper or admin',
`active` TINYINT(4) NOT NULL DEFAULT 1,
`creation_date` DATETIME NOT NULL DEFAULT NOW(),
`update_date` DATETIME NOT NULL DEFAULT NOW(),
PRIMARY KEY (`user_id`),
UNIQUE INDEX `email_UNIQUE` (`email` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `mercadofarma`.`departments` (
department_id INT(11) NOT NULL AUTO_INCREMENT,
`zip_code` VARCHAR(45) NOT NULL,
`name` VARCHAR(255) NOT NULL,
`active` TINYINT(4) NOT NULL DEFAULT 1,
`creation_date` DATETIME NOT NULL DEFAULT now(),
`update_date` DATETIME NOT NULL DEFAULT NOW(),
PRIMARY KEY (department_id))
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;

insert into mercadofarma.departments(zip_code, name, update_date) value('760001','Valle del cauca', NOW());

CREATE TABLE IF NOT EXISTS `mercadofarma`.`cities` (
`city_id` INT(11) NOT NULL AUTO_INCREMENT,
`zip_code` VARCHAR(45) NULL DEFAULT NULL,
`name` VARCHAR(255) NOT NULL,
`department_id` INT(11) NOT NULL,
`active` TINYINT(4) NOT NULL DEFAULT 1,
`creation_date` DATETIME NOT NULL DEFAULT now(),
`update_date` DATETIME NOT NULL DEFAULT NOW(),
PRIMARY KEY (`city_id`),
INDEX `fk_cities_departments1_idx` (`department_id` ASC),
CONSTRAINT `fk_cities_departments1`
FOREIGN KEY (`department_id`)
   REFERENCES `mercadofarma`.`departments` (department_id)
   ON DELETE NO ACTION
   ON UPDATE NO ACTION)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;

insert into mercadofarma.cities(zip_code, name, department_id, update_date) values('760001', 'cali', 1,  now());

CREATE TABLE IF NOT EXISTS `mercadofarma`.`sectors` (
`sector_id` INT(11) NOT NULL AUTO_INCREMENT,
`name` VARCHAR(255) NOT NULL,
`city_id` INT(11) NOT NULL,
`active` TINYINT(4) NOT NULL DEFAULT 1,
`creation_date` DATETIME NOT NULL DEFAULT now(),
`update_date` DATETIME NOT NULL DEFAULT NOW(),
PRIMARY KEY (`sector_id`),
INDEX `fk_sectors_cities1_idx` (`city_id` ASC),
CONSTRAINT `fk_sectors_cities1`
FOREIGN KEY (`city_id`)
    REFERENCES `mercadofarma`.`cities` (`city_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;

insert into mercadofarma.sectors(name, city_id,update_date) value('Oriente', 1, NOW());

CREATE TABLE IF NOT EXISTS `mercadofarma`.`business` (
`business_id` VARCHAR(255) NOT NULL,
`user_id` VARCHAR(255) NOT NULL,
`sector_id` INT(11) NOT NULL,
`tax_id` VARCHAR(255) NOT NULL,
`company_name` VARCHAR(255) NOT NULL,
`business_type` VARCHAR(255) NOT NULL COMMENT 'Ex: farmacia, tienda/marca cosmetica, ortopedia, optica, perfumeria, otros',
`active` TINYINT(4) NOT NULL DEFAULT 1,
`creation_date` DATETIME NOT NULL DEFAULT now(),
`update_date` DATETIME NOT NULL DEFAULT NOW(),
`address` VARCHAR(255) NOT NULL,
`phone_number` VARCHAR(45) NOT NULL,
PRIMARY KEY (`business_id`),
UNIQUE INDEX `national_id_UNIQUE` (`tax_id` ASC),
INDEX `fk_business_users1_idx` (`user_id` ASC),
INDEX `fk_business_sectors1_idx` (`sector_id` ASC),
CONSTRAINT `fk_business_users1`
 FOREIGN KEY (`user_id`)
     REFERENCES `mercadofarma`.`users` (`user_id`)
     ON DELETE NO ACTION
     ON UPDATE NO ACTION,
CONSTRAINT `fk_business_sectors1`
 FOREIGN KEY (`sector_id`)
     REFERENCES `mercadofarma`.`sectors` (`sector_id`)
     ON DELETE NO ACTION
     ON UPDATE NO ACTION)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;


CREATE TABLE IF NOT EXISTS `mercadofarma`.`categories` (
`category_id` INT(11) NOT NULL AUTO_INCREMENT,
`name` VARCHAR(255) NOT NULL,
`parent_id` VARCHAR(45) NULL DEFAULT NULL,
`active` TINYINT(4) NOT NULL DEFAULT 1,
PRIMARY KEY (`category_id`))
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;


CREATE TABLE IF NOT EXISTS `mercadofarma`.`internal_products` (
`internal_product_id` INT(11) NOT NULL,
`ean` VARCHAR(255) NOT NULL,
`national_code` VARCHAR(45) NULL DEFAULT NULL,
`product_name` VARCHAR(255) NOT NULL,
`description` VARCHAR(255) NOT NULL,
`category_id` INT(11) NOT NULL,
`laboratory` VARCHAR(255) NOT NULL,
`brand` VARCHAR(255) NOT NULL,
`active` TINYINT(4) NOT NULL DEFAULT 1,
PRIMARY KEY (`internal_product_id`),
UNIQUE INDEX `ean_UNIQUE` (`ean` ASC),
INDEX `fk_products_categories1_idx` (`category_id` ASC),
CONSTRAINT `fk_products_categories1`
FOREIGN KEY (`category_id`)
  REFERENCES `mercadofarma`.`categories` (`category_id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `mercadofarma`.`business_products` (
`business_product_id` INT(11) NOT NULL AUTO_INCREMENT,
`business_id` VARCHAR(255) NOT NULL,
`internal_product_id` INT(11) NOT NULL COMMENT 'mercadofarma internal product id',
`national_code` VARCHAR(45) NULL DEFAULT NULL,
`ean` VARCHAR(255) NOT NULL,
`product_name` VARCHAR(255) NOT NULL,
`description` VARCHAR(255) NOT NULL,
`price` FLOAT(11) NOT NULL DEFAULT 0,
`vat` FLOAT(11) NOT NULL DEFAULT 0 COMMENT 'iva',
`stock` INT(11) NOT NULL DEFAULT 0,
`active` TINYINT(4) NOT NULL DEFAULT 1,
PRIMARY KEY (`business_product_id`),
INDEX `fk_products_business1_idx` (`business_id` ASC),
INDEX `fk_business_products_products1_idx` (`internal_product_id` ASC),
CONSTRAINT `fk_products_business1`
FOREIGN KEY (`business_id`)
  REFERENCES `mercadofarma`.`business` (`business_id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION,
CONSTRAINT `fk_business_products_products1`
FOREIGN KEY (`internal_product_id`)
  REFERENCES `mercadofarma`.`internal_products` (`internal_product_id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;


CREATE TABLE IF NOT EXISTS `mercadofarma`.`images` (
`image_id` INT(11) NOT NULL AUTO_INCREMENT,
`product_id` INT(11) NOT NULL,
`size` VARCHAR(45) NOT NULL COMMENT 'Can be 300x300 or 150x150 etc',
`path` VARCHAR(255) NOT NULL,
`creation_date` DATETIME NOT NULL DEFAULT now(),
`update_date` DATETIME NOT NULL DEFAULT NOW(),
PRIMARY KEY (`image_id`),
INDEX `fk_images_products1_idx` (`product_id` ASC),
CONSTRAINT `fk_images_products1`
   FOREIGN KEY (`product_id`)
       REFERENCES `mercadofarma`.`internal_products` (`internal_product_id`)
       ON DELETE NO ACTION
       ON UPDATE NO ACTION)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `mercadofarma`.`business_opening_hours` (
`business_opening_hour_id` INT(11) NOT NULL AUTO_INCREMENT,
`business_id` VARCHAR(255) NOT NULL,
`day` VARCHAR(255) NOT NULL,
`start_time` VARCHAR(255) NOT NULL,
`ending_time` VARCHAR(255) NOT NULL,
PRIMARY KEY (`business_opening_hour_id`),
INDEX `fk_business_opening_hours_business_idx` (`business_id` ASC),
CONSTRAINT `fk_business_opening_hours_business1`
   FOREIGN KEY (`business_id`)
       REFERENCES `mercadofarma`.`business` (`business_id`)
       ON DELETE NO ACTION
       ON UPDATE NO ACTION)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `mercadofarma`.`payment_methods` (
`payment_method_id` INT(11) NOT NULL AUTO_INCREMENT,
`name` VARCHAR(255) NOT NULL,
`creation_date` DATETIME NOT NULL DEFAULT NOW(),
`update_date` DATETIME NOT NULL DEFAULT NOW(),
PRIMARY KEY (`payment_method_id`))
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `mercadofarma`.`carts` (
`cart_id` INT(11) NOT NULL,
`user_id` VARCHAR(255) NULL DEFAULT NULL,
`delivery_type` VARCHAR(255) NULL DEFAULT NULL COMMENT 'address_delivery, picking',
`full_name` VARCHAR(255) NULL DEFAULT NULL,
`email` VARCHAR(255) NULL DEFAULT NULL,
`phone_number` VARCHAR(45) NULL DEFAULT NULL,
`address` VARCHAR(255) NULL DEFAULT NULL,
`payment_method_id` INT(11) NULL DEFAULT NULL,
`shipping_cost` FLOAT(11) NOT NULL DEFAULT 0,
`total` FLOAT(11) NOT NULL DEFAULT 0,
`saving` FLOAT(11) NOT NULL DEFAULT 0,
`creation_date` DATETIME NOT NULL DEFAULT now(),
PRIMARY KEY (`cart_id`),
INDEX `fk_shopping_cart_users1_idx` (`user_id` ASC),
INDEX `fk_carts_payment_methods1_idx` (`payment_method_id` ASC),
CONSTRAINT `fk_shopping_cart_users1`
FOREIGN KEY (`user_id`)
  REFERENCES `mercadofarma`.`users` (`user_id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION,
CONSTRAINT `fk_carts_payment_methods1`
FOREIGN KEY (`payment_method_id`)
  REFERENCES `mercadofarma`.`payment_methods` (`payment_method_id`)
  ON DELETE NO ACTION
  ON UPDATE NO ACTION)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;

CREATE TABLE IF NOT EXISTS `mercadofarma`.`cart_details` (
`cart_detail_id` INT(11) NOT NULL AUTO_INCREMENT,
`cart_id` INT(11) NOT NULL,
`business_product_id` INT(11) NOT NULL,
`quantity` INT(11) NOT NULL DEFAULT 1,
`price` FLOAT(11) NOT NULL DEFAULT 0,
`vat` INT(11) NOT NULL DEFAULT 0,
`total_price` FLOAT(11) NOT NULL DEFAULT 0,
PRIMARY KEY (`cart_detail_id`),
INDEX `fk_shopping_cart_details_shopping_cart1_idx` (`cart_id` ASC),
INDEX `fk_shopping_cart_details_business_products1_idx` (`business_product_id` ASC),
CONSTRAINT `fk_shopping_cart_details_shopping_cart1`
 FOREIGN KEY (`cart_id`)
     REFERENCES `mercadofarma`.`carts` (`cart_id`)
     ON DELETE NO ACTION
     ON UPDATE NO ACTION,
CONSTRAINT `fk_shopping_cart_details_business_products1`
 FOREIGN KEY (`business_product_id`)
     REFERENCES `mercadofarma`.`business_products` (`business_product_id`)
     ON DELETE NO ACTION
     ON UPDATE NO ACTION)
    ENGINE = InnoDB
    DEFAULT CHARACTER SET = utf8;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
