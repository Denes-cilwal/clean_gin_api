-- +migrate Up
CREATE TABLE IF NOT EXISTS `sms` (
  `id` BINARY(16) NOT NULL, 
  `message` VARCHAR(100) NOT NULL,
  `to` VARCHAR(20) NOT NULL,
  `created_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE IF EXISTS `sms`;