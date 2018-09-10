package models

type SummarySales struct {
	Omzet         float64 `json:"omzet"`
	GrossProfit   float64 `json:"grossProfit"`
	TotalItemSold int64   `json:"totalItemSold"`
	Records       []Sales `json:"records"`
}
