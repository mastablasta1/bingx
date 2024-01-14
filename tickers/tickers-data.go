package tickers

/*
https://www.bitget.com/api-doc/spot/market/Get-Tickers
*/
import (
	b "bingxstakan/base"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func GetTickersUpdate(timeNow int64) {

	// Создаем HTTP-клиента
	client := &http.Client{}

	// Создаем GET-запрос
	req, err := http.NewRequest("GET", "https://api.bitget.com/api/v2/spot/market/tickers", nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}

	// Выполняем запрос и получаем ответ
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	// Распаковываем JSON-ответ в структуру TickerResponse
	var tickerResp TickerResponse
	err = json.Unmarshal(body, &tickerResp)
	if err != nil {
		fmt.Println("Ошибка при парсинге JSON:", err)
		return
	}

	// Обрабатываем результат
	strTime := strconv.FormatInt(timeNow, 10)
	var megaUpdateArray []string

	if tickerResp.Code == "00000" {
		for _, ticker := range tickerResp.Data {
			// Выводим информацию о тикере
			/*
				fmt.Println("Символ:", ticker.Symbol)
				fmt.Println("Цена открытия:", ticker.OpenPrice)
				fmt.Println("Цена открытия:", ticker.Volume24h)
				fmt.Println("-------------------------")
			*/

			megaUpdateArray = append(megaUpdateArray, "('"+ticker.Symbol+"','"+strTime+"','"+ticker.UsdtVolume+"','1')")
		}
	} else {
		fmt.Println("Ошибка:", tickerResp.Code)
	}

	megaUpdateString := strings.Join(megaUpdateArray, ",")

	//	Обновим данные
	updateTickers(megaUpdateString)

	//	Почистим после себя
	megaUpdateString = ""
	megaUpdateArray = []string{}

	//fmt.Println("Размер массива: ", len(megaUpdateArray))

}

// Апдейтим данные по тикерам - по обхему торгов
func updateTickers(megaUpdateString string) {

	//	Занесем свежие данные и включим символы, где есть данные
	query2 := "REPLACE INTO `bingxs`(`symbol`,`updated_at`,`quote_volume`,`active`) VALUES" + megaUpdateString //	Запрос создадим
	//fmt.Println(query2)

	b.InsertDataInMysql(query2)

}
