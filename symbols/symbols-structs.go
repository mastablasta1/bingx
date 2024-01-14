package symbols

type SymbolsResponse struct {
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
	DebugMsg string `json:"debugMsg"`
	Data     struct {
		Symbols []struct {
			Symbol      string  `json:"symbol"`
			MinQty      float64 `json:"minQty"`
			MaxQty      float64 `json:"maxQty"`
			MinNotional int     `json:"minNotional"`
			MaxNotional int     `json:"maxNotional"`
			Status      int     `json:"status"`
			TickSize    float64 `json:"tickSize"`
			StepSize    float64 `json:"stepSize"`
		} `json:"symbols"`
	} `json:"data"`
}
