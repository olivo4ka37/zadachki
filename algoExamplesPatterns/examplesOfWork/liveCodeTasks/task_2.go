package main

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
)

// Условия задачи:
// Поочередно выполнить http запросы по списку
// в случае выполнения http-кода ответа на запрос 200 Ok
// печатаем на экране "адрес url - ok"
// в случае получения http-кода ответа на запрос отличного от "200 OK"
// либо в случае ошибки печатаем на экране адрес url - not ok.

// Модифицируйте программу таким образом, чтобы использовались каналы для
// коммуникации основного потока с горутинами. Пример: Запросы по списку
// выполняются в горутинах. Печать результатов происходит в основном потоке.

func main() {
	var urls = []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"http://ya.ru",
		"https://ya.ru",
		"http://ыпвфпвям",
	}

	num := runtime.GOMAXPROCS(0)

	for i := 0; i < num; i++ {

	}

	ctx := context.Background()

	/*
		for i:=0;i<len(urls);i++ {
			resp, err := http.Get(ctx, urls[i])
			if err != nil {
				fmt.Printf("%s - not ok \n", urls[i])
			} else if resp.StatusCode != 200 {
				fmt.Printf("%s - not ok \n", urls[i])
			} else {
				fmt.Printf("%s - ok \n", urls[i])
			}
		}
	*/

	// resp, err = http.Get(ctx, url)
	// resp.StatusCode
}
