-- Create table decks
CREATE TABLE `cards` (
  `id` int NOT NULL AUTO_INCREMENT,
  `ranking` int NOT NULL,
  `value` varchar(50) NOT NULL,
  `suit` varchar(50) NOT NULL,
  `code` varchar(10) NOT NULL,
  `suit_type` varchar(30) NOT NULL DEFAULT 'french_suit',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `card_code_idx` (`code`,`status`)
);