package static

import (
	models "tarot-app/models"
)

// InitializeTarotSpreads returns a list of predefined tarot spreads
func InitializeTarotSpreads() []models.TarotSpread {
	return []models.TarotSpread{
		{
			Name:        "One-Card Spread",
			Description: "适合回答简单的是非问题或提供日常指引",
			CardsNeeded: 1,
		},
		{
			Name:        "Three-Card Spread",
			Description: "代表过去、现在和未来，或情况、障碍和建议",
			CardsNeeded: 3,
		},
		{
			Name:        "Celtic Cross Spread",
			Description: "较为复杂，适合深入探讨一个问题，包括现状、挑战、潜在基础、过去、未来等方面",
			CardsNeeded: 10,
		},
		{
			Name:        "Time Flow Spread",
			Description: "适合预测未来，按照时间顺序排列牌阵，以了解事件的发展过程",
			CardsNeeded: 5,
		},
		{
			Name:        "Holy Triangle Spread",
			Description: "适合判断情势和寻找成因，通过牌阵的前因后果来分析问题",
			CardsNeeded: 3,
		},
		{
			Name:        "Hexagram Spread",
			Description: "用于占卜事物发展和预测未来，理清事情的本源",
			CardsNeeded: 6,
		},
		{
			Name:        "Four Elements Spread",
			Description: "通过四张牌来代表火、水、风、土四个元素，分析问题的各个方面",
			CardsNeeded: 4,
		},
		{
			Name:        "Tree of Life Spread",
			Description: "使用11张牌，代表卡巴拉生命之树的各个节点，深入分析问题",
			CardsNeeded: 11,
		},
		{
			Name:        "Golden Zodiac Spread",
			Description: "根据黄道十二宫的布局来排列牌阵，探索与星座相关的主题",
			CardsNeeded: 12,
		},
		{
			Name:        "Celtic Cross Spread",
			Description: "一种古老的塔罗牌阵，由7张牌组成，用于深入分析问题",
			CardsNeeded: 7,
		},
		{
			Name:        "Venus Love Spread",
			Description: "专门用于分析爱情和关系问题，了解双方的感受和关系的走向",
			CardsNeeded: 5,
		},
	}
}
