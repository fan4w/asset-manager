package validator

import (
	"asset-manager/pkg/model"
	"fmt"
)

// ValidateSnapshot 执行所有针对快照的检查
func ValidateJsonSnapshot(snap *model.Snapshot, folderName string) error {
	// 防止复制粘贴了文件，却忘了改 meta.period
	if snap.Meta.Period != folderName {
		return fmt.Errorf("元数据不匹配: 文件夹名为 '%s', 但 JSON 内部 period 为 '%s'",
			folderName, snap.Meta.Period)
	}

	seenIDs := make(map[string]bool)

	for _, acc := range snap.Accounts {
		// 检查 ID 重复
		if acc.Name != "" {
			if seenIDs[acc.Name] {
				return fmt.Errorf("记录 '%s' 中发现重复的账户 ID: %s", snap.Meta.Period, acc.Name)
			}
			seenIDs[acc.Name] = true
		}

		// 跳过注释行或未激活账户
		if acc.Currency == "" || acc.Balance == 0 {
			continue
		}

		// 检查汇率缺失
		if _, ok := snap.ExchangeRates[acc.Currency]; !ok {
			return fmt.Errorf("记录 '%s' 中账户 '%s' 使用了货币 '%s', 但未提供该货币的汇率",
				snap.Meta.Period, acc.Name, acc.Currency)
		}
	}
	return nil
}
