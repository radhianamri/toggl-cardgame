-- Create table decks
CREATE TABLE `decks` (
  `id` int NOT NULL AUTO_INCREMENT,
  `external_id` varchar(40) NOT NULL,
  `shuffled` tinyint NOT NULL DEFAULT '0',
  `status` tinyint NOT NULL DEFAULT '1',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `deck_ext_idx` (`external_id`,`status`)
);