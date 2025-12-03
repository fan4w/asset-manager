package repository

import (
	"asset-manager/pkg/model"
	"asset-manager/pkg/validator"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

type JSONFileRepository struct {
	RootDir string
}

func NewJSONFileRepository(root string) *JSONFileRepository {
	return &JSONFileRepository{RootDir: root}
}

func (r *JSONFileRepository) LoadAllSnapshots() ([]*model.Snapshot, error) {
	var snapshots []*model.Snapshot

	entries, err := os.ReadDir(r.RootDir)
	if err != nil {
		return nil, fmt.Errorf("读取目录失败: %w", err)
	}

	// 过滤并排序目录
	var monthDirs []string
	datePattern := regexp.MustCompile(`^\d{4}-\d{2}$`)
	for _, entry := range entries {
		if entry.IsDir() && datePattern.MatchString(entry.Name()) {
			monthDirs = append(monthDirs, entry.Name())
		}
	}
	sort.Strings(monthDirs)

	for _, month := range monthDirs {
		filePath := filepath.Join(r.RootDir, month, "opening.json")
		snap, err := r.parseFile(filePath)
		if err != nil {
			slog.Warn(err.Error())
			continue
		}
		err = validator.ValidateJsonSnapshot(snap, month)
		if err != nil {
			slog.Error(err.Error())
			continue
		}
		snapshots = append(snapshots, snap)
	}

	return snapshots, nil
}

func (r *JSONFileRepository) parseFile(path string) (*model.Snapshot, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var snap model.Snapshot
	if err := json.Unmarshal(content, &snap); err != nil {
		return nil, err
	}
	return &snap, nil
}

func (r *JSONFileRepository) GetByPeriod(period string) (*model.Snapshot, error) {
	filePath := filepath.Join(r.RootDir, period, "opening.json")
	return r.parseFile(filePath)
}
