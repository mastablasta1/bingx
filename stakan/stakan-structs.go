package stakan

type StakanDataStruct struct {
	Code      int   `json:"code"`
	Timestamp int64 `json:"timestamp"`
	Data      struct {
		Bids [][]string `json:"bids"`
		Asks [][]string `json:"asks"`
		Ts   int64      `json:"ts"`
	} `json:"data"`
}

type TikersDataStruct struct {
	Code      int   `json:"code"`
	Timestamp int64 `json:"timestamp"`
	Data      []struct {
		Symbol      string  `json:"symbol"`
		OpenPrice   float64 `json:"openPrice"`
		HighPrice   float64 `json:"highPrice"`
		LowPrice    float64 `json:"lowPrice"`
		LastPrice   float64 `json:"lastPrice"`
		Volume      float64 `json:"volume"`
		QuoteVolume float64 `json:"quoteVolume"`
		OpenTime    int64   `json:"openTime"`
		CloseTime   int64   `json:"closeTime"`
		AskPrice    float64 `json:"askPrice"`
		AskQty      float64 `json:"askQty"`
		BidPrice    float64 `json:"bidPrice"`
		BidQty      float64 `json:"bidQty"`
	} `json:"data"`
}
