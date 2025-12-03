package analysis

import (
	"asset-manager/pkg/model"
)

// CalculateNetWorth 计算单月净值
func CalculateNetWorth(snap *model.Snapshot) model.AssetSummary {
	s := model.NewAssetSummary(snap.Meta.Period)
	for _, acc := range snap.Accounts {
		rate := snap.ExchangeRates[acc.Currency]
		s.Add(acc.Category, acc.Currency, acc.Balance, rate)
	}

	return s
}
