ALTER TABLE `deck_cards`
ADD CONSTRAINT fk_deckcards_deck
FOREIGN KEY (`deck_id`) REFERENCES `decks`(`id`);

ALTER TABLE `deck_cards`
ADD CONSTRAINT fk_deckcards_card
FOREIGN KEY (`card_id`) REFERENCES `cards`(`id`);