package main

import (
	"time"

	/*
		Парсер API биржи bitget.com

		https://www.bitget.com/api-doc/common/release-note

		https://bingx-api.github.io/docs/#/en-us/swapV2/base-info.html#Rate%20limit

		https://open-api.bingx.com/openApi/spot/v1/common/symbols
		https://open-api.bingx.com/openApi/spot/v1/ticker/24hr?symbol=WIFI-USDT&timestamp=1701263935139
		https://open-api.bingx.com/openApi/spot/v1/market/depth?symbol=WIFI-USDT

	*/

	sym "bingxstakan/symbols"

	sta "bingxstakan/stakan"
)

func main() {

	//	Текущие данные по таймстампам
	var timeFive int64 = 0

	for {

		//	ОБновим данные по парам
		timeNow := time.Now().Unix()

		if timeNow > timeFive {

			timeFive = timeNow + 3600 //	Задержка на 5 минут

			//	Класс обновления данных по парам
			sym.GetSymbolsUpdate(timeNow)

		}

		//	Обновим данные по данным в парах
		sta.StartParser(timeNow)

		//	Тормознем на секунду
		time.Sleep(500 * time.Millisecond)
		//time.Sleep(1 * time.Second)

	}

}
