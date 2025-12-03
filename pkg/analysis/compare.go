package analysis

import (
	"asset-manager/pkg/model"
	"fmt"
)

func CompareMonths(prev, curr model.AssetSummary) int {
	return 0
}

// CompareMonthsInBaseCurrency 对比两个月的变化 (支持未来扩展：收入/支出分析)
func CompareMonthsInBaseCurrency(prev, curr model.AssetSummary) string {
	diff := curr.TotalBaseCurrency - prev.TotalBaseCurrency

	trend := "持平"
	if diff > 0 {
		trend = "增长 📈"
	} else if diff < 0 {
		trend = "减少 📉"
	}

	return fmt.Sprintf("%s -> %s: 变动 %+10.2f (%s)",
		prev.Period, curr.Period, diff, trend)
}
