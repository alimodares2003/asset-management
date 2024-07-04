package model

//type AssetRequest struct {
//	TamomiCount int32 `json:"tamomi_count"`
//	NimiCount   int32 `json:"nimi_count"`
//	RobiCount   int32 `json:"robi_count"`
//	GoldGram    int32 `json:"gold_gram"`
//}

type AssetRequest struct {
	AssetType  string `json:"asset_type"`
	AssetCount int32  `json:"asset_count"`
}

type Asset struct {
	Count      int32  `json:"count"`
	Type       string `json:"type"`
	RawPrice   string `json:"raw_price"`
	TotalPrice string `json:"total_price"`
}
type AssetResponse struct {
	Assets []Asset `json:"assets"`
	Total  string  `json:"total"`
}

type AssetPrice struct {
	Success bool   `json:"success"`
	Price   int64  `json:"price"`
	High    int64  `json:"high"`
	Low     int64  `json:"low"`
	Time    string `json:"time"`
}
