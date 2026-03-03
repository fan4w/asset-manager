package analysis

import (
	"asset-manager/pkg/model"
)

func CompareMonths(prev, curr model.AssetSummary) int {
	return 0
}

// CompareMonthsInBaseCurrency 对比两个月的变化 (支持未来扩展：收入/支出分析)
func CompareMonthsInBaseCurrency(prev, curr model.AssetSummary) float64 {
	return curr.TotalBaseCurrency - prev.TotalBaseCurrency
}

func CompareMonthsByCategory(prev, curr model.AssetSummary) map[string]float64 {
	diff := make(map[string]float64)
	for category, currData := range curr.ByCategory {
		prevData, ok := prev.ByCategory[category]
		if ok {
			diff[category] = currData.TotalBaseCurrency - prevData.TotalBaseCurrency
		} else {
			diff[category] = currData.TotalBaseCurrency
		}
	}
	for category, prevData := range prev.ByCategory {
		if _, ok := curr.ByCategory[category]; !ok {
			diff[category] = -prevData.TotalBaseCurrency
		}
	}
	return diff
}

func CompareMonthsByCurrencyAndCategory(prev, curr model.AssetSummary) map[string]map[string]float64 {
	diff := make(map[string]map[string]float64)
	for category, currData := range curr.ByCategory {
		if _, ok := diff[category]; !ok {
			diff[category] = make(map[string]float64)
		}
		for currency, currCurData := range currData.ByCurrency {
			prevCurData, ok := prev.ByCategory[category].ByCurrency[currency]
			if ok {
				diff[category][currency] = currCurData.ToBaseCurrency - prevCurData.ToBaseCurrency
			} else {
				diff[category][currency] = currCurData.ToBaseCurrency
			}
		}
	}
	for category, prevData := range prev.ByCategory {
		if _, ok := diff[category]; !ok {
			diff[category] = make(map[string]float64)
		}
		for currency, prevCurData := range prevData.ByCurrency {
			if _, ok := curr.ByCategory[category].ByCurrency[currency]; !ok {
				diff[category][currency] = -prevCurData.ToBaseCurrency
			}
		}
	}
	return diff
}
