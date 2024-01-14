package stakan

/*
https://bybit-exchange.github.io/docs/v5/market/orderbook
аски и биды перепутаны в выдаче
*/
import (
	b "bingxstakan/base"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// Получим обновления
/*
	принимает текущее время
*/
func isValidString(str string) bool {
	reg := regexp.MustCompile(`^[a-zA-Z0-9\-]+$`)
	return reg.MatchString(str)
}

func StartParser(timeNow int64) {

	//	Карта с символами для парсинга, и количество потоков
	symbols, _ := b.GetDataFromMysql(1)

	isValid := isValidString(symbols[0].Symbol)
	if isValid {

		symbolsData := getStakanData(symbols[0].Symbol)

		quoteVolume := strconv.FormatFloat(symbolsData.QuoteVolume, 'f', -1, 64)
		updatedAt := strconv.FormatInt(timeNow, 10)

		askOne := strconv.FormatFloat(symbolsData.AskOne, 'f', -1, 64)
		askDuo := strconv.FormatFloat(symbolsData.AskDuo, 'f', -1, 64)
		bidOne := strconv.FormatFloat(symbolsData.BidOne, 'f', -1, 64)
		bidDuo := strconv.FormatFloat(symbolsData.BidDuo, 'f', -1, 64)
		raznitca := strconv.FormatFloat(symbolsData.Raznitca, 'f', -1, 64)

		megaUpdateString := "('" + symbols[0].Symbol + "','0','" + quoteVolume + "','" + updatedAt + "','" + askOne + "','" + askDuo + "','" + bidOne + "','" + bidDuo + "','" + raznitca + "','0','0','1')"

		updateData(megaUpdateString)
	} else {

		//	Сотрем дичь с базы
		b.InsertDataInMysql("DELETE FROM `bingxs` WHERE `bingxs`.`symbol` = '" + symbols[0].Symbol + "'")

	}

}

func getStakanData(symbol string) b.Bingxs {

	//	Получим данные по обороту
	timeNow := strconv.FormatInt(time.Now().Unix(), 10)

	urlTiker := "https://open-api.bingx.com/openApi/spot/v1/ticker/24hr?symbol=" + symbol + "&timestamp=" + timeNow

	respTiker, err := http.Get(urlTiker)
	if err != nil {

		log.Fatalf("Ошибка при выполнении GET-запроса: %s", err)

	}
	defer respTiker.Body.Close()

	bodyTiker, err := io.ReadAll(respTiker.Body)
	if err != nil {

		log.Fatalf("Ошибка при чтении данных ответа: %s", err)

	}

	var tikersDataStruct TikersDataStruct

	err = json.Unmarshal(bodyTiker, &tikersDataStruct)
	if err != nil {

		//	Ошибка разбора
		fmt.Println("======", respTiker, "пусто")

	}

	if tikersDataStruct.Code != 0 {
		tikersDataStruct.Data[0].QuoteVolume = 0
	}

	//	Получим данные по стакану
	url := "https://open-api.bingx.com/openApi/spot/v1/market/depth?symbol=" + symbol

	resp, err := http.Get(url)
	if err != nil {

		log.Fatalf("Ошибка при выполнении GET-запроса: %s", err)

	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {

		log.Fatalf("Ошибка при чтении данных ответа: %s", err)

	}

	//	Разберем JSON в структуру ошибки
	var stakanDataStruct StakanDataStruct

	err = json.Unmarshal(body, &stakanDataStruct)
	if err != nil {

		//	Ошибка разбора
		fmt.Println(url, "пусто")

	}

	/*
		Если в стакане есть данные по текушему времени, значит торги по этой паре идут
		Если в предложения в Аске или в Бидах отсутствуют то все окау
	*/

	var mysqlDataStruct b.Bingxs

	if stakanDataStruct.Code == 0 {

		if len(stakanDataStruct.Data.Asks) > 0 {
			if len(stakanDataStruct.Data.Bids) > 0 {

				/*
					ask — цена покупки
					Buy - покупка

					bid — цена продажи
					Sell - продажа
				*/

				//  Берем первые данные в стакане с обеих сторн
				mysqlDataStruct.BidOne, _ = strconv.ParseFloat(stakanDataStruct.Data.Bids[0][0], 64)
				mysqlDataStruct.BidDuo, _ = strconv.ParseFloat(stakanDataStruct.Data.Bids[0][1], 64)

				mysqlDataStruct.AskOne, _ = strconv.ParseFloat(stakanDataStruct.Data.Asks[0][0], 64)
				mysqlDataStruct.AskDuo, _ = strconv.ParseFloat(stakanDataStruct.Data.Asks[0][0], 64)

				mysqlDataStruct.Raznitca = (mysqlDataStruct.BidOne/mysqlDataStruct.AskOne - 1) * 100

			} else {

				mysqlDataStruct.BidOne = 0
				mysqlDataStruct.BidDuo = 0
				mysqlDataStruct.Raznitca = 0

			}
		} else {

			mysqlDataStruct.AskOne = 0
			mysqlDataStruct.AskDuo = 0
			mysqlDataStruct.Raznitca = 0

		}
	} else {

		mysqlDataStruct.AskOne = 0
		mysqlDataStruct.AskDuo = 0
		mysqlDataStruct.BidOne = 0
		mysqlDataStruct.BidDuo = 0
		mysqlDataStruct.Raznitca = 0

	}

	//	Время обновления данных
	mysqlDataStruct.UpdatedAt = time.Now().Unix()
	mysqlDataStruct.QuoteVolume = tikersDataStruct.Data[0].QuoteVolume
	mysqlDataStruct.Active = 1

	return mysqlDataStruct
}

// Апдейтим данные
func updateData(megaUpdateString string) {

	//	Занесем свежие данные и включим символы, где есть данные
	query2 := "REPLACE INTO `bingxs`(`symbol`,`volume`,`quote_volume`,`updated_at`,`ask_one`,`ask_duo`,`bid_one`,`bid_duo`,`raznitca`,`trades_count_old`,`trades_count_new`,`active`) VALUES" + megaUpdateString //	Запрос создадим

	b.InsertDataInMysql(query2)

}
