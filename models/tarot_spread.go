package models

// TarotSpread represents the structure of a tarot card spread
type TarotSpread struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"` // 自增主键
	Name        string // 排阵的名称
	Description string // 描述
	CardsNeeded int    // 需要的牌数
}
