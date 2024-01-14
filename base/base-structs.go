package base

// Структура таблицы
type Bingxs struct {
	Symbol         string  `json:"symbol"`
	Volume         float64 `json:"volume"`
	QuoteVolume    float64 `json:"quote_volume"`
	UpdatedAt      int64   `json:"updated_at"`
	AskOne         float64 `json:"ask_one"`
	AskDuo         float64 `json:"ask_duo"`
	BidOne         float64 `json:"bid_one"`
	BidDuo         float64 `json:"bid_duo"`
	Raznitca       float64 `json:"raznitca"`
	TradesCountOld int     `json:"trades_count_old"`
	TradesCountNew int     `json:"trades_count_new"`
	Active         int     `json:"active"`
}
