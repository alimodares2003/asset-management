package service

import (
	"assets-management/app/model"
	"encoding/json"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/unicode/bidi"
	"io"
	"net/http"
	"strconv"
)

const Mesghal = "gold-miskal"
const NimiCoin = "coin-baharazadi-nim"
const RobiCoin = "coin-baharazadi-rob"
const AzadiCoin = "coin-baharazadi"
const EmamiCoin = "coin-emami"
const GeramiCoin = "coin-gerami"
const Gold18 = "gold-gram-18"
const GheyreCoin = "coin-gheyre"

func GetAssetsTotal(request []model.AssetRequest) model.AssetResponse {
	var assets []model.Asset
	var total int64 = 0
	for i, assetRequest := range request {
		var assetType = assetRequest.AssetType
		if assetType == Gold18 || assetType == GheyreCoin {
			assetType = Mesghal
		}
		var rawPrice = getGoldPrices(assetType).Price
		fmt.Println(strconv.Itoa(i) + ": prices get " + strconv.FormatInt(rawPrice, 10))
		if assetRequest.AssetType == Gold18 {
			rawPrice = int64(float64(rawPrice) / 4.331802)
		}
		if assetRequest.AssetType == GheyreCoin {
			var gramPrice = float64(rawPrice) / 4.331802
			rawPrice = int64((8.13 / 750) * 900 * gramPrice)
		}

		var totalPrice = rawPrice * int64(assetRequest.AssetCount)
		total += totalPrice
		assets = append(assets, model.Asset{TotalPrice: displayPrice(rialToToman(totalPrice)), RawPrice: displayPrice(rialToToman(rawPrice)), Type: assetRequest.AssetType, Count: assetRequest.AssetCount})
	}
	return model.AssetResponse{Assets: assets, Total: displayPrice(rialToToman(total))}
}

func getGoldPrices(assetType string) model.AssetPrice {
	priceResponse, err := http.Get(fmt.Sprintf("https://api.priceto.day/v1/latest/irr/%s", assetType))
	if err != nil {
		fmt.Println(err.Error())
		return model.AssetPrice{}
	}

	var assetPrice model.AssetPrice
	jsonBytes, err := io.ReadAll(priceResponse.Body)
	if err != nil {
		return model.AssetPrice{}
	}
	if err != nil {
		fmt.Println("Read Bytes Error: " + err.Error())
		return model.AssetPrice{}
	}
	err = json.Unmarshal(jsonBytes, &assetPrice)
	if err != nil {
		fmt.Println("Json Parse Error: " + err.Error())
		return model.AssetPrice{}
	}

	return assetPrice
}

func rialToToman(rial int64) int64 {
	return rial / 10
}

func displayPriceV2(price int64) string {
	return message.NewPrinter(language.English).Sprint(price)
}

func displayPrice(price int64) string {
	var strPrice = strconv.FormatInt(price, 10)
	println(strPrice)
	var displayPriceArr = make([]byte, 0)
	var res = ""
	for i := len(strPrice) - 1; i >= 0; i-- {
		if len(displayPriceArr) != 0 && len(displayPriceArr)%3 == 0 {
			displayPriceArr = append(displayPriceArr, strPrice[i])
			res += ","
			res += string(strPrice[i])
		} else {
			displayPriceArr = append(displayPriceArr, strPrice[i])
			res += string(strPrice[i])
		}
	}
	return bidi.ReverseString(res)
}
