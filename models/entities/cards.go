package entities

type Card struct {
	Base
	Ranking  int64  `gorm:"column:ranking" json:"-"`
	Value    string `gorm:"column:value" json:"value"`
	Suit     string `gorm:"column:suit" json:"suit"`
	Code     string `gorm:"column:code" json:"code"`
	SuitType string `gorm:"column:suit_type" json:"-"`
}

func (i Card) TableName() string {
	return "cards"
}
