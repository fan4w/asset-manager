package main

import (
	"asset-manager/pkg/analysis"
	"asset-manager/pkg/model"
	"asset-manager/pkg/repository"
	"asset-manager/pkg/show"
	"fmt"
	"log/slog"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelInfo)
	var repo repository.AssetRepository = repository.NewJSONFileRepository("./data")

	snapshots, err := repo.LoadAllSnapshots()
	if err != nil {
		panic(err)
	}

	if len(snapshots) == 0 {
		slog.Error("未找到任何资产快照数据。")
		return
	}

	// 执行分析
	var summaries []model.AssetSummary
	for _, snap := range snapshots {
		summaries = append(summaries, analysis.CalculateNetWorth(snap))
	}

	// 展示报告
	show.PrintWealthReport(summaries)
	show.PrintTrendReport(summaries)

	snapshot, err := repo.GetByPeriod("2025-11")
	if err != nil {
		slog.Error("获取特定月份快照失败:", "error", err)
	}

	show.PrintAnalyzedSnapshot(snapshot)
	show.PrintDetailedSnapshot(snapshot)
	fmt.Printf("当前基础币种：%s\n", snapshot.Meta.BaseCurrency)
}
