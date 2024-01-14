package tickers

type TickerResponse struct {
	Code string `json:"code"`
	Data []struct {
		AskPr        string `json:"askPr"`
		AskSz        string `json:"askSz"`
		BaseVolume   string `json:"baseVolume"`
		BidPr        string `json:"bidPr"`
		BidSz        string `json:"bidSz"`
		Change24H    string `json:"change24h"`
		ChangeUtc24H string `json:"changeUtc24h"`
		High24H      string `json:"high24h"`
		LastPr       string `json:"lastPr"`
		Low24H       string `json:"low24h"`
		Open         string `json:"open"`
		OpenUtc      string `json:"openUtc"`
		QuoteVolume  string `json:"quoteVolume"`
		Symbol       string `json:"symbol"`
		Ts           string `json:"ts"`
		UsdtVolume   string `json:"usdtVolume"`
	} `json:"data"`
	Msg string `json:"msg"`
}
