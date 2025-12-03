package repository

import "asset-manager/pkg/model"

// AssetRepository 定义了数据访问的标准行为
// 任何想要提供数据的源都必须实现这个接口
type AssetRepository interface {
	// LoadAllSnapshots 加载所有历史快照
	LoadAllSnapshots() ([]*model.Snapshot, error)
	// GetByPeriod 根据月份加载特定快照
	GetByPeriod(period string) (*model.Snapshot, error)
}
