package model

// currencyBreakdown 货币拆分结构体
type currencyBreakdown struct {
	Amount         float64 // 原币金额
	ToBaseCurrency float64 // 折算成基准货币金额
}

type categorySummary struct {
	ByCurrency        map[string]currencyBreakdown // 根据货币的拆分
	TotalBaseCurrency float64                      // 该类目折算成基准货币的金额
}

type AssetSummary struct {
	Period            string
	TotalBaseCurrency float64
	ByCurrency        map[string]currencyBreakdown // 全局按货币统计
	ByCategory        map[string]categorySummary   // 按 category -> currency 统计
}

// NewAssetSummary 构造器
func NewAssetSummary(period string) AssetSummary {
	return AssetSummary{
		Period:     period,
		ByCurrency: make(map[string]currencyBreakdown),
		ByCategory: make(map[string]categorySummary),
	}
}

// Add 用于将单个账户（或一笔资产）按 category + currency 累加进 Summary
func (s *AssetSummary) Add(category, currency string, amount, rate float64) {
	if amount == 0 || currency == "" {
		return
	}
	toBaseCurrency := amount * rate
	s.addCurrency(currency, amount, toBaseCurrency)
	
	// 按 category 的拆分
	cs := s.ByCategory[category]
	if cs.ByCurrency == nil {
		cs.ByCurrency = make(map[string]currencyBreakdown)
	}
	// 当前账户对应的 category 的 currency
	cb := cs.ByCurrency[currency]
	cb.Amount += amount
	cb.ToBaseCurrency += toBaseCurrency
	cs.ByCurrency[currency] = cb
	cs.TotalBaseCurrency += toBaseCurrency
	s.ByCategory[category] = cs
}

// 按货币更新统计
func (s *AssetSummary) addCurrency(currency string, amount, toBaseCurrency float64) {
	cb := s.ByCurrency[currency]
	cb.Amount += amount
	cb.ToBaseCurrency += toBaseCurrency
	s.ByCurrency[currency] = cb
	s.TotalBaseCurrency += toBaseCurrency
}
