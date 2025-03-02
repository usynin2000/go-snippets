package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ExchangeRates struct {
	Valute map[string] struct {
		Value float64 `json:"Value"`
	} `json:"Valute"`
}

func getExchangeRate() (float64, error) {
	url := "https://www.cbr-xml-daily.ru/daily_json.js"

	client := http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(url)

	if err != nil {
		return 0, nil
	}

	defer resp.Body.Close()

	// defer откладывает выполнение указанной функции до выхода из текущей функции.
	// Это значит, что resp.Body.Close() выполнится в конце функции, даже если внутри возникнет ошибка.
	// Если не закрывать resp.Body, то:
	// Утечка памяти — Go будет держать соединение в памяти, пока программа работает.
	// Закончатся соединения — если запросов много, программа упрётся в лимит открытых соединений и начнёт падать с ошибкой "too many open files".
	
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("Ошибка запроса, статус код: %d", resp.StatusCode)
	}

	var rates ExchangeRates 
	//  Почему var rates ExchangeRates, а не :=?
	// Можно было бы написать так: rates := ExchangeRates{}
	// Но Go-программисты чаще объявляют структуры через var rates ExchangeRates, потому что:
	// rates создаётся с нулевыми значениями полей автоматически.
	// Не нужно писать ExchangeRates{} явно, Go сам создаст "пустую" структуру.

	err = json.NewDecoder(resp.Body).Decode(&rates)
	// лол тут указатель

	if err != nil {
		return 0, err
	}

	usdRate, exists := rates.Valute["USD"]

	if !exists {
		return 0, fmt.Errorf("Курс USD не найдет")
	}

	return usdRate.Value, nil
}


func main() {
	fmt.Println("Запрашиваем курс USD → RUB...")
	rate, err := getExchangeRate()

	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Текущий курс: 1 USD = %.2f RUB\n", rate)

}