package show

import (
	"asset-manager/pkg/analysis"
	"asset-manager/pkg/model"
	"fmt"
	"strings"
)

func PrintWealthReport(sums []model.AssetSummary) {
	fmt.Println("\n=== 💰 资产快照表 ===")
	fmt.Printf("%-10s \t|   %-12s\t| %s\n", "月份", "总估值", "资产构成")
	fmt.Println(strings.Repeat("-", 60))

	for _, s := range sums {
		// 简单格式化各币种构成
		var details []string
		for k, v := range s.ByCurrency {
			if v.ToBaseCurrency > 0 {
				details = append(details, fmt.Sprintf("%s:%.2f", k, v.ToBaseCurrency))
			}
		}
		fmt.Printf("%-10s\t| %10.2f\t\t| %v\n", s.Period, s.TotalBaseCurrency, details)
	}
}

func PrintReportByMonths(sums []model.AssetSummary) {
	if len(sums) < 2 {
		return
	}
	fmt.Println("\n=== 📊 趋势分析 (Month over Month) ===")
	for i := 1; i < len(sums); i++ {
		prev := sums[i-1]
		curr := sums[i]
		diff := analysis.CompareMonthsInBaseCurrency(prev, curr)
		trend := "持平"
		if diff > 0 {
			trend = "增长 📈"
		} else if diff < 0 {
			trend = "减少 📉"
		}

		fmt.Printf("%s -> %s: 变动 %+10.2f (%s)\n",
			prev.Period, curr.Period, diff, trend)
	}
}

func PrintTrendReportByCategory(sums []model.AssetSummary) {
	if len(sums) < 2 {
		return
	}
	fmt.Println("\n=== 📊 趋势分析 by Category (Month over Month) ===")
	for i := 1; i < len(sums); i++ {
		prev := sums[i-1]
		curr := sums[i]
		diff := analysis.CompareMonthsByCategory(prev, curr)
		fmt.Printf("%s -> %s:\n", prev.Period, curr.Period)
		for category, change := range diff {
			trend := "持平"
			if change > 0 {
				trend = "增长 📈"
			} else if change < 0 {
				trend = "减少 📉"
			}
			fmt.Printf("  %s: %+10.2f (%s)\n", category, change, trend)
		}
	}
}

func PrintDetailedSnapshot(snap *model.Snapshot) {
	fmt.Printf("\n=== 📝 详细快照 (%s - %s) ===\n", snap.Meta.Period, snap.Meta.Type)
	fmt.Printf("快照ID: %s\n", snap.Meta.SnapshotID)
	fmt.Println("账户列表:")

	// 遍历打印
	for _, acc := range snap.Accounts {
		// 使用 \t 分割列
		// 格式: - [ID] 名称 \t | 类别: CATEGORY \t | 余额: ...
		fmt.Printf("- [%s] %s\t| 类别: %s\t| 余额: %.2f %s\n",
			acc.ID, acc.Name, acc.Category, acc.Balance, acc.Currency)
	}
}

func PrintAnalyzedSnapshot(snap *model.Snapshot) {
	asset := analysis.CalculateNetWorth(snap)
	fmt.Printf("\n=== 📝 按类别统计 (%s - %s) ===\n", snap.Meta.Period, snap.Meta.Type)
	for k, v := range asset.ByCategory {
		if v.TotalBaseCurrency > 0 {
			sum := 0.0
			for curr, amount := range v.ByCurrency {
				fmt.Printf("类别: %s | 币种: %s | 余额: %.2f | 折合%s: %.2f \n",
					k, curr, amount.Amount, snap.Meta.BaseCurrency, amount.ToBaseCurrency)
				sum += amount.ToBaseCurrency
			}
			fmt.Printf("%s 共计 %.2f %s\n", k, sum, snap.Meta.BaseCurrency)
		}
	}
}
