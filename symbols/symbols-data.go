package symbols

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	b "bingxstakan/base"
)

// Получим обновления
/*
	принимает текущее время
*/
func GetSymbolsUpdate(timeNow int64) {

	//	ОТкуда парсим данные
	url := "https://open-api.bingx.com/openApi/spot/v1/common/symbols"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка при выполнении GET-запроса: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Ошибка при чтении данных ответа: %s", err)
	}

	//	Разберем все в структуру
	var symbolsResponse SymbolsResponse
	err = json.Unmarshal(body, &symbolsResponse)
	if err != nil {

		//	Ошибка разбора
		fmt.Println("======", url, "пусто")

	}

	//	Соберем запрос для отправки
	var megaUpdateArray []string

	if symbolsResponse.Code == 0 {

		for _, v := range symbolsResponse.Data.Symbols {

			megaUpdateArray = append(megaUpdateArray, "('"+v.Symbol+"','1')")

		}
	}

	megaUpdateString := strings.Join(megaUpdateArray, ",")

	//	Обновим данные
	updateSymbols(megaUpdateString)

	//	Почистим после себя
	megaUpdateString = ""
	megaUpdateArray = []string{}

	//fmt.Println("Размер массива: ", len(megaUpdateArray))

}

/*
Данные по объему торгов
*/

// Апдейтим данные
func updateSymbols(megaUpdateString string) {

	//	Отключим все символы в БД
	query := "UPDATE `bingxs` SET `active`= '0'" //	Запрос создадим
	b.InsertDataInMysql(query)

	//	Занесем свежие данные и включим символы, где есть данные
	query2 := "REPLACE INTO `bingxs`(`symbol`, `active`) VALUES" + megaUpdateString //	Запрос создадим
	b.InsertDataInMysql(query2)

}
