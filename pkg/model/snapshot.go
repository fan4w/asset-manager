package model

// Snapshot 对应单个 JSON 文件的完整结构
type Snapshot struct {
	Meta          Meta               `json:"meta"`
	ExchangeRates map[string]float64 `json:"exchange_rates"`
	Accounts      []Account          `json:"accounts"`
}

type Meta struct {
	SnapshotID   string `json:"snapshot_id"`
	Period       string `json:"period"` // e.g. "2025-11"
	Type         string `json:"type"`   // OPENING, CLOSING
	BaseCurrency string `json:"base_currency"`
}

type Account struct {
	ID       string  `json:"account_id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}
